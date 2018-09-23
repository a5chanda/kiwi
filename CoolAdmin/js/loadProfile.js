$(document).ready(() => {
    let d = new Date();
    var monthNames = ["January", "February", "March", "April", "May","June","July", "August", "September", "October", "November","December"];
    $('#transaction-title').html($('#transaction-title').html() + monthNames[d.getMonth()] + " " + d.getFullYear());      
});