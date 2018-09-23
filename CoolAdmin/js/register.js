const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Mzc3NDIwMDUsInVzZXJuYW1lIjoiVXNlck9yZzEiLCJvcmdOYW1lIjoiT3JnMSIsImlhdCI6MTUzNzcwNjAwNX0.j_VlemeawMG79sWiqmdTwLVizLlH9QdQwgsuXLd5OIA";
$('#registerButton').click((e) => {
    e.preventDefault();
    let funcName = 'add' + $("#accountType").val();
    let options = {
        method: 'POST',
        url: 'http://localhost:4000/channels/kiwi-channel/chaincodes/mycc',
        headers: {
            'Authorization': 'Bearer ' + token,
            'Content-Type': 'application/json'
        },
        data: JSON.stringify({
            "peers": ["peer0.org1.kiwi.com","peer1.org1.kiwi.com"],
            "fcn": "addBusiness",
            "args": [$('#username').val()+"", uuidv4(),"[]","{}", String($('#equity').val())]
        })
    };
    console.log(options);
    $.ajax(options).then((res) => {
        window.open('profile.html?username='+ $('#username').val(), "_self");
    }).catch((err) => {
        console.log(err);
    })
});