const loading_div = '<div class="spinner-border"></div>'

const null_div = '<div class="empty">\n' +
    '  <div class="empty-icon">\n' +
    '    <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">\n' +
    '      <path stroke="none" d="M0 0h24v24H0z" fill="none" />\n' +
    '      <circle cx="12" cy="12" r="9" />\n' +
    '      <line x1="9" y1="10" x2="9.01" y2="10" />\n' +
    '      <line x1="15" y1="10" x2="15.01" y2="10" />\n' +
    '      <path d="M9.5 15.25a3.5 3.5 0 0 1 5 0" />\n' +
    '    </svg>\n' +
    '  </div>\n' +
    '  <p class="empty-title">未找到结果</p>\n' +
    '  <p class="empty-subtitle text-muted">\n' +
    '    尝试添加记录或重试\n' +
    '  </p>\n' +
    '</div>\n'

const viewCard = document.getElementById('view-card')
const billCard = document.getElementById('bill-card')
query_select_callback()
change_card_view(false)
time_range_select_callback()
uid_init()
function change_card_view(block){
    if (block){
        viewCard.style.display = "block"
        billCard.style.display = "block"
    }else{
        viewCard.style.display = "none"
        billCard.style.display = "none"
    }

}

function query_select_callback(){
    const queryDiv = document.getElementById('query-div');
    const conditionInput = document.getElementById('condition');
    const orderInput = document.getElementById('order');
    const radios = document.querySelectorAll('input[name="query"]');
    queryDiv.style.display = "none"
    radios.forEach(radio => {
        radio.addEventListener('change', function (event){
            if (event.target.value ==="advanced"){
                queryDiv.style.display = "block"
            }else{
                queryDiv.style.display = "none"
                conditionInput.value = ""
                orderInput.value = "time desc"
            }
        });
    });
}



function time_range_select_callback(){

    let range = document.getElementById('time-range')
    let timeSelect = document.getElementById('time-select')
    timeSelect.style.display = "none"
    range.addEventListener('change',function (){
        console.log(range.value)
        if (range.value === "自定义"){
            timeSelect.style.display = "block"
        }else{
            timeSelect.style.display = "none"
        }
    })

}

function uid_init() {
    // 获取当前 URL
    const url = new URL(window.location.href);

    // 创建 URLSearchParams 对象
    const params = new URLSearchParams(url.search);

    // 获取 'uid' 参数的值
    const uid = params.get('uid');

    console.log('UID:', uid);
    let uid_input = document.getElementById('uid-input')

    if (uid!=null){
        init_user_info(uid)
    }

}


function init_user_info(uid){
    let url = 'api/user?uid='+uid
    let userName = document.getElementById('user-name')
    let uidInput = document.getElementById('uid-input')
    const fetchUser = fetch(url,{method:'Get'})
        .then(response => response.json())
        .then(data => {
            console.log(data)
            console.log(data.Code<=0)
            if ('Code' in data){
            }else {
                userName.innerHTML = data.Name
                uidInput.value = data.UID
                uidInput.readonly = true
            }

        })
        .catch(error => console.error(error));
}

function transaction_search(){


    const loadingModal = document.querySelector('#transaction-search-callback-modal');
    const uidInput = document.querySelector('#transaction-search-modal input[name="uid"]');
    const timeRangeInput = document.querySelector('#transaction-search-modal select[name="time_range"]');
    const startTimeInput = document.querySelector('#transaction-search-modal input[name="start_time"]');
    const endTimeInput = document.querySelector('#transaction-search-modal input[name="end_time"]');
    const conditionInput = document.querySelector('#transaction-search-modal input[name="condition"]');
    const orderInput = document.querySelector('#transaction-search-modal input[name="order"]');
    const startTime = new Date(startTimeInput.value).getTime() / 1000; // 将毫秒时间戳转换为秒时间戳
    const endTime = new Date(endTimeInput.value).getTime() / 1000;     // 将毫秒时间戳转换为秒时间戳
    const params = {
        uid : uidInput.value,
        time_range : timeRangeInput.options[timeRangeInput.selectedIndex].value,
        start_time : startTimeInput.value,
        end_time : endTimeInput.value,
        condition : conditionInput.value,
        order : orderInput.value,
    }
    const req = {
        uid : uidInput.value,
        time_range : timeRangeInput.options[timeRangeInput.selectedIndex].value,
        start_time : startTime,
        end_time : endTime,
        condition : conditionInput.value,
        order : orderInput.value,
    }

    transaction_search_init(params)
    fetch('api/money/view',{
        method : 'POST',
        headers: {
            'Content-Type': 'application/json' // 指定请求头为JSON格式
        },
        body:JSON.stringify(req),
    }).then(function(response) {
        return response.json()
    }).then(function (viewData){
        if ('Code' in viewData) {
            loadingModal.innerHTML = get_fail_modal("查询失败",viewData.Message)
            change_card_view(false)
        }else{
            loadingModal.innerHTML = get_success_modal("查询成功","结果数:"+viewData.Count)
            transaction_view_init(params,viewData)
            change_card_view(true)
            // ResInit
            transaction_search_init(params)
        }

    }).catch(function(error) {
        loadingModal.innerHTML = get_fail_modal("查询失败",error)
        change_card_view(false)
    })
}

function transaction_view_init(params,viewData) {
    const viewTitle = document.getElementById('view-title');
    const viewTitleCond = document.getElementById("view-title-condition");
    if (params.time_range==="自定义"){
        viewTitle.textContent = params.start_time+" 到 "+params.end_time+" 的总览"
    }else{
        viewTitle.textContent = params.time_range+"的总览"
    }
    viewTitleCond.innerHTML = params.condition

    const countDiv = document.getElementById("count-view")
    countDiv.innerHTML = viewData.Count
    const incomeCountDiv = document.getElementById("income-count-view")
    if ('IncomeCount' in viewData){
        incomeCountDiv.innerHTML = viewData.IncomeCount
    }else{
        incomeCountDiv.innerHTML = "无"
    }
    const expenditureCountDiv = document.getElementById("expenditure-count-view")
    if ('ExpenditureCount' in viewData){
        expenditureCountDiv.innerHTML = viewData.ExpenditureCount
    }else{
        expenditureCountDiv.innerHTML = "无"
    }
    const incomeDiv = document.getElementById('income-view')
    if ('Income' in viewData){
        incomeDiv.innerHTML = "+"+viewData.Income
    }else{
        incomeDiv.innerHTML = "无"
    }
    const expenditureDiv = document.getElementById('expenditure-view')
    if ('Expenditure' in viewData){
        expenditureDiv.innerHTML = "-"+viewData.Expenditure
    }else{
        expenditureDiv.innerHTML = "无"
    }
    const largestIncomeDiv = document.getElementById('largest-income-view')
    if ('LargestIncome' in viewData){
        let date = formatTimestamp(viewData.LargestIncome.time)
        largestIncomeDiv.innerHTML = date+" "+viewData.LargestIncome.category+" 收入 "+viewData.LargestIncome.amount+"元"
    }else{
        largestIncomeDiv.innerHTML = "无"
    }


    const largestExpenditureDiv = document.getElementById('largest-expenditure-view')
    if ('LargestExpenditure' in viewData){
        let date = formatTimestamp(viewData.LargestExpenditure.time)
        largestExpenditureDiv.innerHTML = date+" "+viewData.LargestExpenditure.category+" 支出 "+viewData.LargestExpenditure.amount+"元"
    }else{
        largestExpenditureDiv.innerHTML = "无"
    }

}

let global_params
const bill_ltb=new LimitTables(document.getElementById('bill-limit-table'),10,'btn btn-success',bill_inner)
function transaction_search_init(params){
    const startTime = new Date(params.start_time).getTime() / 1000; // 将毫秒时间戳转换为秒时间戳
    const endTime = new Date(params.end_time).getTime() / 1000;     // 将毫秒时间戳转换为秒时间戳
    global_params = {
        uid : params.uid,
        time_range : params.time_range,
        start_time : startTime,
        end_time : endTime,
        condition : params.condition,
        order : params.order,
    }


    bill_ltb.update_table()


    const billTitle = document.getElementById('bill-title');
    const billTitleCond = document.getElementById("bill-title-condition");
    if (params.time_range==="自定义"){
        billTitle.textContent = params.start_time+" 到 "+params.end_time+" 的账单"
    }else{
        billTitle.textContent = params.time_range+"的账单"
    }
    billTitleCond.innerHTML = params.condition
}

function bill_inner(limit,offset){

    const billTable = document.getElementById("bill-table")
    return  new Promise(function(resolve, reject) {
        if (global_params==null){
            billTable.innerHTML = null_div
            return
        }

        billTable.innerHTML = loading_div
        let req = {
            count : limit,
            start : offset,
            uid : global_params.uid,
            time_range : global_params.time_range,
            start_time : global_params.start_time,
            end_time : global_params.end_time,
            condition : global_params.condition,
            order : global_params.order,
        }

        fetch('api/money',{
            method : 'POST',
            headers: {
                'Content-Type': 'application/json' // 指定请求头为JSON格式
            },
            body:JSON.stringify(req),
        }).then(function(response) {
            return response.json();
        }).then(function(responseData){
            if (!('Transactions' in responseData)){
                billTable.innerHTML = null_div
                resolve(0)
                return
            }
            billTable.innerHTML = `
                <thead>
                  <tr>
                    <th>日期</th>
                    <th>类型</th>
                    <th>来源/用途</th>
                    <th>金额</th>
                    <th>备注</th>
                  </tr>
                  </thead>

                  <tbody id="bill-tbody">
                    {tbs}
                  </tbody>
            `

            const transactions = responseData.Transactions
            const tbs = document.createElement('div')
            const tForEach = transactions.forEach(transaction=>{

                const tr = document.createElement('tr')
                const th1 = document.createElement('th')
                th1.innerHTML = formatTimestamp(transaction.time)
                const th2 = document.createElement('th')
                th2.innerHTML = get_score_title_span(transaction.type)
                const th3 = document.createElement('th')
                th3.innerHTML = get_score_title_span(transaction.type,transaction.category)
                const th4 = document.createElement('th')
                th4.innerHTML = get_score_title_span(transaction.type,transaction.amount)
                const th5 = document.createElement('th')
                th5.innerHTML = transaction.note



                tr.appendChild(th1)
                tr.appendChild(th2)
                tr.appendChild(th3)
                tr.appendChild(th4)
                tr.appendChild(th5)


                tbs.appendChild(tr)

            })
            Promise.all([tForEach]).then(()=>{
                if (tbs.innerHTML===''){
                    billTable.innerHTML = null_div
                }
                else {
                    billTable.innerHTML =billTable.innerHTML.replace("{tbs}",tbs.innerHTML)
                }
                resolve(transactions.length); // 将game_infos数组的长度作为结果传递
            })
        }).catch(function (error) {
            reject(error)
        })
    });
}

function formatTimestamp(timestamp) {
    // 将秒级时间戳转换为毫秒级时间戳
    const date = new Date(timestamp * 1000);

    // 获取年份、月份、日期、小时和分钟
    const year = date.getFullYear();
    const month = date.getMonth() + 1; // 月份从0开始，所以需要加1
    const day = date.getDate();
    const hours = date.getHours();
    const minutes = date.getMinutes();

    // 格式化为两位数
    const formattedMonth = month.toString().padStart(2, '0');
    const formattedDay = day.toString().padStart(2, '0');
    const formattedHours = hours.toString().padStart(2, '0');
    const formattedMinutes = minutes.toString().padStart(2, '0');

    return `${year}年${formattedMonth}月${formattedDay}日 ${formattedHours}:${formattedMinutes}`;
}