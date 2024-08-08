
let global_uid
const loading_div = '<div class="spinner-border"></div>'
const loading_div_2 = `<div class="spinner-grow" role="status"></div>`

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

const trash_svg = `<svg  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-trash"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M4 7l16 0" /><path d="M10 11l0 6" /><path d="M14 11l0 6" /><path d="M5 7l1 12a2 2 0 0 0 2 2h8a2 2 0 0 0 2 -2l1 -12" /><path d="M9 7v-3a1 1 0 0 1 1 -1h4a1 1 0 0 1 1 1v3" /></svg>`
const check_svg = `<svg  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-check"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M5 12l5 5l10 -10" /></svg>`
const x_svg = `<svg  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-x"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M18 6l-12 12" /><path d="M6 6l12 12" /></svg>`
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
        init_transaction_personal(uid)
        global_uid = uid

        init_transaction_personal_daily("7")
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


function init_transaction_personal(uid){
    const radios= document.querySelectorAll('input[name="daily"]');
    const dailyLoading = document.getElementById('daily-loading');
    radios.forEach(radio => {
        radio.addEventListener('change', function (event){
            dailyLoading.style.visibility = "visible"
            event.target.disabled = true
            init_transaction_personal_daily(event.target.value,function (){
                event.target.disabled = false

                dailyLoading.style.visibility = "hidden"
            })

        });
    });

    const userDay = document.getElementById("user-day")
    const userIncome = document.getElementById("user-income")
    const userExpenditure = document.getElementById("user-expenditure")
    const userIncomeSpan = document.getElementById("user-income-span")
    const userBalanceSpan = document.getElementById("user-balance-span")
    const userExpenditureSpan = document.getElementById("user-expenditure-span")
    fetch("api/money/personal",{
        method:"POST",
        headers: {
            'Content-Type': 'application/json' // 指定请求头为JSON格式
        },
        body:JSON.stringify({"uid":uid})
    }).then(response=>response.json()).then(
        data =>{
            // 取得StartTime的Unix时间戳
            const startTime = data.StartTime * 1000; // 转换为毫秒
            // 取得当前时间的Unix时间戳
            const currentTime = new Date().getTime();
            // 计算两个时间戳之间的差异，以毫秒为单位
            const differenceInMilliseconds = currentTime - startTime;
            // 将毫秒差异转换为天数
            const days = Math.ceil(differenceInMilliseconds / (1000 * 60 * 60 * 24));
            userDay.innerHTML = "记账"+days+"天"
            userIncome.innerHTML = "收入"+data.IncomeCount+"次"
            userExpenditure.innerHTML = "支出"+data.ExpenditureCount+"次"
            userIncomeSpan.innerHTML = data.Income
            userBalanceSpan.innerHTML = data.Balance
            userExpenditureSpan.innerHTML = data.Expenditure
        }
    )
}

function init_transaction_personal_daily(days,whenUpdated) {
    const uid = global_uid



    // 三个对应的图表
    const chartIncomePie = document.getElementById("chart-income-pie")
    const chartExpenditurePie = document.getElementById("chart-expenditure-pie")
    const chartDaily = document.getElementById('chart-daily')


    // 获取当前日期
    let currentDate = new Date();
    let endTime = Math.floor(currentDate.getTime() / 1000); // 当前时间的Unix时间戳（秒）

    // 计算days天前的日期
    let startDate = new Date();
    startDate.setDate(currentDate.getDate() - parseInt(days));
    startDate.setHours(0, 0, 0, 0); // 将时间部分设置为0点
    let startTime = Math.floor(startDate.getTime() / 1000); // days天前0点的Unix时间戳（秒）

    // 构造参数对象
    const params = {
        uid: uid,
        time_range: "自定义",
        start_time: startTime,
        end_time: endTime,
        condition: "", // 这里可以根据需要填写查询条件
        order: "time",
        count: -1
    };

    fetch("api/money",{
        method:"POST",
        headers: {
            'Content-Type': 'application/json' // 指定请求头为JSON格式
        },
        body:JSON.stringify(params)
    }).then(response=>response.json()).then(data=>{
        const transactions = data.Transactions
        let expenditures = {}; // 支出字典
        let incomes = {}; // 收入字典
        let dailyExpenditures = {}; // 每日支出字典
        let dailyIncomes = {}; // 每日收入字典
        let dailies = {}; // 每日日期字典

        transactions.forEach(transaction => {
            let { type, category, amount, time } = transaction;

            // 将 time 转换为 Date 对象并格式化为 MM-DD
            let date = new Date(time * 1000);
            let dateKey = `${date.getMonth() + 1}-${date.getDate()}`; // MM-DD 格式
            dailies[dateKey] = null
            if (type === "支出") {
                // 更新支出字典
                if (expenditures[category]) {
                    expenditures[category] += amount;
                } else {
                    expenditures[category] = amount;
                }

                // 更新每日支出字典
                if (dailyExpenditures[dateKey]) {
                    dailyExpenditures[dateKey] += amount;
                } else {
                    dailyExpenditures[dateKey] = amount;
                }
            } else if (type === "收入") {
                // 更新收入字典
                if (incomes[category]) {
                    incomes[category] += amount;
                } else {
                    incomes[category] = amount;
                }

                // 更新每日收入字典
                if (dailyIncomes[dateKey]) {
                    dailyIncomes[dateKey] += amount;
                } else {
                    dailyIncomes[dateKey] = amount;
                }
            }
        });

        chartIncomePie.innerHTML = ""
        window.ApexCharts && (new ApexCharts(chartIncomePie, {
            chart: {
                type: "donut",
                fontFamily: 'inherit',
                height: 240,
                sparkline: {
                    enabled: true
                },
                animations: {
                    enabled: true
                },
            },

            fill: {
                opacity: 1,
            },
            series: Object.values(incomes),
            labels: Object.keys(incomes),
            tooltip: {
                theme: 'dark',
                y: {
                    formatter: function (val) {
                        return '￥' + val.toFixed(2); // 工具提示也格式化为金额模式，四舍五入到两位小数
                    }
                },
                fillSeriesColor: false
            },
            grid: {
                strokeDashArray: 4,
            },
            //colors: [tabler.getColor("green"),tabler.getColor("red"),tabler.getColor("black")],
            legend: {
                show: true,
                position: 'bottom',
                offsetY: 12,
                markers: {
                    width: 10,
                    height: 10,
                    radius: 100,
                },
                itemMargin: {
                    horizontal: 8,
                    vertical: 8
                },
            },

        })).render();


        chartExpenditurePie.innerHTML = ""
        window.ApexCharts && (new ApexCharts(chartExpenditurePie, {
            chart: {
                type: "donut",
                fontFamily: 'inherit',
                height: 240,
                sparkline: {
                    enabled: true
                },
                animations: {
                    enabled: true
                },
            },

            fill: {
                opacity: 1,
            },
            series: Object.values(expenditures),
            labels: Object.keys(expenditures),
            tooltip: {
                theme: 'dark',
                y: {
                    formatter: function (val) {
                        return '￥' + val.toFixed(2); // 工具提示也格式化为金额模式，四舍五入到两位小数
                    }
                },
                fillSeriesColor: false
            },
            grid: {
                strokeDashArray: 4,
            },
            //colors: [tabler.getColor("green"),tabler.getColor("red"),tabler.getColor("black")],
            legend: {
                show: true,
                position: 'bottom',
                offsetY: 12,
                markers: {
                    width: 10,
                    height: 10,
                    radius: 100,
                },
                itemMargin: {
                    horizontal: 8,
                    vertical: 8
                },
            },

        })).render();


        chartDaily.innerHTML = ""
        window.ApexCharts && (new ApexCharts(chartDaily, {
            chart: {
                type: "area",
                fontFamily: 'inherit',
                height: 240,
                parentHeightOffset: 0,
                toolbar: {
                    show: false,
                },
                animations: {
                    enabled: true
                },
            },

            fill: {
                opacity: 1,
            },
            dataLabels: {
                enabled: false
            },
            stroke: {
                width: 2,
                lineCap: "round",
                curve: "straight",
            },
            series: [{
                name: "收入",
                data: Object.values(dailyIncomes)
            },{name:"支出",data:Object.values(dailyExpenditures)}],
            tooltip: {
                theme: 'dark'
            },
            grid: {
                padding: {
                    top: -20,
                    right: 0,
                    left: -4,
                    bottom: -4
                },
                strokeDashArray: 4,
            },
            xaxis: {
                labels: {
                    padding: 0,
                },
                tooltip: {
                    enabled: false
                },
                tickAmount: 'dataPoints', // 根据实际数据点数量显示刻度
                tickPlacement: 'on', // 确保刻度放置在数据点上
                //type: 'datetime',
            },
            yaxis: {
                labels: {
                    padding: 4,
                    formatter: function (val) {
                        return '￥' + val.toFixed(2); // Y轴标签也显示为金额模式，四舍五入到两位小数
                    }
                },

            },
            labels: Object.keys(dailies),
            colors: [ tabler.getColor("red"), tabler.getColor("green")],
            legend: {
                show: true,
                position: 'bottom',
                offsetY: 12,
                markers: {
                    width: 10,
                    height: 10,
                    radius: 100,
                },
                itemMargin: {
                    horizontal: 8,
                    vertical: 8
                },
            },
        })).render();

        if (whenUpdated!=null){
            whenUpdated();
        }

    })

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
    if (params.condition===""){
        viewTitleCond.innerHTML = "ALL"
    }else{
        viewTitleCond.innerHTML = params.condition
    }


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

    const balanceDiv = document.getElementById('balance-view')
    if ('Balance' in viewData){
        if (viewData.Balance>0){
            balanceDiv.innerHTML = "+"+viewData.Balance
            balanceDiv.className = 'status status-green'
        }else {
            balanceDiv.innerHTML = viewData.Balance
            balanceDiv.className = 'status status-red'
        }
    }else{
        balanceDiv.innerHTML = "0"
        balanceDiv.className = 'status status-cyan'
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
    if (params.condition===""){
        billTitleCond.innerHTML = "ALL"
    }else{
        billTitleCond.innerHTML = params.condition
    }

}


function delete_func(event){
    const target = event.currentTarget
    const parma = JSON.parse(target.value)
    const btnDiv = document.getElementById("btn-div-"+parma.id)

    btnDiv.innerHTML = ''

    const primaryBtn =  get_primary_button(check_svg)

    const dangerBtn = get_danger_button(x_svg)
    const alertDiv = document.getElementById("table-alert")
    primaryBtn.addEventListener("click",function (){
        btnDiv.innerHTML = loading_div


        let req = {
            uid :document.getElementById('uid-input').value,
            transaction:parma
        }

        fetch('api/money',{
            method : 'DELETE',
            headers: {
                'Content-Type': 'application/json' // 指定请求头为JSON格式
            },
            body:JSON.stringify(req),
        }).then(function(response) {
            return response.json();
        }).then(function (data){
            btnDiv.innerHTML = ''
            btnDiv.appendChild(get_hidden_btn())
            btnDiv.appendChild(target)
            btnDiv.appendChild(get_hidden_btn())
            if ('Code' in data){
                if (data.Code>0){
                    // 失败
                    alertDiv.innerHTML = get_fail_alert("失败！",data.Message)
                    return
                }
            }
            // 成功
            alertDiv.innerHTML = get_success_alert("成功！","你删除了一条记录")
            bill_ltb.update_page()

        }).catch(function (error){
            alertDiv.innerHTML = get_fail_alert("失败！",error)
            btnDiv.innerHTML = ''
            btnDiv.appendChild(get_hidden_btn())
            btnDiv.appendChild(target)
            btnDiv.appendChild(get_hidden_btn())
        })
    })
    dangerBtn.addEventListener("click",function (){
        btnDiv.innerHTML = ''
        btnDiv.appendChild(get_hidden_btn())
        btnDiv.appendChild(target)
        btnDiv.appendChild(get_hidden_btn())
    })

    btnDiv.appendChild(primaryBtn)
    btnDiv.appendChild(get_hidden_btn())
    btnDiv.appendChild(dangerBtn)

}

function bill_inner(limit,offset){

    const billTable = document.getElementById("bill-table")
    return  new Promise(function(resolve, reject) {
        if (global_params==null){
            billTable.innerHTML = null_div
            return
        }

        //billTable.innerHTML = loading_div
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
                    <th></th>
                
                  </tr>
                  </thead>

                  <tbody id="bill-tbody">
                    
                  </tbody>
            `

            const transactions = responseData.Transactions

            const tbody = document.getElementById("bill-tbody")
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

                const th6 = document.createElement('th')



                const deleteBtn = get_danger_button(trash_svg)


                const btnDiv = document.createElement('div')
                btnDiv.id = "btn-div-" + transaction.id
                deleteBtn.value = JSON.stringify(transaction)
                deleteBtn.addEventListener('click', delete_func)
                btnDiv.appendChild(get_hidden_btn())
                btnDiv.appendChild(deleteBtn)
                btnDiv.appendChild(get_hidden_btn())
                th6.appendChild(btnDiv)


                tr.appendChild(th1)
                tr.appendChild(th2)
                tr.appendChild(th3)
                tr.appendChild(th4)
                tr.appendChild(th5)
                tr.appendChild(th6)
                tbody.appendChild(tr)



            })
            Promise.all([tForEach]).then(()=>{

                if (tbody.innerHTML===''){
                    billTable.innerHTML = null_div
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