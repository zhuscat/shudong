var parseURL = function(url) {
    var parse_url = /^(?:([A-Za-z]+):)?(\/{0,3})([0-9.\-A-Za-z]+)(?::(\d+))?(?:(\/[^?#]*))?(?:\?([^#]*))?(?:#(.*))?$/,
        result = parse_url.exec(url),
        names = ['url', 'scheme', 'slash', 'host', 'port', 'path', 'query', 'hash'],
        urlComponent = {},
        i = 0;
    for (i = 0; i < names.length; i++) {
        urlComponent[names[i]] = result[i];
    }
    return urlComponent;
}

$(function() {
    $('.send-message-box textarea').on('focus', function() {
        this.style.color = '#4a4a4a';
    });
    $('.send-message-box textarea').on('blur', function() {
        this.style.color = '#9b9b9b';
    });
    $('#broadcast-button').on('click', function() {
        $(this).addClass('trigger-menu-item-active');
        $('#notification-button').removeClass('trigger-menu-item-active');
        $('#broadcast').css('display', 'block');
        $('#notification').css('display', 'none');
    });
    $('#notification-button').on('click', function() {
        $(this).addClass('trigger-menu-item-active');
        $('#broadcast-button').removeClass('trigger-menu-item-active');
        $('#broadcast').css('display', 'none');
        $('#notification').css('display', 'block');
    });
})

$(function() {
    $('#book-management .bar-item').each(function(index, el) {
        var url = $(this).find('a')[0].href;
        var component = parseURL(url);
        var path = component.path;
        var currentURL = window.location.pathname;
        if (path === currentURL) {
            $(this).addClass('bar-item-selected')
        }
    });
})

$(function() {
    $('.delete-button').on('click', function(e) {
        var url = e.target.href,
            id = $(this).data('id');
        e.preventDefault();
        $.ajax({
            url: url,
            type: 'POST',
            dataType: 'json',
            data: {
                id: id
            }
        })
        .done(function(data) {
            if (data['success'] === true) {
                alert('操作成功');
                window.location.reload();
            } else {
                alert('操作失败');
            }
        });
    });
    $('.change-can-comment-btn').on('click', function(e) {
        var url = e.target.href,
            id = $(this).data('id');
        e.preventDefault();
        $.ajax({
            url: url,
            type: 'POST',
            dataType: 'json',
            data: {
                id: id
            }
        })
        .done(function(data) {
            if (data['success'] === true) {
                alert(data['info']);
                window.location.reload();
            } else {
                alert(data['info']);
            }
        });
    })
    $('#broadcast button').on('click', function(e) {
        e.preventDefault();
        var content = $('#broadcast').find('textarea').val();
        $.ajax({
            url: '/management/broadcast',
            type: 'POST',
            dataType: 'json',
            data: {
                content: content
            }
        })
        .done(function(data) {
            if (data['success'] === true) {
                alert('操作成功');
            } else {
                alert('操作失败');
            }
        });
    });
    $('#notification button').on('click', function(e) {
        e.preventDefault();
        var content = $('#notification').find('textarea').val(),
            username = $('#notification').find('input').val()
        $.ajax({
            url: '/management/notification',
            type: 'POST',
            dataType: 'json',
            data: {
                username: username,
                content: content
            }
        })
        .done(function(data) {
            alert(data['info']);
        });
    });
})

$(function () {
  $('[data-toggle="tooltip"]').tooltip({
    container: 'body'
  });
});

// 可以改成web socket
$(document).ready(function() {
  function haveNewMessage() {
    $.ajax({
      type: "GET",
      dataType: "json",
      url: "/message/have-new-message",
      success: function(data) {
        if (data["new"] == true) {
          $(".fa-bell").addClass("new-message");
        } else {
          $(".fa-bell").removeClass("new-message");
        }
      }
    });
  }

  setInterval(haveNewMessage, 5000);

  $(".tab-menu>li").click(function() {
      $(".tab-menu>li").toggleClass("item-selected");
      $(".tab-menu>li").each(function() {
          var item = $(this).data("item");
          $(item).css("display", "none");
      });
      var thisItem = $(this).data("item");
      $(thisItem).css("display", "block");
  });

  $(".comment-page a").click(function(e) {
      // 获取链接里面的参数
      var href = $(this).attr("href");
      var re = /\/book\/detail\/(\d+)\?p=(\d+)/gi;
      var matches = re.exec(href);
      var getUrl = "/comment/get/" + matches[1] + "?p=" + matches[2];
      $(".comment-page a").parent().removeClass("active");
      $(this).parent().addClass("active");
      $.ajax({
          url: getUrl,
          type: "GET",
          success: function(data) {
              $("#comment-display").html(data);
          }
      });
      e.preventDefault();
  });
});
