<!DOCTYPE html>
<html>
<head>
	<title>提示</title>
	  <link rel="stylesheet" type="text/css" href="/static/css/font-awesome.css">
  	<link rel="stylesheet" type="text/css" href="/static/css/main.css">
  	<script type="text/javascript" src="/static/js/jquery-2.2.4.js"></script>
  	<script type="text/javascript" src="/static/js/shudong.js"></script>
</head>
<body>
<div class="alert-div">
	<p>{{.Alert}}</p>
	<a href="/">回到首页</a>
	<a href="{{.Redirect}}">返回</a>
</div>
</body>
</html>