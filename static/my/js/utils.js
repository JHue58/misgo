const income = '<span class="badge bg-red">收</span>'
const expenditure = '<span class="badge bg-green">支</span>'
const income_text = `<span class="status status-red">
                      {text}
                    </span>`
const expenditure_text = `<span class="status status-green">
                      {text}
                    </span>`

const success_alert = `<div class="alert alert-success alert-dismissible m-0" role="alert">
                  <div class="d-flex">
                    <div>
                      <svg  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-check"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M5 12l5 5l10 -10" /></svg>

                    </div>
                    <div>
                      <h4 class="alert-title">{title}</h4>
                      <div class="text-secondary">{text}</div>
                    </div>
                  </div>
                  <a class="btn-close" data-bs-dismiss="alert" aria-label="close"></a>
                </div>\``
const fail_alert = `<div class="alert alert-danger alert-dismissible m-0" role="alert">
                  <div class="d-flex">
                    <div>
                      <svg  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-alert-circle"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M3 12a9 9 0 1 0 18 0a9 9 0 0 0 -18 0" /><path d="M12 8v4" /><path d="M12 16h.01" /></svg>

                    </div>
                    <div>
                      <h4 class="alert-title">{title}</h4>
                      <div class="text-secondary">{text}</div>
                    </div>
                  </div>
                  <a class="btn-close" data-bs-dismiss="alert" aria-label="close"></a>
                </div>\``

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


function get_primary_button(text){
    let btn = document.createElement("button")

    btn.className = "btn btn-primary btn-icon"
    btn.setAttribute('aria-label', 'Button');

    btn.innerHTML = text
    return btn
}

function get_danger_button(text){
    let btn = document.createElement("button")

    btn.className = "btn btn-danger btn-icon"
    btn.setAttribute('aria-label', 'Button');

    btn.innerHTML = text


    return btn
}

function get_success_alert(title,text){
    return success_alert.replace("{title}",title).replace("{text}",text)
}

function get_fail_alert(title,text){
    return fail_alert.replace("{title}",title).replace("{text}",text)
}

function get_hidden_btn(){
    let btn = document.createElement("button")
    btn.className = "btn btn-primary btn-icon"
    btn.setAttribute('aria-label', 'Button');
    btn.style.visibility = "hidden";
    return btn
}




