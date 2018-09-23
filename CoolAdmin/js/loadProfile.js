const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Mzc3NDIwMDUsInVzZXJuYW1lIjoiVXNlck9yZzEiLCJvcmdOYW1lIjoiT3JnMSIsImlhdCI6MTUzNzcwNjAwNX0.j_VlemeawMG79sWiqmdTwLVizLlH9QdQwgsuXLd5OIA"
$(document).ready(() => {
    let d = new Date();
    var monthNames = ["January", "February", "March", "April", "May","June","July", "August", "September", "October", "November","December"];
    $('#transaction-title').html($('#transaction-title').html() + monthNames[d.getMonth()] + " " + d.getFullYear());   
    let urlParams = new URLSearchParams(window.location.search);
    let userParam = urlParams.get('username');
    console.log(userParam);
    var qURL = 'http://localhost:4000/channels/kiwi-channel/chaincodes/mycc?peer=peer0.org1.kiwi.com&fcn=query&args=%5B%22asd%22%5D';
    console.log(qURL);
    $.ajax({
        method: 'GET',
        uri: qURL,
        headers: {
            'content-type': 'application.json',
            'Authorization': 'Bearer ' + token
        }
    }).then((res) => {
        console.log(res);
        $('#accountBalance').html(res.netWorth);    
    });
});