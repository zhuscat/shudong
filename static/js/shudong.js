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
