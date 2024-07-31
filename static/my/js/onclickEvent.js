const loading_modal = '<div class="modal-dialog" role="document">\n' +
    '                                    <div class="modal-content">\n' +
    '                                        <div class="modal-header">\n' +
    '                                            <h5 class="modal-title">{title}</h5>\n' +
    '                                        </div>\n' +
    '                                        <div class="modal-body">\n' +
    '                                            <div class="spinner-border"></div>\n' +
    '                                        </div>\n' +
    '                                    </div>\n' +
    '                                </div>'

const success_modal = '<div class="modal-dialog modal-sm" role="document">\n' +
    '    <div class="modal-content">\n' +
    '      <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>\n' +
    '      <div class="modal-status bg-success"></div>\n' +
    '      <div class="modal-body text-center py-4">\n' +
    '        <svg xmlns="http://www.w3.org/2000/svg" class="icon mb-2 text-green icon-lg" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">\n' +
    '          <path stroke="none" d="M0 0h24v24H0z" fill="none" />\n' +
    '          <circle cx="12" cy="12" r="9" />\n' +
    '          <path d="M9 12l2 2l4 -4" />\n' +
    '        </svg>\n' +
    '        <h3>{title}</h3>\n' +
    '        <div class="text-muted">{msg}</div>\n' +
    '      </div>\n' +
    '    </div>\n' +
    '  </div>'

const fail_modal = '<div class="modal-dialog modal-sm" role="document">\n' +
    '    <div class="modal-content">\n' +
    '      <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>\n' +
    '      <div class="modal-status bg-danger"></div>\n' +
    '      <div class="modal-body text-center py-4">\n' +
    '        <svg xmlns="http://www.w3.org/2000/svg" class="icon mb-2 text-danger icon-lg" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">\n' +
    '          <path stroke="none" d="M0 0h24v24H0z" fill="none" />\n' +
    '          <path d="M12 9v2m0 4v.01" />\n' +
    '          <path d="M5 19h14a2 2 0 0 0 1.84 -2.75l-7.1 -12.25a2 2 0 0 0 -3.5 0l-7.1 12.25a2 2 0 0 0 1.75 2.75" />\n' +
    '        </svg>\n' +
    '        <h3>{title}</h3>\n' +
    '        <div class="text-muted">{msg}</div>\n' +
    '      </div>\n' +
    '      <div id="error-info" class="modal-body text-center py-4">\n' +
    '      </div>\n' +
    '    </div>\n' +
    '  </div>\n'

function get_success_modal(title,msg){
    return success_modal.replace("{title}",title).replace("{msg}",msg)
}
function get_fail_modal(title,msg,errorInfo){
    var html = fail_modal.replace("{title}",title).replace("{msg}",msg);
    if (errorInfo==null){
        return html
    }
    // 将HTML字符串转换为DOM节点
    var tempDiv = document.createElement('div');
    tempDiv.innerHTML = html;

    // 获取错误信息的div元素
    var errorInfoDiv = tempDiv.querySelector('#error-info');
    // 创建按钮元素
    var button = document.createElement('div');
    button.setAttribute('class', 'alert alert-danger');
    button.setAttribute('role', 'alert');
    button.textContent = errorInfo;
    // 将按钮添加到错误信息的div元素中
    errorInfoDiv.appendChild(button);
    // 获取修改后的HTML字符串
    var modifiedHtml = tempDiv.innerHTML;
    return modifiedHtml
}

function get_loading_modal(title){
    return loading_modal.replace("{title}",title)
}



function transaction_default(){
    const callbackModal = document.querySelector('#transaction-search-callback-modal');
    callbackModal.innerHTML = get_loading_modal("等待响应中...")
}















