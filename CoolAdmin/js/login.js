$('#login').click((e) => {
    e.preventDefault();
    window.open('index.html?username='+ $('#username').val(), "_self");
});