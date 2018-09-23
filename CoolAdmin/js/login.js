$('#login').click((e) => {
    e.preventDefault();
    window.open('profile.html?username='+ $('#username').val(), "_self");
});