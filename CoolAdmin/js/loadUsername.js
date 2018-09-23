$(document).ready(() => {
    let urlParams = new URLSearchParams(window.location.search);
    let userParam = urlParams.get('username');
    $('li a').each(function() {
        console.log(userParam);
        $(this).attr('href', $(this).attr("href") + "?username=" + userParam);    
    });
});