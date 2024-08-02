import sqlite3
import uuid
import sys

def generate_uid():
    return uuid.uuid4().hex  # 生成一个长度为 32 的十六进制字符串

def main(name):
    # 连接到 SQLite 数据库（如果文件不存在，会自动创建）
    conn = sqlite3.connect('misgo.db')
    cursor = conn.cursor()

    
    # 生成随机 UID
    uid = generate_uid()

    # 插入或更新数据
    cursor.execute('''
    INSERT OR IGNORE INTO users (UID, Name)
    VALUES (?, ?)
    ''', (uid, name))

    cursor.execute('''
    UPDATE users
    SET UID = ?
    WHERE Name = ?
    ''', (uid, name))



    # 提交事务
    conn.commit()

    # 打印 UID
    print(f'UID: {uid}')

    # 关闭连接
    conn.close()

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python register.py <Name>")
    else:
        name = sys.argv[1]
        main(name)
