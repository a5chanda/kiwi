$(document).ready(() => {
    let urlParams = new URLSearchParams(window.location.search);
    let userParam = urlParams.get('username');
    $('#header-username').html(userParam);
});