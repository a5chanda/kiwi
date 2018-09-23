const token = "";
$('#registerButton').click((e) => {
    e.preventDefault();
    let funcName = 'add' + $("#accountType").val();
    let options = {
        method: 'POST',
        url: 'http://localhost:4200/channels/kiwi-channel/chaincodes/mycc',
        headers: {
            'Authorization': 'Bearer' + token,
            'content-type': 'application/json'
        },
        data: {
            peers: ["peer0.org1.kiwi.com","peer1.org1.kiwi.com"],
            fcn: funcName,
            args: [$('#username').val(), uuidv4(),[],{}, $('#equity').val()]
        }
    };
    console.log(options);
    $.ajax(options).then((res) => {
        window.open('profile.html?username='+ $('#username').val(), "_self");
    }).catch((err) => {
        console.log(err);
    })
});