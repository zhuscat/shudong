$(function () {
    /*------------------------------------------------用 户 名 检 测 部 分------------------------------------------------*/
    var account_flag = false;
    var pwd_flag = false;
    var email_flag = false;
    $('#name-input').keyup(function () {
        $("#name-message").css("display", "none");
        document.getElementById("name-message").innerHTML = "";
    });
    $('#name-input').blur(function () {
        account_flag = false;
        var account = document.getElementById("name-input").value;
        if (account.length > 0 && account.length < 5) {
            $("#name-message").css("display", "block");
            document.getElementById("name-message").innerHTML = "*用户名小于5个字符";
        }
        if (account.length > 20) {
            $("#name-message").css("display", "block");
            document.getElementById("name-message").innerHTML = "*用户名超过20个字符";
        }
        if (account.length == 0) {
            $("#name-message").css("display", "block");
            document.getElementById("name-message").innerHTML = "*用户名不能为空";
        }
        if (account.length >= 5 && account.length <= 20) {
            $("#name-message").css("display", "none");
            document.getElementById("name-message").innerHTML = "";
            account_flag = true;
        }
    });
    /*-------------------------------------------------邮 箱 检 测 部 分--------------------------------------------------*/
    $('#email-input').keyup(function () {
        $("#email-message").css("display", "none");
        document.getElementById("email-message").innerHTML = "";
    });
    $('#email-input').blur(function () {
        email_flag = false;
        var pattern = /^(\w)+(\.\w+)*@(\w)+((\.\w{2,3}){1,3})$/;
        var eamil = document.getElementById("email-input").value;
        if (!pattern.test(eamil)) {
            $("#email-message").css("display", "block");
            document.getElementById("email-message").innerHTML = "*邮箱不符合要求";
        } else {
            email_flag = true;
        }
        if (email.length == 0) {
            document.getElementById("email-message").innerHTML = "*邮箱不能为空";
        }
    });
    /*-------------------------------------------------密 码 检 测 部 分--------------------------------------------------*/
    $('#pwd-input').keyup(function () {
        $("#pwd-message").css("display", "none");
        document.getElementById("pwd-message").innerHTML = "";
    });
    $('#pwdconf-input').keyup(function () {
        $("#pwdconf-message").css("display", "none");
        document.getElementById("pwdconf-message").innerHTML = "";
    });
    $('#pwdconf-input').blur(function () {
        pwd_flag = false;
        var pwd = document.getElementById('pwd-input').value;
        var pwdconf = document.getElementById('pwdconf-input').value;
        if (pwd.length >= 6 && pwd.length <= 20) {
            if (pwd != pwdconf) {
                $("#pwdconf-message").css("display", "block");
                document.getElementById("pwdconf-message").innerHTML = "*两次输入的密码不同，请重新输入";
            }
            if (pwd == pwdconf && pwdconf != "") {
                pwd_flag = true;//此时密码符合要求
            }
        }
        if (pwd.length < 6) {
            $("#pwdconf-message").css("display", "none");
            document.getElementById("pwdconf-message").innerHTML = "";
        }
        if (pwd.length > 20) {
            $("#pwdconf-message").css("display", "none");
            document.getElementById("pwdconf-message").innerHTML = "";
        }
    });

    $('#pwd-input').blur(function () {
        var pwd = document.getElementById('pwd-input').value;
        var pwdconf = document.getElementById('pwdconf-input').value;
        if (pwd.length > 0 && pwd.length < 6) {
            $("#pwd-message").css("display", "block");
            document.getElementById("pwd-message").innerHTML = "*密码小于6个字符，请重新输入";
        }
        if (pwd.length > 20) {
            $("#pwd-message").css("display", "block");
            document.getElementById("pwd-message").innerHTML = "*密码大于20个字符，请重新输入";
        }
        if (pwd.length >= 6 && pwd.length <= 20) {
        }
        if (pwd == "") {
            $("#pwd-message").css("display", "block");
            document.getElementById("pwd-message").innerHTML = "*密码不能为空";
        }
    });

    $("#signup-submit").click(function () {
        if (account_flag && pwd_flag && email_flag) {
            var data = {
                username: $("#name-input").val(),
                email: $("#email-input").val(),
                password: $("#pwd-input").val()
            };
            $.ajax({
                type: "post",
                data: data,
                url: "../static/data/signup.json",
                dataType: "json",
                success: function (data) {
                    if (data.success) {
                        alert("注册成功");
                    }
                }
            });
        }
        return false;
    });

})
