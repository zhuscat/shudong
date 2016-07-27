$(function () {
    /*------------------------------------------------用 户 名 检 测 部 分------------------------------------------------*/
    var account_flag = false;
    var pwd_flag = false;
    $('#name-input').keyup(function () {
        $("#name-message").css("display", "none");
        document.getElementById("name-message").innerHTML = "";
    });

    $('#pwd-input').keyup(function () {
        $("#pwd-message").css("display", "none");
        document.getElementById("pwd-message").innerHTML = "";
    });


    $("#signin-submit").click(function () {
        var username=document.getElementById("name-input").value;
        var password=document.getElementById("pwd-input").value;
        if(username==""||password==""){
            if(username==""){
                $("#name-message").css("display", "block");
                document.getElementById("name-message").innerHTML = "!用户名不能为空";
            }
            if(password==""){
                $("#pwd-message").css("display", "block");
                document.getElementById("pwd-message").innerHTML = "!密码不能为空";
            }
        }
        if (username!==""&&password!=="") {
            var data = {
                username: $("#name-input").val(),
                password: $("#pwd-input").val()
            };
            $.ajax({
                type: "post",
                data: data,
                url: "../static/data/signup.json",
                dataType: "json",
                success: function (data) {
                    if (data.success) {
                        alert("登录成功");
                    }
                }
            });
        }
        return false;
    });

});