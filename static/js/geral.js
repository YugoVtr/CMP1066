
function setCheckedCheckBox(checkbox_name) {
    var selector = "[name=" + checkbox_name + "]"; 
    if ($(selector).val() == false) {
        $(selector).prop('checked', false);
    } else { 
        $(selector).prop('checked', true);
    }
}

function setValueCheckBox(checkbox_name, value_checked, value_unchecked) {
    $("[name=" + checkbox_name + "]").click(function() {
        if ($(this).is(":checked")) {
            $(this).val(value_checked);
        } else { 
            $(this).val(value_unchecked)
        }
    });
}

function getXSRFData() {
    var token = $("input[name=_xsrf]").val()
    if(token) {
        return token;
    } else {
        return null; 
    }
}

function redirect(url) {
    location.href = url;
}