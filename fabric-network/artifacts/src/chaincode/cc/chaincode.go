package main

import (
	//"bytes"
	"fmt"
	"encoding/json"
	"strconv"
	"errors"
	//"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)


//******************************* STRUCTURES *************************************//

// SmartContract implements a simple chaincode to manage an asset
type SmartContract struct {
}

// Peoples IDs List
var PersonsListStr = "personList"

type PersonsList struct{ // Stores the list of quote IDs
	List []string `json:"personsList"`
}

// Businesses ID List
var BusinessesListStr = "businessList"

type BusinessesList struct{ // Stores the list of Businesses IDs
	List []string `json:"businessesList"`
}

// Services ID List
var ServicesListStr = "serviceList"

type ServicesList struct{ // Stores the list of Businesses IDs
	List []string `json:"servicesList"`
}


type Business struct{ // Business STRUCTURE
	ObjectType					string			`json:"docType"`
	BusinessName    			string			`json:businessName`	
	BusinessID			 	 	string 			`json:"businessID"`
	Employees	 				[]int			`json:"employees"`
	Inventory					map[string]int	`json:"inventory"`
	NetWorth					float64			`json:"netWorth"`
}

type Person struct{ // Person STRUCTURE
	ObjectType					string			`json:"docType"`
	Name    					string			`json:"personName"`	
    PersonID			 	 	string			`json:"businessID"`
	CompanyName	 				string			`json:"companyName"`
	NetWorth					float64			`json:"netWorth"`
	AssetsInventory 			[]int			`json:"assetsInventory"`
}

type Service struct{
	ObjectType					string			`json:"docType"`
	ServiceID					string			`json:"serviceID"`
	ServiceName    				string			`json:"serviceName"`
	Price 						int				`json:"servicePrice"`
	Description					string			`json:"description"`
	ServiceType					string			`json:"serviceType"`
	PosterID					string			`json:"posterID"`
	PostingStatus				bool   			`json:"postingStatus"`
	LastDateOfTransaction  		string			`json:"lastDateOfTransaction"`
}
//*******************************************************************************//


// ******************************* INITIALIZATION *******************************//

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(SmartContract)); err != nil {
		fmt.Printf("Error starting SmartContract chaincode: %s", err)
	}
}


// Init is called during chaincode instantiation to initialize any
// data. Note that chaincode upgrade also calls this function to reset
// or to migrate data.
func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// Invoke is called per transaction on the chaincode. Each transaction is
// either a 'get' or a 'set' on the asset created by Init function. The Set
// method may create a new asset by specifying a new key-value pair.
func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	// Extract the function and args from the transaction proposal
	fn, args := stub.GetFunctionAndParameters()

	if fn == "addService" {
		return s.addService(stub,args)
	} else if fn =="addBusiness"{
		return s.addBusiness(stub, args)	
	} else if fn =="addPerson"{
		return s.addPerson(stub, args)	
	} else if fn == "query"{ // Query by name
		return s.query(stub, args)
	} else if fn == "queryPersonsList"{ 
		return s.queryPersonsList(stub, args)
	} else if fn == "queryServicesList" { 
		return s.queryServicesList(stub, args)
	} else if fn == "queryBusinessesList" { 
		return s.queryBusinessesList(stub, args)
	} 


	// Return the result as success payload
	fmt.Println("Receieved unknown invoke function name - '" + fn + "'")
	return shim.Error("Received Unknown invoke function name - '" + fn + "'")
}

//*****************************************************************************//



//********************************************** INITIALIZER FUNCTIONS*************************************//
func stringToIntArr(args string) []int{
	str := args
    var ints []int
    err := json.Unmarshal([]byte(str), &ints)
    if err != nil {
        fmt.Println("Error")
    }
	return ints;
}

func (s *SmartContract) addPerson(stub shim.ChaincodeStubInterface, args []string) sc.Response{
	var err error

	// Name    					string			`json:"personName"`	
    // PersonID			 	 	string			`json:"businessID"`
	// CompanyName	 			string			`json:"companyName"`
	// NetWorth					float64			`json:"netWorth"`
	// AssetsInventory 			[]int			`json:"assetsInventory"`

	fmt.Println("Number of args: ")
	fmt.Println(len(args))

 //*********************************** Error handling **********************************//
	netWorth, err := strconv.ParseFloat(args[3], 64)
	if err != nil{
		return shim.Error("4th argument 'NetWorth' must be a numeric string,"+ args[3])
	}

 //*****************************************************************************************//

	//Creating the quote object
	person := Person{
		ObjectType: "Person",
		Name: args[0],
	  	PersonID: args[1],
		CompanyName: args[2],
		NetWorth: netWorth,
		AssetsInventory: stringToIntArr(args[4]),	
	}

	fmt.Println("Person being added: ")
	fmt.Println(person)


	// 1. Storing quote with ID 
	personAsBytes, err := json.Marshal(person)

	// Save quote to state with ID
	err = stub.PutState(args[0], personAsBytes) //storing by name
	if err != nil {
		return shim.Error(err.Error())
	}
	
	// 2. Append personID into personslist
	personsListAsBytes, err := stub.GetState(PersonsListStr) //getting the personsList
	if err !=nil{
		return shim.Error("Failed to get personsList")
	}

	var pList PersonsList
	json.Unmarshal(personsListAsBytes, &pList)

	pList.List = append([]string{args[0]}, pList.List...)
	fmt.Println("! appended quote to personslist")

	pListAsBytes, _ := json.Marshal(pList)
	err = stub.PutState(PersonsListStr, pListAsBytes)
	if err != nil{
		return shim.Error("Error pushing personslist into blockchain")
	}

	fmt.Println("- end init person")
	return shim.Success(nil)
}

func (s *SmartContract) addBusiness(stub shim.ChaincodeStubInterface, args []string) sc.Response{

	///********//////	
	var err error

	// ObjectType					string			`json:"docType"`
	// BusinessName    			string			`json:businessName`	
	// BusinessID			 	 	string 			`json:"businessID"`
	// Employees	 				[]int			`json:"employees"`
	// Inventory					map[string]int	`json:"inventory"`
	// NetWorth					float64			`json:"netWorth"`

	fmt.Println("Number of args: ")
	fmt.Println(len(args))
	fmt.Println(args[0]) // Business Name
	fmt.Println(args[1]) // BusinessID
	fmt.Println(args[2]) // Employees
	fmt.Println(args[3]) //Inventory
						// networth
	//*********************************Error Handling*****************************************//
	netWorth, err := strconv.ParseFloat(args[4], 64)
	if err != nil{
		return shim.Error("3rd argument 'NetWorth' must be a numeric string")
	}


	var inventory = make(map[string]int)
	inventory["init"] = -1
	//Creating the quote object
	business := Business{
		ObjectType: "Business",
		BusinessName: args[0],
		BusinessID: args[1],
	  	Employees: stringToIntArr(args[2]),
		Inventory: inventory,
		NetWorth: netWorth,	
	}

	fmt.Println("Business being added: ")
	fmt.Println(business)


	// 1. Storing quote with ID 
	businessAsBytes, err := json.Marshal(business)

	// Save quote to state
	err = stub.PutState(args[0], businessAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	
	// 2. Append personID into personslist
	businessesListAsBytes, err := stub.GetState(BusinessesListStr)
	if err !=nil{
		return shim.Error("Failed to get businessesList")
	}

	var bList BusinessesList
	json.Unmarshal(businessesListAsBytes, &bList)

	bList.List = append([]string{args[0]}, bList.List...)
	fmt.Println("! appended quote to businesseslist")

	bListAsBytes, _ := json.Marshal(bList)
	err = stub.PutState(BusinessesListStr, bListAsBytes)
	if err != nil{
		return shim.Error("Error pushing list into blockchain")
	}

	fmt.Println("- end init business")
	return shim.Success(nil)
}

func (s *SmartContract) addService(stub shim.ChaincodeStubInterface, args []string) sc.Response{
	var err error
 //*****************************************************************************************//
	// ObjectType					string			`json:"docType"`
	// ServiceID					string			`json:"serviceID"`
	// ServiceName    				string			`json:"serviceName"`
	// Price 						int				`json:"servicePrice"`
	// Description					string			`json:"description"`
	// ServiceType					string			`json:"serviceType"`
	// PosterID					string			`json:"posterID"`
	// PostingStatus				bool   			`json:"postingStatus"`
	// LastDateOfTransaction  		string			`json:"lastDateOfTransaction"`

	fmt.Println("Number of args: ")
	fmt.Println(len(args))
	fmt.Println(args[0]) // Business Name
	fmt.Println(args[1]) // BusinessID
	fmt.Println(args[2]) // Employees
	fmt.Println(args[3]) //Inventory
						// networth

	price, err := strconv.Atoi(args[2])
	if(err != nil){
		fmt.Println("Error parsing args[2], Price must be a numeric string")
	}

	postingStatus, err := strconv.ParseBool(args[6])
	if(err != nil){
		fmt.Println("Error parsing postingStatus, posting status must be a boolean string")
	}
 //*****************************************************************************************//

	//Creating the service object
	service := Service{
		ObjectType: "Business",
		ServiceID: args[0],
	  	ServiceName: args[1],
		Price: price,
		Description: args[3],
		ServiceType: args[4],
		PosterID: args[5],
		PostingStatus: postingStatus,
		LastDateOfTransaction: args[7],	
	}

	fmt.Println("Service being added: ")
	fmt.Println(service)


	// 1. Storing service with ID 
	serviceAsBytes, err := json.Marshal(service)

	// Save service to state
	err = stub.PutState(args[0], serviceAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	
	// 2. Append serviceID into servicelist
	servicesListAsBytes, err := stub.GetState(ServicesListStr)
	if err !=nil{
		return shim.Error("Failed to get servicesList")
	}

	var sList ServicesList
	json.Unmarshal(servicesListAsBytes, &sList)

	sList.List = append([]string{args[0]}, sList.List...)
	fmt.Println("! appended service ID to serviceslist")

	sListAsBytes, _ := json.Marshal(sList)
	err = stub.PutState(ServicesListStr, sListAsBytes)
	if err != nil{
		return shim.Error("Error pushing serviceslist into blockchain")
	}
	

	fmt.Println("- end init service")
	return shim.Success(nil)
}


// ********************GETTER FUNCTIONS*******************/

// ============================================================================================================================
// Get Person - get the person asset from ledger
// ============================================================================================================================
func get_person(stub shim.ChaincodeStubInterface, personID string) (Person, error) {
	var person Person
	personAsBytes, err := stub.GetState(personID)//getState retreives a key/value from the ledger

	if err != nil {                                            //this seems to always succeed, even if key didn't exist
		return person, errors.New("Failed to get person - " + personID)
	}
	json.Unmarshal(personAsBytes, &person)                       //un stringify it aka JSON.parse()

	if len(person.PersonID) == 0 {                              //test if owner is actually here or just nil
		return person, errors.New("person does not exist - " + personID + ", '" + person.Name)
	}
	
	return person, nil
}

// ============================================================================================================================
// Get Service - get the service asset from ledger
// ============================================================================================================================
func get_service(stub shim.ChaincodeStubInterface, serviceID string) (Service, error) {
	var service Service
	serviceAsBytes, err := stub.GetState(serviceID)                     //getState retreives a key/value from the ledger
	if err != nil {                                            //this seems to always succeed, even if key didn't exist
		return service, errors.New("Failed to get service - " + serviceID)
	}
	json.Unmarshal(serviceAsBytes, &service)                       //un stringify it aka JSON.parse()

	if len(service.ServiceID) == 0 {                              //test if owner is actually here or just nil
		return service, errors.New("Service does not exist - " + serviceID + ", '" + service.ServiceName)
	}
	
	return service, nil
}


// ============================================================================================================================
// Get Business - get the business asset from ledger
// ============================================================================================================================
func get_business(stub shim.ChaincodeStubInterface, businessID string) (Business, error) {
	var business Business
	businessAsBytes, err := stub.GetState(businessID)                     //getState retreives a key/value from the ledger
	if err != nil {                                            //this seems to always succeed, even if key didn't exist
		return business, errors.New("Failed to get business - " + businessID)
	}
	json.Unmarshal(businessAsBytes, &business)                       //un stringify it aka JSON.parse()

	if len(business.BusinessID) == 0 {                              //test if owner is actually here or just nil
		return business, errors.New("Business does not exist - " + businessID + ", '" + business.BusinessName)
	}
	
	return business, nil
}


//********************* QUERIES ********************//

func (s *SmartContract) query(stub shim.ChaincodeStubInterface, args []string) sc.Response{
	fmt.Println("============================Beginning Query================================")

	if len(args) != 1{
		return shim.Error("Incorrect number of arguments. Expecting Person ID, Business ID or Service ID")
	}

	var id string
	var err error

	//var person Person
	id = args[0]
	queryAsBytes, err := stub.GetState(id)
	
	if err != nil{
		return shim.Error("json error")
	}

	//json.Unmarshal(personBytes, &person)
	return shim.Success(queryAsBytes)
}

func (s *SmartContract) queryBusinessesList(stub shim.ChaincodeStubInterface, args []string) sc.Response{
	fmt.Println("============================Beginning Businesses List Query================================")

	if len(args) != 1{
		return shim.Error("Incorrect number of arguments. Expecting empty string")
	}

	var err error

	//var person Person
	businessList, err := stub.GetState(BusinessesListStr)
	
	if err != nil{
		return shim.Error("json error")
	}

	//json.Unmarshal(personBytes, &person)
	return shim.Success(businessList)
}

func (s *SmartContract) queryPersonsList(stub shim.ChaincodeStubInterface, args []string) sc.Response{
	fmt.Println("============================Beginning Persons List Query================================")

	if len(args) != 1{
		return shim.Error("Incorrect number of arguments. Expecting empty string")
	}

	var err error

	list, err := stub.GetState(PersonsListStr)
	
	if err != nil{
		return shim.Error("json error")
	}

	return shim.Success(list)
}

func (s *SmartContract) queryServicesList(stub shim.ChaincodeStubInterface, args []string) sc.Response{
	fmt.Println("============================Beginning Services List Query================================")

	if len(args) != 1{
		return shim.Error("Incorrect number of arguments. Expecting empty string")
	}

	var err error

	list, err := stub.GetState(ServicesListStr)
	
	if err != nil{
		return shim.Error("json error")
	}

	return shim.Success(list)
}



//********************************************** INITIALIZER FUNCTIONS*************************************//

// func (s *SmartContract) sellService(stub shim.ChaincodeStubInterface, args []string){ //args[0] => buyerID, args[2]=serviceID
// 	var buyerID = args[0]
// 	var sellerID = args[1]
	
// 	var service Service
// 	var buyer Person
	




	
// }