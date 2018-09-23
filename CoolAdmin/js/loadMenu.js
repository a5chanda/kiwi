$(document).ready(() => {
    $(".menu").load("menu.html", function () {
        var fileName = location.href.split("/").slice(-1)[0];
        fileName = '#' + fileName.substr(0, fileName.indexOf(".html"));
        $(fileName).addClass('active');
        $(fileName).addClass('has-sub');
    });
});