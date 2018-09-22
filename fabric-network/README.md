# Kiwi Fabric Network Setup



## Generate crypto material command
./bin/cryptogen generate --config=./crypto-config.yaml

## Generate genesis block 
./bin/configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block

## Generate Channel tx
./bin/configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID kiwi-channel

## Update Peer Anchors for org1 and org2


# Starting the Network

'docker-compose -f docker-compose-cli.yaml up'


# Starting the Network

'docker-compose -f docker-compose-cli.yaml up'


## Devmode

### Setting up Channel

1. Get the channel.block from the network orderer

`peer channel create -o orderer.kiwi.com:7050 -c kiwi-channel -f ./channel-artifacts/channel.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/kiwi.com/orderers/orderer.kiwi.com/msp/tlscacerts/tlsca.kiwi.com-cert.pem`

2. Join the channel
`peer channel join -b kiwi-channel.block`

3. Add the other peer from Org2 to the channel
`CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.kiwi.com/users/Admin@org2.kiwi.com/msp CORE_PEER_ADDRESS=peer0.org2.kiwi.com:7051 CORE_PEER_LOCALMSPID="Org2MSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.kiwi.com/peers/peer0.org2.kiwi.com/tls/ca.crt peer channel join -b kiwi-channel.block`

4. Add the rest of the peers by adjusting the numbers accordingly

5. Update anchor peers
    -  This command updates the anchor peer for org1.
    `peer channel update -o orderer.kiwi.com:7050 -c kiwi-channel -f ./channel-artifacts/Org1MSPanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/kiwi.com/orderers/orderer.kiwi.com/msp/tlscacerts/tlsca.kiwi.com-cert.pem`
    - The second command updates the anchor peer for org2
    `CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.kiwi.com/users/Admin@org2.kiwi.com/msp CORE_PEER_ADDRESS=peer0.org2.kiwi.com:7051 CORE_PEER_LOCALMSPID="Org2MSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.kiwi.com/peers/peer0.org2.kiwi.com/tls/ca.crt peer channel update -o orderer.kiwi.com:7050 -c kiwi-channel -f ./channel-artifacts/Org2MSPanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/kiwi.com/orderers/orderer.kiwi.com/msp/tlscacerts/tlsca.kiwi.com-cert.pem`




# Testing the Node SDK and interacting with the network

The commands are taken from the balance-transfer folder from fabric-samples and have been modified to interact with our network.

## Open a new terminal for running the node application

Navigate to the  directory that contains app.js and then start the node application on port 4000:

`PORT=4000 node app`

## Open another terminal for invoking the node commands to interact with the network

## Login Request

### Enrolling a user for Org1

The following command registers and enrolls a user with the name `User` for an Organization, in this case Org1:
`curl -s -X POST   http://localhost:4000/users   -H "content-type: application/x-www-form-urlencoded"   -d 'username=UserOrg1&orgName=Org1'`

This should Output
```
{
    "success":true,
    "secret":"",
    "message":"UserOrg1 enrolled Successfully",
    "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzA4MzY5MTEsInVzZXJuYW1lIjoiVXNlck9yZzEiLCJvcmdOYW1lIjoiT3JnMSIsImlhdCI6MTUzMDgwMDkxMX0.x5S4dSqEIGZLQKINpfz_UXvH2Un7_m1-J12zBGcMXZc"
}
```
The token value will be unique and not the same as above.

The response contains the success/failure status, an **enrollment Secret** and a **JSON Web Token (JWT)** that is a required string in the Request Headers for subsequent requests. 

### Enrolling a user for Org12

The following command registers and enrolls a user with the name `User` for an Organization, in this case Org1:
`curl -s -X POST   http://localhost:4000/users   -H "content-type: application/x-www-form-urlencoded"   -d 'username=UserOrg2&orgName=Org2'`

This should Output
```
{
    "success":true,
    "secret":"",
    "message":"UserOrg2 enrolled Successfully",
    "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzA2NjMyMTEsInVzZXJuYW1lIjoiVXNlck9yZzEiLCJvcmdOYW1lIjoiT3JnMSIsImlhdCI6MTUzMDYyNzIxMX0.sHxWIvPPX8HP7X_T1KZzCJypp3_3qkND3xA_PP43yoE"
}
```


## Channel Configuration
### Create Channel Request
```
curl -s -X POST \
  http://localhost:4000/channels \
  -H "authorization: Bearer <Token Value For User in ORG1>" \
  -H "content-type: application/json" \
  -d '{
	"channelName":"kiwi-channel",
	"channelConfigPath":"../artifacts/channel/channel-artifacts/channel.tx"
}'
```

### Request for Org1 to Join Channel 
```
curl -s -X POST \
  http://localhost:4000/channels/kiwi-channel/peers \
  -H "authorization: Bearer <Token Value For User in ORG1>" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.org1.kiwi.com","peer1.org1.kiwi.com"]
}'
```

### Request for Join Channel on Org2 (Not necessary for testing)
```
curl -s -X POST \
  http://localhost:4000/channels/kiwi-channel/peers \
  -H "authorization: Bearer <Token Value Org2>" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.org2.kiwi.com","peer1.org2.kiwi.com"]
}'
```
## Chaincode Interactions

### Install Chaincode on Org1 Peers
```
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzA4MzcxNDMsInVzZXJuYW1lIjoiVXNlck9yZzEiLCJvcmdOYW1lIjoiT3JnMSIsImlhdCI6MTUzMDgwMTE0M30.yAWEyRolt38XOOfKM_h0Y7vDH1RoFYhKu_hjzlS5eYA" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.org1.kiwi.com","peer1.org1.kiwi.com"],
	"chaincodeName":"mycc",
	"chaincodePath":"chaincode/cc",
	"chaincodeVersion":"1.0",
    "chaincodeType":"golang"
}'
```

### Install Chaincode on Org2 Peers
```
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer <Token Value UserOrg2>" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.org2.kiwi.com","peer1.org2.kiwi.com"],
	"chaincodeName":"mycc",
	"chaincodePath":"chaincode/cc",
	"chaincodeVersion":"1.0",
    "chaincodeType":"golang"
}'
```

## Instantiate Chaincode
```
curl -s -X POST \
  http://localhost:4000/channels/kiwi-channel/chaincodes \
  -H "authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzA4MzcxNDMsInVzZXJuYW1lIjoiVXNlck9yZzEiLCJvcmdOYW1lIjoiT3JnMSIsImlhdCI6MTUzMDgwMTE0M30.yAWEyRolt38XOOfKM_h0Y7vDH1RoFYhKu_hjzlS5eYA" \
  -H "content-type: application/json" \
  -d '{
    "peers": ["peer0.org1.kiwi.com","peer1.org1.kiwi.com"],
	"chaincodeName":"mycc",
	"chaincodeVersion":"1.0",
    "chaincodeType":"golang",
	"args":["init", "a"]
}'
```

## Invoke Chaincode
```
curl -s -X POST \
  http://localhost:4000/channels/kiwi-channel/chaincodes/mycc \
  -H "authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzA4MzcxNDMsInVzZXJuYW1lIjoiVXNlck9yZzEiLCJvcmdOYW1lIjoiT3JnMSIsImlhdCI6MTUzMDgwMTE0M30.yAWEyRolt38XOOfKM_h0Y7vDH1RoFYhKu_hjzlS5eYA" \
  -H "content-type: application/json" \
  -d '{
    "peers": ["peer0.org1.kiwi.com","peer1.org1.kiwi.com"],
	"fcn":"",
	"args":[]
}'
```

## Query Chaincode

curl -s -X GET \
  "http://localhost:4000/channels/kiwi-channel/chaincodes/mycc?peer=peer0.org1.kiwi.com&fcn=query&args=%5B%2220870%22%5D" \
  -H "authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzA4MzcxNDMsInVzZXJuYW1lIjoiVXNlck9yZzEiLCJvcmdOYW1lIjoiT3JnMSIsImlhdCI6MTUzMDgwMTE0M30.yAWEyRolt38XOOfKM_h0Y7vDH1RoFYhKu_hjzlS5eYA" \
  -H "content-type: application/json"


## Query Chaincode and have record of who accessed it

curl -s -X POST \
  http://localhost:4000/channels/kiwi-channel/chaincodes/mycc \
  -H "authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzA1NzkyMDksInVzZXJuYW1lIjoiVXNlck9yZzEiLCJvcmdOYW1lIjoiT3JnMSIsImlhdCI6MTUzMDU0MzIwOX0.lUYCSkW6K3fZKBUGWGNl_GVoisujU6TEKoHyaNow-0o" \
  -H "content-type: application/json" \
  -d '{
    "peers": ["peer0.org1.kiwi.com","peer1.org1.kiwi.com"],
	"fcn":"query",
	"args":[""]
}'
