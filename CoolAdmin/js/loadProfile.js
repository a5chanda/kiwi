const token = "";
$(document).ready(() => {
    
    let d = new Date();
    var monthNames = ["January", "February", "March", "April", "May","June","July", "August", "September", "October", "November","December"];
    $('#transaction-title').html($('#transaction-title').html() + monthNames[d.getMonth()] + " " + d.getFullYear());   
    let urlParams = new URLSearchParams(window.location.search);
    let userParam = urlParams.get('username');
    $.ajax({
        method: 'GET',
        url: 'ttp://localhost:4200/channels/kiwi-channel/chaincodes/mycc?peer=peer0.org1.kiwi.com&fcn=query&args=%5B%22*' + userParam + '*%22%5D',
        headers: {
            'content-type': 'application.json',
            'Authorization': 'Bearer ' + token
        }
    }).then((res) => {
        $('#accountBalance').html(res.value.netWorth);    
    });
});