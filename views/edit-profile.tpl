<div class="button-field">
	<a class="large-text-button-selected" href="/edit-profile">个人信息</a>
	<a class="large-text-button" href="/edit-password">修改密码</a>
</div>
<div class="card">
<!-- 头像 -->
<form id="avatar-form" method="POST" enctype="multipart/form-data" action="/upload-avatar">
<div class="avatar-wrapper">
	<label class="avatar-title">头像</label>
	<label class="avatar" style="background-image: url(static/www/avatar/{{.Avatar}})">
		<input type="file" name="avatar" id="avatar">
	</label>
</div>
<div class="avatar-preview">
<img id="avatar-preview-img" src="#" alt="avatar-preview" />
<input class="radius-button radius-button-green  radius-button-small" type="submit" value="上传">
<button id="cancel-button" style="margin-left: 0px;" class="radius-button radius-button-gray radius-button-small">取消</button>
</div>
</form>
<!-- 头像结束 -->
<!-- 个人信息 -->
<form id="profile-form" method="POST" enctype="multipart/form-data">
<div class="profile-field">
	<h3>邮件</h3>
	<p>{{.Email}}</p>
</div>
<div class="profile-field">
	<h3>登录名</h3>
	<p>{{.Username}}</p>
</div>
<div class="shudong-normal-input">
	<label for="phone-number">手机号</label>
	<input id="phone-number" type="text" name="phone-number" placeholder="手机号" value="{{.Phone}}">
</div>
<div class="shudong-normal-input">
	<label for="qq-number">QQ</label>
	<input id="qq-number" type="text" name="qq" placeholder="QQ号" value="{{.QQ}}">
</div>
<div class="shudong-normal-input">
	<label for="weixin-number">微信</label>
	<input id="weixin-number" type="text" name="weixin" placeholder="微信号" value="{{.Weixin}}">
</div>
<div class = "shudong-select">
	<label for="address-select">地址</label>
	<select name="address">
	<option value="未选择">未选择</option>
	<option value="华南理工大学">华南理工大学</option>
	<option value="中山大学">中山大学</option>
	</select>
</div>
<input class="radius-button radius-button-green margin-top-32" type="submit" value="提交">
</form>
<!-- 个人信息结束 -->
</div>
<script type="text/javascript">
	function readURL(input) {
		if (input.files && input.files[0]) {
			var reader = new FileReader();
			reader.onload = function (e) {
				$("#avatar-preview-img").attr("src", e.target.result);
			}

			reader.readAsDataURL(input.files[0]);
		}
	}

	$("#avatar").change(function() {
		$(".avatar").css("display", "none");
		$(".avatar-preview").css("display", "block");
		readURL(this);
	});

	$("#cancel-button").click(function() {
		$(".avatar-preview").css("display", "none");
		$(".avatar").css("display", "block");
		return false;
	})

</script>