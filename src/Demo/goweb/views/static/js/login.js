function check(){
    let Username=document.getElementById("username").value.length
    let Password =document.getElementById("password").value.length

    let msg = document.getElementById("msg").value
    if(Username===0&&Password===0){
        // window.alert("用户名不能为空")
        msg.innerText("请输入用户名和密码")
    }else if (Username===0){
        // window.alert("用户名不能为空")
        msg.innerText("用户名不能为空")
    }else if(Password===0) {
        // window.alert("密码不能为空")
        msg.innerText("密码不能为空")
    }
}

$(document).ready(function () {
    $("#msg").hide()
})

