## Query Chaincode by Name

curl -s -X GET \
  "http://localhost:4000/channels/kiwi-channel/chaincodes/mycc?peer=peer0.org1.kiwi.com&fcn=query&args=%5B%22*NAMEGOESHERE*%22%5D" \
  -H "authorization: Bearer <Bearer Token>" \
  -H "content-type: application/json"

## Query Chaincode for List of businesses

curl -s -X GET \
  "http://localhost:4000/channels/kiwi-channel/chaincodes/mycc?peer=peer0.org1.kiwi.com&fcn=queryBusinessesList&args=%5B%22%22%5D" \
  -H "authorization: Bearer <Bearer Token>" \
  -H "content-type: application/json"

## Query Chaincode for List of People

curl -s -X GET \
  "http://localhost:4000/channels/kiwi-channel/chaincodes/mycc?peer=peer0.org1.kiwi.com&fcn=queryPersonsList&args=%5B%22%22%5D" \
  -H "authorization: Bearer <Bearer Token>" \
  -H "content-type: application/json"


## Query Chaincode for List of Services

curl -s -X GET \
  "http://localhost:4000/channels/kiwi-channel/chaincodes/mycc?peer=peer0.org1.kiwi.com&fcn=queryServicesList&args=%5B%22%22%5D" \
  -H "authorization: Bearer <Bearer Token>" \
  -H "content-type: application/json"