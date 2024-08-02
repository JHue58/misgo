import argparse
from datetime import datetime
import numpy as np
import pandas as pd
import openpyxl
import sqlite3

import requests

addr = "http://127.0.0.1:8888/"

class MisgoDB:
    

    def __enter__(self):
        self.db  = sqlite3.connect('misgo.db')
        return self
        
    def __exit__(self, exc_type, exc_value, traceback):
        self.db.close()



class Transcation:
    time_data = 0
    type_data = "支出"
    amount_data = 0.0
    note_data = "无"
    category_data = ""
    
    def to_json(self) -> dict:
        json_dict = {
            "time":self.time_data,
            "type":self.type_data,
            "amount":self.amount_data,
            "note":self.note_data,
            "category":self.category_data
        }
        return json_dict
        

    def __str__(self) -> str:
        dt = datetime.fromtimestamp(self.time_data)
        formatted_time = dt.strftime('%Y-%m-%d %H:%M:%S')
        s = f'{formatted_time} {self.category_data} {self.type_data} {self.amount_data} {self.note_data}'
        return s


class YiMu():
    
    def __init__(self,path) -> None:
        df = pd.read_excel(path)
        self.df = df

   

    def get_transcations(self) -> list:
        # 获取所有列名
        column_names = self.df.columns.tolist()
        print("列名：", column_names)
        frame = {}
        value_count = 0
        for column in column_names:
            values = self.df[column].tolist()
            value_count = len(values)
            print(f"列 '{column}' 获取'{value_count}'个值")
            frame[column] = values
            
        # 日期列名称
        date_cloumn = "日期"
        # 类型列名称
        type_cloumn = "收支类型"
        # 金额列名称
        amount_cloumn = "金额"
        # 备注列名称
        note_cloumn = "备注"
        # 来源/用途列名称
        category_cloumn = "子类"
        # 当category_cloumn无值时用这个
        backup_category_cloumn = "类别"
        ledger_cloumn = "账本"

        transcations = []

        time_format = "%Y-%m-%d %H:%M"
        for i in range(value_count):
            if frame[ledger_cloumn][i]!="日常账本":
                continue
            
            category = frame[category_cloumn][i]
            if pd.isna(category):
                category = frame[backup_category_cloumn][i]
            note = frame[note_cloumn][i]
            if pd.isna(note):
                note = "无"
            amount = str(frame[amount_cloumn][i])
            amount=amount.replace("元","")
            amount=amount.replace("-","")
            amount=amount.replace("+","")
            
            _type = frame[type_cloumn][i]
            date = frame[date_cloumn][i]
            
            dt = datetime.strptime(date, time_format)
            time = int(dt.timestamp())
            
            transcation = Transcation()
            transcation.time_data = int(time)
            transcation.category_data = category
            transcation.type_data = _type
            transcation.amount_data = float(amount)
            transcation.note_data = note

            transcations.append(transcation)

        print(f'解析出{len(transcations)}条数据')

        return transcations
    

def write_to_misgo(uid:str,transactions:list):
   url = addr+"/api/money"
   headers = {
    'Content-Type': 'application/json'
}
   idx = 0
   for transaction in transactions:
        parma = {
            'uid':uid,
            'transaction':transaction.to_json()
        }
        response = requests.put(url,json=parma,headers=headers)
        # 检查响应状态码
        if response.status_code == 200:
            # 解析 JSON 响应
            response_data = response.json()
            print(f"{idx}:Response data:", response_data)
        else:
            print(f"Request failed with status code: {response.status_code}")
            print(f"{idx}:Response text:", response.text)

if __name__ == "__main__":
    # 创建解析器
    parser = argparse.ArgumentParser(description="解析命令行参数")

    # 添加参数
    parser.add_argument('uid', type=str, help='UID')
    parser.add_argument('ledger', type=str, help='账本类型')
    parser.add_argument('path',type=str,help='表格文件路径')

    # 解析参数
    args = parser.parse_args()

    # 打印解析的参数
    print(f"UID: {args.uid}")
    print(f"Ledger: {args.ledger}")
    print(f"Path: {args.path}")

    if args.ledger == "yimu":
        yi_mu = YiMu(args.path)
        transactions = yi_mu.get_transcations()
        print(len(transactions))

        for transaction in transactions:
            print(transaction)
    else:
        raise Exception("不支持的账单类型")
    
    write_to_misgo(args.uid,transactions)