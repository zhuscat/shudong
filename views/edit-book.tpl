<div class="card">
<form id="book-form" method="POST" enctype="multipart/form-data">
<div class="book-pic-wrapper">
    <label class="book-pic-title">图片</label>
    <label class="book-pic">
        <img src="/static/www/book/{{.Book.Picture}}">
        <input type="file" name="picture" id="picture">
    </label>
</div>
<div class="shudong-normal-input">
    <label for="title">书名</label>
    <input id="title" type="text" name="title" placeholder="书名" value="{{.Book.Title}}">
</div>
<div class="shudong-normal-input">
    <label for="author">作者</label>
    <input id="author" type="text" name="author" placeholder="作者" value="{{.Book.Author}}">
</div>
<div class="shudong-normal-input">
    <label for="publisher">出版商</label>
    <input id="publisher" type="text" name="publisher" placeholder="出版商" value="{{.Book.Publisher}}">
</div>
<div class="shudong-normal-input">
    <label for="price">价格</label>
    <input id="price" type="text" name="price" placeholder="价格" value="{{.Book.Price}}">
</div>
<div class="shudong-normal-input">
    <label for="isbn">ISBN</label>
    <input id="isbn" type="text" name="isbn" placeholder="ISBN" value="{{.Book.Isbn}}">
</div>
<div class="shudong-normal-input">
    <label for="description">描述</label>
    <input id="description" type="text" name="description" placeholder="简单的描述一下..." value="{{.Book.Description}}">
</div>
<input class="radius-button radius-button-green margin-top-32" type="submit" value="修改">
</form>
<!-- 个人信息结束 -->
</div>
<script type="text/javascript">
    function readURL(input) {
        if (input.files && input.files[0]) {
            var reader = new FileReader();
            reader.onload = function (e) {
                $(".book-pic>img").attr("src", e.target.result);
            }

            reader.readAsDataURL(input.files[0]);
        }
    }

    $("#picture").change(function() {
        readURL(this);
    });
</script>
