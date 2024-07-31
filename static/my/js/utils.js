const income = '<span class="badge bg-red">收</span>'
const expenditure = '<span class="badge bg-green">支</span>'
const income_text = `<span class="status status-red">
                      {text}
                    </span>`
const expenditure_text = `<span class="status status-green">
                      {text}
                    </span>`
function get_score_title_span(types,text) {
    if (text==null){
        switch (types){
            case '收入':return income;
            case '支出':return expenditure;
        }
    }else {
        switch (types){
            case '收入':return income_text.replace("{text}",text);
            case '支出':return expenditure_text.replace("{text}",text);
        }
    }

}



