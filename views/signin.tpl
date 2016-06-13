<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>登录</title>
	<link rel="stylesheet" type="text/css" href="/static/css/font-awesome.css">
	<link rel="stylesheet" type="text/css" href="/static/css/main.css">
</head>
<body id="signin">
<h1 class="shudong-header1 text-align-center">
书洞
</h1>
<h2 class="shudong-header2 text-align-center">
一个二手书交易平台
</h2>
<div class="button-field">
	<a class="large-text-button-selected" href="/signin">登录</a>
	<a class="large-text-button" href="/signup">注册</a>
</div>
<div class="card">
<form method="POST">
<div class="shudong-input">
	<div class="shudong-input-left-div">
	<i class="fa fa-user fa-lg"></i>
	</div>
	<input class="shudong-input-right-field" type="text" name="username" placeholder="用户名">
</div>
<div class="shudong-input">
	<div class="shudong-input-left-div">
		<i class="fa fa-unlock-alt fa-lg"></i>
	</div>
	<input class="shudong-input-right-field" type="password" name="password" placeholder="密码">
</div>
<input class="radius-button" type="submit" value="登录">
<a id="forget-password" href="/forgot">忘记密码?</a>
</form>
</div>
</body>
</html>