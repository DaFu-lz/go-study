function checkregister(){
    let Username = document.getElementById("username").value.length
    let Password =document.getElementById("password").value.length
    if(Username===0&&Password===0){
        window.alert("用户名和密码不能为空！")
    }else if (Username===0){
        window.alert("用户名不能为空！")
    }else if (Password===0){
        window.alert("密码不能为空！")
    }
}