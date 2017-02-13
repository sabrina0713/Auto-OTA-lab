<<<<<<< HEAD
/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
	"errors"
	"fmt"
	"strconv"
	"encoding/json"
	"time"
	"github.com/hyperledger/fabric/core/chaincode/shim"

)

// Test comment
// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// Maximum number of transactions to return
const NUM_TX_TO_RETURN = 27
const NUM_OP_TO_RETURN = 27

// Smart Contract Id, Compatibility Reference numbers
const TRAVEL_CONTRACT   = "Paris"
const FEEDBACK_CONTRACT = "Feedback"
const REFERENCE_TABLES = "Compatibility"

// Blockchain point transaction record
type Transaction struct {
	RefNumber   string   `json:"RefNumber"`
	Date 		time.Time   `json:"Date"`
	Description string   `json:"description"`
	Type 		string   `json:"Type"`
	Amount    	float64  `json:"Amount"`
	Money    	float64  `json:"Money"`
	Activities  int      `json:"FeedbackActivitiesDone"`
	To			string   `json:"ToUserid"`
	From		string   `json:"FromUserid"`
	ToName	    string   `json:"ToName"`
	FromName	string   `json:"FromName"`
	ContractId	string   `json:"ContractId"`
	StatusCode	int 	 `json:"StatusCode"`
	StatusMsg	string   `json:"StatusMsg"`
}

// Blockchain operation record
type Operation struct {
	RefNumber	string		`json:"RefNumber"`
	Date 		time.Time	`json:"Date"`
	From		string		`json:"FromUserid"`
	FromName	string		`json:"FromName"`
	To		string		`json:"ToUserid"`
	ToName		string		`json:"ToName"`
	Description	string		`json:"description"`
	Version		string		`json:"Version"`
	StatusCode	int		`json:"StatusCode"`
	StatusMsg	string		`json:"StatusMsg"`
}

// Smart contract metadata record
type Contract struct {
	Id			string   `json:"ID"`
	BusinessId  string   `json:"BusinessId"`
	BusinessName string   `json:"BusinessName"`
	Title		string   `json:"Title"`
	Description string   `json:"Description"`
	Conditions  []string `json:"Conditions"`
	Icon        string 	 `json:"Icon"`
	StartDate   time.Time   `json:"StartDate"`
	EndDate		time.Time   `json:"EndDate"`
	Method	    string   `json:"Method"`
	DiscountRate float64  `json:"DiscountRate"`
}

type Compatibility struct {
	Chassis		[]string	`json:"Chassis"`
	Powertrain	[]string	`json:"Powertrain"`
	Safety		[]string	`json:"Safety"`
	Telematics	[]string	`json:"Telematics"`
}

// Subsystem running version record
type Subsystem struct {
	SsId		string		`json:"SsId"`
	Name   		string		`json:"Name"`
	osVersion	string		`json:"osVersion"`
	NumTxs		int		`json:"NumberOfTransactions"`
	Effectivity	string		`json:"EffectivityDate"`
	Chassis		[]string	`json:"Chassis"`
	Powertrain	[]string	`json:"Powertrain"`
	Safety		[]string	`json:"Safety"`
	Telematics	[]string	`json:"Telematics"`
	Manufactured	string		`json:"MfgDate"`
	Modified	string		`json:"LastModifiedDate"`
}

// Open Points member record
type User struct {
	UserId		string   `json:"UserId"`
	Name   		string   `json:"Name"`
	Balance 	float64  `json:"Balance"`
	NumTxs 	    int      `json:"NumberOfTransactions"`
	Status      string 	 `json:"Status"`
	Expiration  string   `json:"ExpirationDate"`
	Join		string   `json:"JoinDate"`
	Modified	string   `json:"LastModifiedDate"`
}


// Array for storing all open points transactions
type AllTransactions struct{
	Transactions []Transaction `json:"transactions"`
}

// Array for storing all operations
type AllOperations struct{
	Operations []Operation `json:"operations"`
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// ============================================================================================================================
// Init - reset all the things
// ============================================================================================================================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	
	// Create the 'Bank' user and add it to the blockchain
	var bank User
	bank.UserId = "B1928564";
	bank.Name = "OpenFN"
	bank.Balance = 1000000
	bank.Status  = "Originator"
	bank.Expiration = "2099-12-31"
	bank.Join  = "2015-01-01"
	bank.Modified = "2016-05-06"
	bank.NumTxs  = 0
	
	
	jsonAsBytes, _ := json.Marshal(bank)
	err = stub.PutState(bank.UserId, jsonAsBytes)								
	if err != nil {
		fmt.Println("Error Creating Bank user account")
		return nil, err
	}
	
	
    // Create the 'Travel Agency' user and add it to the blockchain
	var travel User
	travel.UserId = "T5940872";
	travel.Name = "Open Travel"
	travel.Balance = 500000
	travel.Status  = "Member"
	travel.Expiration = "2099-12-31"
	travel.Join  = "2015-01-01"
	travel.Modified = "2016-05-06"
	travel.NumTxs  = 0
	
	jsonAsBytes, _ = json.Marshal(travel)
	err = stub.PutState(travel.UserId, jsonAsBytes)								
	if err != nil {
		fmt.Println("Error Creating Travel user account")
		return nil, err
	}
	
	
	// Create the 'Natalie' user and add her to the blockchain
	var natalie User
	natalie.UserId = "U2974034";
	natalie.Name = "Natalie"
	natalie.Balance = 1000
	natalie.Status  = "Platinum"
	natalie.Expiration = "2017-06-01"
	natalie.Join  = "2015-05-31"
	natalie.Modified = "2016-05-06"
	natalie.NumTxs  = 0
	
	jsonAsBytes, _ = json.Marshal(natalie)
	err = stub.PutState(natalie.UserId, jsonAsBytes)								
	if err != nil {
		fmt.Println("Error Creating Natalie user account")
		return nil, err
	}
	
	
	// Create the 'Anthony' user and add him to the blockchain
	var anthony User
	anthony.UserId = "U3151672";
	anthony.Name = "Anthony"
	anthony.Balance = 50000
	anthony.Status  = "Silver"
	anthony.Expiration = "2017-03-15"
	anthony.Join  = "2015-08-15"
	anthony.Modified = "2016-04-17"
	anthony.NumTxs  = 0
	
	jsonAsBytes, _ = json.Marshal(anthony)
	err = stub.PutState(anthony.UserId, jsonAsBytes)								
	if err != nil {
		fmt.Println("Error Creating Anthony user account")
		return nil, err
	}
	
	
	// Create an array for storing all transactions, and store the array on the blockchain
	var transactions AllTransactions
	jsonAsBytes, _ = json.Marshal(transactions)
	err = stub.PutState("allTx", jsonAsBytes)
	if err != nil {
		return nil, err
	}
	
	// Create transaction reference number and store it on the blockchain
	var refNumber int
	
	refNumber = 2985674978
	jsonAsBytes, _ = json.Marshal(refNumber)
	err = stub.PutState("refNumber", jsonAsBytes)								
	if err != nil {
		fmt.Println("Error Creating reference number")
		return nil, err
	}

	
	// Create contract metadata for double points and add it to the blockchain
	var double Contract
	double.Id = TRAVEL_CONTRACT
	double.BusinessId  = "T5940872"
	double.BusinessName = "Open Travel"
	double.Title = "Paris for Less"
	double.Description = "All Paris travel activities are half the stated point price"
	double.Conditions = append(double.Conditions, "Half off dining and travel activities in Paris")
	double.Conditions = append(double.Conditions, "Valid from May 11, 2016") 
	double.Icon = ""
	double.Method = "travelContract"
	
	startDate, _  := time.Parse(time.RFC822, "11 May 16 12:00 UTC")
	double.StartDate = startDate
	endDate, _  := time.Parse(time.RFC822, "31 Dec 60 11:59 UTC")
	double.EndDate = endDate
	
	jsonAsBytes, _ = json.Marshal(double)
	err = stub.PutState(TRAVEL_CONTRACT, jsonAsBytes)								
	if err != nil {
		fmt.Println("Error creating double contract")
		return nil, err
	}
	
	
	// Create contract metadata for feedback points and add it to the blockchain
    var feedback Contract
	feedback.Id = FEEDBACK_CONTRACT
	feedback.BusinessId  = "T5940872"
	feedback.BusinessName = "Open Travel"
	feedback.Title = "Points for Feedback"
	feedback.Description = "Earn points by sharing your thoughts on travel package and activities"
	feedback.Conditions = append(feedback.Conditions, "1,000 points for travel package ")
	feedback.Conditions = append(feedback.Conditions, "Valid from May 24, 2016")
	feedback.Icon = ""
	feedback.Method = "feedbackContract"
	startDate, _  = time.Parse(time.RFC822, "24 May 16 12:00 UTC")
	feedback.StartDate = startDate
	endDate, _  = time.Parse(time.RFC822, "31 Dec 60 11:59 UTC")
	feedback.EndDate = endDate
	
	jsonAsBytes, _ = json.Marshal(feedback)
	err = stub.PutState(FEEDBACK_CONTRACT, jsonAsBytes)								
	if err != nil {
		fmt.Println("Error creating feedback contract")
		return nil, err
	}

	// Create the 'Chassis' subsystem and add it to the blockchain
	var chassis Subsystem
	chassis.SsId = "Chassis"
	chassis.Name = "Chassis Subsystem"
	chassis.osVersion = "001"
	chassis.NumTxs  = 0
	chassis.Powertrain = append(chassis.Powertrain,"120")
	chassis.Powertrain = append(chassis.Powertrain,"131")
	chassis.Powertrain = append(chassis.Powertrain,"175")
	chassis.Powertrain = append(chassis.Powertrain,"192")
	chassis.Safety = append(chassis.Safety,"137")
	chassis.Safety = append(chassis.Safety,"139")
	chassis.Safety = append(chassis.Safety,"140")
	chassis.Telematics = append(chassis.Telematics,"036")
	chassis.Telematics = append(chassis.Telematics,"047")
	chassis.Telematics = append(chassis.Telematics,"091")

	jsonAsBytes, _ = json.Marshal(chassis)
	err = stub.PutState(chassis.SsId, jsonAsBytes)								
	if err != nil {
		fmt.Println("Error Creating Chassis user account")
		return nil, err
	}

	// Create the 'Powertrain' subsystem and add it to the blockchain
	var powertrain Subsystem
	powertrain.SsId = "Powertrain"
	powertrain.Name = "Powertrain Subsystem"
	powertrain.osVersion = "001"
	powertrain.NumTxs  = 0
	powertrain.Chassis = append(powertrain.Chassis,"023")
	powertrain.Chassis = append(powertrain.Chassis,"024")
	powertrain.Chassis = append(powertrain.Chassis,"025")
	powertrain.Chassis = append(powertrain.Chassis,"031")
	powertrain.Safety = append(powertrain.Safety,"137")
	powertrain.Safety = append(powertrain.Safety,"139")
	powertrain.Safety = append(powertrain.Safety,"140")
	powertrain.Telematics = append(powertrain.Telematics,"036")
	powertrain.Telematics = append(powertrain.Telematics,"047")
	powertrain.Telematics = append(powertrain.Telematics,"091")
	
	jsonAsBytes, _ = json.Marshal(powertrain)
	err = stub.PutState(powertrain.SsId, jsonAsBytes)								
	if err != nil {
		fmt.Println("Error Creating Powertrain user account")
		return nil, err
	}

	// Create the 'Safety' subsystem and add it to the blockchain
	var safety Subsystem
	safety.SsId = "Safety"
	safety.Name = "Safety Subsystem"
	safety.osVersion = "001"
	safety.NumTxs  = 0
	safety.Chassis = append(safety.Chassis,"023")
	safety.Chassis = append(safety.Chassis,"024")
	safety.Chassis = append(safety.Chassis,"025")
	safety.Chassis = append(safety.Chassis,"031")
	safety.Powertrain = append(safety.Powertrain,"120")
	safety.Powertrain = append(safety.Powertrain,"131")
	safety.Powertrain = append(safety.Powertrain,"175")
	safety.Powertrain = append(safety.Powertrain,"192")
	safety.Telematics = append(safety.Telematics,"036")
	safety.Telematics = append(safety.Telematics,"047")
	safety.Telematics = append(safety.Telematics,"091")

	jsonAsBytes, _ = json.Marshal(safety)
	err = stub.PutState(safety.SsId, jsonAsBytes)								
	if err != nil {
		fmt.Println("Error Creating Safety user account")
		return nil, err
	}

	// Create the 'Telematics' subsystem and add it to the blockchain
	var telematics Subsystem
	telematics.SsId = "Telematics"
	telematics.Name = "Telematics Subsystem"
	telematics.osVersion = "001"
	telematics.NumTxs  = 0
	telematics.Chassis = append(telematics.Chassis,"023")
	telematics.Chassis = append(telematics.Chassis,"024")
	telematics.Chassis = append(telematics.Chassis,"025")
	telematics.Chassis = append(telematics.Chassis,"031")
	telematics.Powertrain = append(telematics.Powertrain,"120")
	telematics.Powertrain = append(telematics.Powertrain,"131")
	telematics.Powertrain = append(telematics.Powertrain,"175")
	telematics.Powertrain = append(telematics.Powertrain,"192")
	telematics.Safety = append(telematics.Safety,"137")
	telematics.Safety = append(telematics.Safety,"139")
	telematics.Safety = append(telematics.Safety,"140")

	jsonAsBytes, _ = json.Marshal(telematics)
	err = stub.PutState(telematics.SsId, jsonAsBytes)								
	if err != nil {
		fmt.Println("Error Creating Telematics user account")
		return nil, err
	}
	// Create compatibility reference data and add it to the blockchain
	var reference Compatibility
		reference.Chassis = append(reference.Chassis,"023")
		reference.Chassis = append(reference.Chassis,"024")
		reference.Chassis = append(reference.Chassis,"025")
		reference.Chassis = append(reference.Chassis,"031")
		reference.Powertrain = append(reference.Powertrain,"120")
		reference.Powertrain = append(reference.Powertrain,"131")
		reference.Powertrain = append(reference.Powertrain,"175")
		reference.Powertrain = append(reference.Powertrain,"192")
		reference.Safety = append(reference.Safety,"137")
		reference.Safety = append(reference.Safety,"139")
		reference.Safety = append(reference.Safety,"140")
		reference.Telematics = append(reference.Telematics,"036")
		reference.Telematics = append(reference.Telematics,"047")
		reference.Telematics = append(reference.Telematics,"091")
	
	jsonAsBytes, _ = json.Marshal(reference)
	err = stub.PutState(REFERENCE_TABLES, jsonAsBytes)								
	if err != nil {
		fmt.Println("Error creating reference tables")
		return nil, err
	}

	// Create an array of contract ids to keep track of all contracts
	var contractIds []string
	contractIds = append(contractIds, TRAVEL_CONTRACT);
	contractIds = append(contractIds, FEEDBACK_CONTRACT);
	//contractIds = append(contractIds, REFERENCE_TABLES);
	
	jsonAsBytes, _ = json.Marshal(contractIds)
	err = stub.PutState("contractIds", jsonAsBytes)								
	if err != nil {
		fmt.Println("Error storing contract Ids on blockchain")
		return nil, err
	}
	
	return nil, nil
}

// ============================================================================================================================
// Run - Our entry point for Invocations - [LEGACY] obc-peer 4/25/2016
// ============================================================================================================================
func (t *SimpleChaincode) Run(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("run is running " + function)
	return t.Invoke(stub, function, args)
}

// ============================================================================================================================
// Invoke - Our entry point for Invocations
// ============================================================================================================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)


	
	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	} else if function == "transferPoints" {											//create a transaction
		return t.transferPoints(stub, args)
	} else if function == "updateEmbedded" {											//create a transaction
		return t.updateEmbedded(stub, args)
	} else if function == "addSmartContract" {											//create a transaction
		return t.addSmartContract(stub, args)
	} else if function == "incrementReferenceNumber" {											//create a transaction
		return t.incrementReferenceNumber(stub, args)
	} 
		
	fmt.Println("invoke did not find func: " + function)					//error

	return nil, errors.New("Received unknown function invocation")
}

// ============================================================================================================================
// Query - Our entry point for Queries
// ============================================================================================================================
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	
	if function == "getTxs" { return t.getTxs(stub, args[1]) }
	if function == "getOps" { return t.getOps(stub, args[1]) }
	if function == "getUserAccount" { return t.getUserAccount(stub, args[1]) }
	if function == "getSubsystem" { return t.getSubsystem(stub, args[1]) }
	if function == "getAllContracts" { return t.getAllContracts(stub) }
	if function == "getReferenceNumber" { return t.getReferenceNumber(stub) }
	if function == "getReferenceTables" { return t.getReferenceTables(stub) }
	
	fmt.Println("query did not find func: " + function)						//error

	return nil, errors.New("Received unknown function query")
}

// ============================================================================================================================
// Get Open Points member account from the blockchain
// ============================================================================================================================
func (t *SimpleChaincode) getUserAccount(stub shim.ChaincodeStubInterface, userId string)([]byte, error){
	
	fmt.Println("Start getUserAccount")
	fmt.Println("Looking for user with ID " + userId);

	//get the User index
	fdAsBytes, err := stub.GetState(userId)
	if err != nil {
		return nil, errors.New("Failed to get user account from blockchain")
	}

	return fdAsBytes, nil
	
}

// ============================================================================================================================
// Get Subsystem 'account' from the blockchain
// ============================================================================================================================
func (t *SimpleChaincode) getSubsystem(stub shim.ChaincodeStubInterface, SsId string)([]byte, error){
	
	fmt.Println("Start getSubsystem()")
	fmt.Println("Looking for subsystem with ID " + SsId);

	//get the User index
	fdAsBytes, err := stub.GetState(SsId)
	if err != nil {
		return nil, errors.New("Failed to get subsystem from blockchain")
	}

	return fdAsBytes, nil
	
}

// ============================================================================================================================
// Get all transactions that involve a particular user
// ============================================================================================================================
func (t *SimpleChaincode) getTxs(stub shim.ChaincodeStubInterface, userId string)([]byte, error){
	
	var res AllTransactions

	fmt.Println("Start find getTransactions")
	fmt.Println("Looking for " + userId);

	//get the AllTransactions index
	allTxAsBytes, err := stub.GetState("allTx")
	if err != nil {
		return nil, errors.New("Failed to get all Transactions")
	}

	var txs AllTransactions
	json.Unmarshal(allTxAsBytes, &txs)
	numTxs := len(txs.Transactions)

	for i := numTxs -1; i >= 0; i-- {
	    if txs.Transactions[i].From == userId{
			res.Transactions = append(res.Transactions, txs.Transactions[i])
		}

		if txs.Transactions[i].To == userId{
			res.Transactions = append(res.Transactions, txs.Transactions[i])
		}
		
		if (len(res.Transactions) >= NUM_TX_TO_RETURN) { break }
	}

	resAsBytes, _ := json.Marshal(res)

	return resAsBytes, nil
	
}

// ============================================================================================================================
// Get all operations that involve a particular subsystem
// ============================================================================================================================
func (t *SimpleChaincode) getOps(stub shim.ChaincodeStubInterface, SsId string)([]byte, error){
	
	var res AllOperations

	fmt.Println("Start getOps")
	fmt.Println("Looking for " + SsId);

	//get the AllOperations index
	allOpAsBytes, err := stub.GetState("allOp")
	if err != nil {
		return nil, errors.New("Failed to get all Operations")
	}

	var ops AllOperations
	json.Unmarshal(allOpAsBytes, &ops)
	numOps := len(ops.Operations)

	for i := numOps -1; i >= 0; i-- {
	    if ops.Operations[i].From == SsId{
			res.Operations = append(res.Operations, ops.Operations[i])
		}

		if ops.Operations[i].To == SsId{
			res.Operations = append(res.Operations, ops.Operations[i])
		}
		
		if (len(res.Operations) >= NUM_OP_TO_RETURN) { break }
	}

	resAsBytes, _ := json.Marshal(res)

	return resAsBytes, nil
	
}

// ============================================================================================================================
// Get the smart contract metadata from the blockchain
// ============================================================================================================================
func (t *SimpleChaincode) getAllContracts(stub shim.ChaincodeStubInterface)([]byte, error)  {

	contractIdsAsBytes, _ := stub.GetState("contractIds")
	var contractIds []string
	json.Unmarshal(contractIdsAsBytes, &contractIds)
	
	var allContracts []Contract
	for i := range contractIds{
		contractAsBytes, _ := stub.GetState(contractIds[i])
		var thisContract Contract
		json.Unmarshal(contractAsBytes, &thisContract)
		allContracts = append(allContracts, thisContract)
	}

	asBytes, _ := json.Marshal(allContracts)
	return asBytes, nil

}

func (t *SimpleChaincode) getReferenceTables(stub shim.ChaincodeStubInterface)([]byte, error)  {

	refTablesAsBytes, _ := stub.GetState(REFERENCE_TABLES)
	return refTablesAsBytes, nil

}

func (t *SimpleChaincode) incrementReferenceNumber(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {


	var refNumber int
	refNumberBytes, numErr := stub.GetState("refNumber")
	if numErr != nil {
		fmt.Println("Error Getting  ref number")
		return nil, numErr
	}
	
	json.Unmarshal(refNumberBytes, &refNumber)
	refNumber = refNumber + 1;
	refNumberBytes, _ = json.Marshal(refNumber)
	err := stub.PutState("refNumber", refNumberBytes)								
	if err != nil {
		fmt.Println("Error Creating updating ref number")
		return nil, err
	}

	return nil, nil
}

func (t *SimpleChaincode) getReferenceNumber(stub shim.ChaincodeStubInterface)([]byte, error)  {

	refNumberBytes, numErr := stub.GetState("refNumber")
	if numErr != nil {
		fmt.Println("Error Getting  ref number")
		return nil, numErr
	}
	
	return refNumberBytes, nil

}

// ============================================================================================================================
// Smart contract for giving user double points
// ============================================================================================================================
func travelContract(tx Transaction, stub shim.ChaincodeStubInterface) float64 {


	contractAsBytes, err := stub.GetState(TRAVEL_CONTRACT)
	if err != nil {
		return -99
	}
	var contract Contract
	json.Unmarshal(contractAsBytes, &contract)
	
	var pointsToTransfer float64
	pointsToTransfer = tx.Amount
	if (tx.Date.After(contract.StartDate) && tx.Date.Before(contract.EndDate)) {
	     pointsToTransfer = pointsToTransfer * 0.5
	}
 
 
  return pointsToTransfer
  
  
}


// ============================================================================================================================
// Smart contract for giving user points for completing feedback surveys
// ============================================================================================================================
func feedbackContract(tx Transaction, stub shim.ChaincodeStubInterface) float64 {
  

	contractAsBytes, err := stub.GetState(FEEDBACK_CONTRACT)
	if err != nil {
		return -99
	}
	var contract Contract
	json.Unmarshal(contractAsBytes, &contract)
	
	var pointsToTransfer float64
	pointsToTransfer = 0
	if (tx.Date.After(contract.StartDate) && tx.Date.Before(contract.EndDate)) {
	     pointsToTransfer = 1000
		 
		 if (tx.Activities > 0) {
			pointsToTransfer = pointsToTransfer + float64(tx.Activities)*100
		 }
	}
  
  
  return pointsToTransfer
  
  
}

func (t *SimpleChaincode) addSmartContract(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {


	// Create new smart contract based on user input
	var smartContract Contract
	
	discountRate, err := strconv.ParseFloat(args[4], 64)
	if err != nil {
		smartContract.Title= "Invalid Contract"
	}else{
		smartContract.DiscountRate = discountRate
	}
	
	
	smartContract.Id = args[0]
	smartContract.BusinessId  = "T5940872"
	smartContract.BusinessName = "Open Travel"
	smartContract.Title = args[1]
	smartContract.Description = ""
	smartContract.Conditions = append(smartContract.Conditions, args[2])
	smartContract.Conditions = append(smartContract.Conditions, args[3]) 
	smartContract.Icon = ""
	smartContract.Method = "travelContract"
	
	
	jsonAsBytes, _ := json.Marshal(smartContract)
	err = stub.PutState(smartContract.Id, jsonAsBytes)								
	if err != nil {
		fmt.Println("Error adding new smart contract")
		return nil, err
	}

	contractIdsAsBytes, _ := stub.GetState("contractIds")
	var contractIds []string
	json.Unmarshal(contractIdsAsBytes, &contractIds)
	
	
	var contractIdFound bool
	contractIdFound = false;
	for i := range contractIds{
		if (contractIds[i] == smartContract.Id)  {
			contractIdFound = true;
		}
	}
	
	if (!contractIdFound) {
		contractIds = append(contractIds, smartContract.Id);
	}
	
	
	jsonAsBytes, _ = json.Marshal(contractIds)
	err = stub.PutState("contractIds", jsonAsBytes)								
	if err != nil {
		fmt.Println("Error storing contract Ids on blockchain")
		return nil, err
	}

	return nil, nil

}


// ============================================================================================================================
// Transfer points between members of the Open Points Network
// ============================================================================================================================
func (t *SimpleChaincode) transferPoints(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	fmt.Println("Running transferPoints")
	currentDateStr := time.Now().Format(time.RFC822)
	startDate, _  := time.Parse(time.RFC822, currentDateStr)

	
	var tx Transaction
	tx.Date 		= startDate
	tx.To 			= args[0]
	tx.From 		= args[1]
	tx.Type 	    = args[2]
	tx.Description 	= args[3]
	tx.ContractId 	= args[4]
	activities, _  := strconv.Atoi(args[5])
	tx.Activities   = activities
	tx.StatusCode 	= 1
	tx.StatusMsg 	= "Transaction Completed"
	
	
	amountValue, err := strconv.ParseFloat(args[6], 64)
	if err != nil {
		tx.StatusCode = 0
		tx.StatusMsg = "Invalid Amount"
	}else{
		tx.Amount = amountValue
	}
	
	moneyValue, err := strconv.ParseFloat(args[7], 64)
	if err != nil {
		tx.StatusCode = 0
		tx.StatusMsg = "Invalid Amount"
	}else{
		tx.Money = moneyValue
	}
	
	
	// Get the current reference number and update it
	var refNumber int
	refNumberBytes, numErr := stub.GetState("refNumber")
	if numErr != nil {
		fmt.Println("Error Getting  ref number for transferring points")
		return nil, err
	}
	
	json.Unmarshal(refNumberBytes, &refNumber)
	tx.RefNumber 	= strconv.Itoa(refNumber)
	refNumber = refNumber + 1;
	refNumberBytes, _ = json.Marshal(refNumber)
	err = stub.PutState("refNumber", refNumberBytes)								
	if err != nil {
		fmt.Println("Error Creating updating ref number")
		return nil, err
	}
	
	// Determine point amount to transfer based on contract type
	if (tx.ContractId == TRAVEL_CONTRACT) {
		tx.Amount = travelContract(tx, stub)
	} else if (tx.ContractId == FEEDBACK_CONTRACT) {
		tx.Amount = feedbackContract(tx, stub)
	}  else {
	
		contractIdsAsBytes, _ := stub.GetState("contractIds")
		var contractIds []string
		json.Unmarshal(contractIdsAsBytes, &contractIds)
	
		for i := range contractIds{
			contractAsBytes, _ := stub.GetState(contractIds[i])
			var thisContract Contract
			json.Unmarshal(contractAsBytes, &thisContract)
			
			if (tx.ContractId == thisContract.Id) {
				tx.Amount = tx.Amount -  (tx.Amount * thisContract.DiscountRate);
			}
		}
	}
	
	
	// Get Receiver account from BC and update point balance
	rfidBytes, err := stub.GetState(tx.To)
	if err != nil {
		return nil, errors.New("transferPoints Failed to get Receiver from BC")
	}
	var receiver User
	fmt.Println("transferPoints Unmarshalling User Struct");
	err = json.Unmarshal(rfidBytes, &receiver)
	receiver.Balance = receiver.Balance  + tx.Amount
	receiver.Modified = currentDateStr
	receiver.NumTxs = receiver.NumTxs + 1
	tx.ToName = receiver.Name;
	
	
	//Commit Receiver to ledger
	fmt.Println("transferPoints Commit Updated receiver To Ledger");
	txsAsBytes, _ := json.Marshal(receiver)
	err = stub.PutState(tx.To, txsAsBytes)	
	if err != nil {
		return nil, err
	}
	
	// Get Sender account from BC nd update point balance
	rfidBytes, err = stub.GetState(tx.From)
	if err != nil {
		return nil, errors.New("transferPoints Failed to get Financial Institution")
	}
	var sender User
	fmt.Println("transferPoints Unmarshalling Sender");
	err = json.Unmarshal(rfidBytes, &sender)
	sender.Balance   = sender.Balance  - tx.Amount
	sender.Modified = currentDateStr
	sender.NumTxs = sender.NumTxs + 1
	tx.FromName = sender.Name;
	
	//Commit Sender to ledger
	fmt.Println("transferPoints Commit Updated Sender To Ledger");
	txsAsBytes, _ = json.Marshal(sender)
	err = stub.PutState(tx.From, txsAsBytes)	
	if err != nil {
		return nil, err
	}
	
	
	//get the AllTransactions index
	allTxAsBytes, err := stub.GetState("allTx")
	if err != nil {
		return nil, errors.New("transferPoints: Failed to get all Transactions")
	}

	//Update transactions arrary and commit to BC
	fmt.Println("SubmitTx Commit Transaction To Ledger");
	var txs AllTransactions
	json.Unmarshal(allTxAsBytes, &txs)
	txs.Transactions = append(txs.Transactions, tx)
	txsAsBytes, _ = json.Marshal(txs)
	err = stub.PutState("allTx", txsAsBytes)	
	if err != nil {
		return nil, err
	}
	
	
	return nil, nil

}

// ============================================================================================================================
// Update subsystem with new embedded operating software
// ============================================================================================================================
func (t *SimpleChaincode) updateEmbedded(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var err error

	fmt.Println("Running updateEmbedded()")
	currentDateStr := time.Now().Format(time.RFC822)
	startDate, _  := time.Parse(time.RFC822, currentDateStr)

	
	var op Operation
	op.Date		= startDate
	op.To 		= args[0]
	op.From 	= args[1]
	op.Description 	= args[2]
	op.Version	= args[3]
	op.StatusCode 	= 1
	op.StatusMsg 	= "Operation accepted"
	
	// Get the current reference number and update it
	var refNumber int
	refNumberBytes, numErr := stub.GetState("refNumber")
	if numErr != nil {
		fmt.Println("Error Getting  ref number for operation")
		return nil, err
	}
	
	json.Unmarshal(refNumberBytes, &refNumber)
	op.RefNumber 	= strconv.Itoa(refNumber)
	refNumber = refNumber + 1;
	refNumberBytes, _ = json.Marshal(refNumber)
	err = stub.PutState("refNumber", refNumberBytes)								
	if err != nil {
		fmt.Println("Error updating ref number")
		return nil, err
	}
	
	// Get Receiver account from BC and update point balance
	rfidBytes, err := stub.GetState(op.To)
	if err != nil {
		return nil, errors.New("updateEmbedded failed to get receiver from BC")
	}

	var receiver Subsystem
	fmt.Println("updateEmbedded unmarshalling Subsystem struct");
	err = json.Unmarshal(rfidBytes, &receiver)
	receiver.osVersion = op.Version
	receiver.Modified = currentDateStr
	receiver.NumTxs = receiver.NumTxs + 1
	op.ToName = receiver.Name;
	
	
	//Commit Receiver to ledger
	fmt.Println("updateEmbedded commit updated receiver To ledger");
	opsAsBytes, _ := json.Marshal(receiver)
	err = stub.PutState(op.To, opsAsBytes)	
	if err != nil {
		return nil, err
	}
	
	//get the AllOperations index
	allOpAsBytes, err := stub.GetState("allOp")
	if err != nil {
		return nil, errors.New("updateEmbedded: Failed to get all Operations")
	}

	//Update operations array and commit to BC
	fmt.Println("updateEmbedded commit operation to ledger");
	var ops AllOperations
	json.Unmarshal(allOpAsBytes, &ops)
	ops.Operations = append(ops.Operations,op)
	opsAsBytes, _ = json.Marshal(ops)
	err = stub.PutState("allOp", opsAsBytes)	
	if err != nil {
		return nil, err
	}
	
	
	return nil, nil

}
=======
/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"encoding/json"
	"time"
	"github.com/hyperledger/fabric/core/chaincode/shim"

)

// Test comment
// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// Maximum number of transactions to return
const NUM_TX_TO_RETURN = 27
const NUM_OP_TO_RETURN = 27

// Smart Contract Id, Compatibility Reference numbers
const TRAVEL_CONTRACT   = "Paris"
const FEEDBACK_CONTRACT = "Feedback"
const REFERENCE_TABLES = "Compatibility"

// Blockchain point transaction record
type Transaction struct {
	RefNumber   string   `json:"RefNumber"`
	Date 		time.Time   `json:"Date"`
	Description string   `json:"description"`
	Type 		string   `json:"Type"`
	Amount    	float64  `json:"Amount"`
	Money    	float64  `json:"Money"`
	Activities  int      `json:"FeedbackActivitiesDone"`
	To			string   `json:"ToUserid"`
	From		string   `json:"FromUserid"`
	ToName	    string   `json:"ToName"`
	FromName	string   `json:"FromName"`
	ContractId	string   `json:"ContractId"`
	StatusCode	int 	 `json:"StatusCode"`
	StatusMsg	string   `json:"StatusMsg"`
}

// Blockchain operation record
type Operation struct {
	RefNumber	string		`json:"RefNumber"`
	Date 		time.Time	`json:"Date"`
	From		string		`json:"FromUserid"`
	FromName	string		`json:"FromName"`
	To		string		`json:"ToUserid"`
	ToName		string		`json:"ToName"`
	Description	string		`json:"description"`
	Version		string		`json:"Version"`
	StatusCode	int		`json:"StatusCode"`
	StatusMsg	string		`json:"StatusMsg"`
}

// Smart contract metadata record
type Contract struct {
	Id			string   `json:"ID"`
	BusinessId  string   `json:"BusinessId"`
	BusinessName string   `json:"BusinessName"`
	Title		string   `json:"Title"`
	Description string   `json:"Description"`
	Conditions  []string `json:"Conditions"`
	Icon        string 	 `json:"Icon"`
	StartDate   time.Time   `json:"StartDate"`
	EndDate		time.Time   `json:"EndDate"`
	Method	    string   `json:"Method"`
	DiscountRate float64  `json:"DiscountRate"`
}

type Compatibility struct {
	Chassis		[]string	`json:"Chassis"`
	Powertrain	[]string	`json:"Powertrain"`
	Safety		[]string	`json:"Safety"`
	Telematics	[]string	`json:"Telematics"`
}

// Subsystem running version record
type Subsystem struct {
	SsId		string		`json:"SsId"`
	Name   		string		`json:"Name"`
	osVersion	string		`json:"osVersion"`
	NumTxs		int		`json:"NumberOfTransactions"`
	Effectivity	string		`json:"EffectivityDate"`
	Manufactured	string		`json:"MfgDate"`
	Modified	string		`json:"LastModifiedDate"`
}

// Open Points member record
type User struct {
	UserId		string   `json:"UserId"`
	Name   		string   `json:"Name"`
	Balance 	float64  `json:"Balance"`
	NumTxs 	    int      `json:"NumberOfTransactions"`
	Status      string 	 `json:"Status"`
	Expiration  string   `json:"ExpirationDate"`
	Join		string   `json:"JoinDate"`
	Modified	string   `json:"LastModifiedDate"`
}


// Array for storing all open points transactions
type AllTransactions struct{
	Transactions []Transaction `json:"transactions"`
}

// Array for storing all operations
type AllOperations struct{
	Operations []Operation `json:"operations"`
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// ============================================================================================================================
// Init - reset all the things
// ============================================================================================================================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	
	// Create the 'Bank' user and add it to the blockchain
	var bank User
	bank.UserId = "B1928564";
	bank.Name = "OpenFN"
	bank.Balance = 1000000
	bank.Status  = "Originator"
	bank.Expiration = "2099-12-31"
	bank.Join  = "2015-01-01"
	bank.Modified = "2016-05-06"
	bank.NumTxs  = 0
	
	
	jsonAsBytes, _ := json.Marshal(bank)
	err = stub.PutState(bank.UserId, jsonAsBytes)								
	if err != nil {
		fmt.Println("Error Creating Bank user account")
		return nil, err
	}
	
	
    // Create the 'Travel Agency' user and add it to the blockchain
	var travel User
	travel.UserId = "T5940872";
	travel.Name = "Open Travel"
	travel.Balance = 500000
	travel.Status  = "Member"
	travel.Expiration = "2099-12-31"
	travel.Join  = "2015-01-01"
	travel.Modified = "2016-05-06"
	travel.NumTxs  = 0
	
	jsonAsBytes, _ = json.Marshal(travel)
	err = stub.PutState(travel.UserId, jsonAsBytes)								
	if err != nil {
		fmt.Println("Error Creating Travel user account")
		return nil, err
	}
	
	
	// Create the 'Natalie' user and add her to the blockchain
	var natalie User
	natalie.UserId = "U2974034";
	natalie.Name = "Natalie"
	natalie.Balance = 1000
	natalie.Status  = "Platinum"
	natalie.Expiration = "2017-06-01"
	natalie.Join  = "2015-05-31"
	natalie.Modified = "2016-05-06"
	natalie.NumTxs  = 0
	
	jsonAsBytes, _ = json.Marshal(natalie)
	err = stub.PutState(natalie.UserId, jsonAsBytes)								
	if err != nil {
		fmt.Println("Error Creating Natalie user account")
		return nil, err
	}
	
	
	// Create the 'Anthony' user and add him to the blockchain
	var anthony User
	anthony.UserId = "U3151672";
	anthony.Name = "Anthony"
	anthony.Balance = 50000
	anthony.Status  = "Silver"
	anthony.Expiration = "2017-03-15"
	anthony.Join  = "2015-08-15"
	anthony.Modified = "2016-04-17"
	anthony.NumTxs  = 0
	
	jsonAsBytes, _ = json.Marshal(anthony)
	err = stub.PutState(anthony.UserId, jsonAsBytes)								
	if err != nil {
		fmt.Println("Error Creating Anthony user account")
		return nil, err
	}
	
	
	// Create an array for storing all transactions, and store the array on the blockchain
	var transactions AllTransactions
	jsonAsBytes, _ = json.Marshal(transactions)
	err = stub.PutState("allTx", jsonAsBytes)
	if err != nil {
		return nil, err
	}
	
	// Create transaction reference number and store it on the blockchain
	var refNumber int
	
	refNumber = 2985674978
	jsonAsBytes, _ = json.Marshal(refNumber)
	err = stub.PutState("refNumber", jsonAsBytes)								
	if err != nil {
		fmt.Println("Error Creating reference number")
		return nil, err
	}

	
	// Create contract metadata for double points and add it to the blockchain
	var double Contract
	double.Id = TRAVEL_CONTRACT
	double.BusinessId  = "T5940872"
	double.BusinessName = "Open Travel"
	double.Title = "Paris for Less"
	double.Description = "All Paris travel activities are half the stated point price"
	double.Conditions = append(double.Conditions, "Half off dining and travel activities in Paris")
	double.Conditions = append(double.Conditions, "Valid from May 11, 2016") 
	double.Icon = ""
	double.Method = "travelContract"
	
	startDate, _  := time.Parse(time.RFC822, "11 May 16 12:00 UTC")
	double.StartDate = startDate
	endDate, _  := time.Parse(time.RFC822, "31 Dec 60 11:59 UTC")
	double.EndDate = endDate
	
	jsonAsBytes, _ = json.Marshal(double)
	err = stub.PutState(TRAVEL_CONTRACT, jsonAsBytes)								
	if err != nil {
		fmt.Println("Error creating double contract")
		return nil, err
	}
	
	
	// Create contract metadata for feedback points and add it to the blockchain
    var feedback Contract
	feedback.Id = FEEDBACK_CONTRACT
	feedback.BusinessId  = "T5940872"
	feedback.BusinessName = "Open Travel"
	feedback.Title = "Points for Feedback"
	feedback.Description = "Earn points by sharing your thoughts on travel package and activities"
	feedback.Conditions = append(feedback.Conditions, "1,000 points for travel package ")
	feedback.Conditions = append(feedback.Conditions, "Valid from May 24, 2016")
	feedback.Icon = ""
	feedback.Method = "feedbackContract"
	startDate, _  = time.Parse(time.RFC822, "24 May 16 12:00 UTC")
	feedback.StartDate = startDate
	endDate, _  = time.Parse(time.RFC822, "31 Dec 60 11:59 UTC")
	feedback.EndDate = endDate
	
	jsonAsBytes, _ = json.Marshal(feedback)
	err = stub.PutState(FEEDBACK_CONTRACT, jsonAsBytes)								
	if err != nil {
		fmt.Println("Error creating feedback contract")
		return nil, err
	}

	// Create the 'Chassis' subsystem and add it to the blockchain
	var chassis Subsystem
	chassis.SsId = "Chassis"
	chassis.Name = "Chassis Subsystem"
	chassis.osVersion = "001"
	chassis.NumTxs = 0

	jsonAsBytes, _ = json.Marshal(chassis)
	err = stub.PutState(chassis.SsId, jsonAsBytes)								
	if err != nil {
		fmt.Println("Error Creating Chassis user account")
		return nil, err
	}

	// Create the 'Powertrain' subsystem and add it to the blockchain
	var powertrain Subsystem
	powertrain.SsId = "Powertrain"
	powertrain.Name = "Powertrain Subsystem"
	powertrain.osVersion = "001"
	powertrain.NumTxs  = 0
	
	jsonAsBytes, _ = json.Marshal(powertrain)
	err = stub.PutState(powertrain.SsId, jsonAsBytes)								
	if err != nil {
		fmt.Println("Error Creating Powertrain user account")
		return nil, err
	}

	// Create the 'Safety' subsystem and add it to the blockchain
	var safety Subsystem
	safety.SsId = "Safety"
	safety.Name = "Safety Subsystem"
	safety.osVersion = "001"
	safety.NumTxs  = 0

	jsonAsBytes, _ = json.Marshal(safety)
	err = stub.PutState(safety.SsId, jsonAsBytes)								
	if err != nil {
		fmt.Println("Error Creating Safety user account")
		return nil, err
	}

	// Create the 'Telematics' subsystem and add it to the blockchain
	var telematics Subsystem
	telematics.SsId = "Telematics"
	telematics.Name = "Telematics Subsystem"
	telematics.osVersion = "001"
	telematics.NumTxs  = 0

	jsonAsBytes, _ = json.Marshal(telematics)
	err = stub.PutState(telematics.SsId, jsonAsBytes)								
	if err != nil {
		fmt.Println("Error Creating Telematics user account")
		return nil, err
	}

	// Create compatibility reference data and add it to the blockchain
	var reference Compatibility
		reference.Chassis = append(reference.Chassis,"023")
		reference.Chassis = append(reference.Chassis,"024")
		reference.Chassis = append(reference.Chassis,"025")
		reference.Chassis = append(reference.Chassis,"031")
		reference.Powertrain = append(reference.Powertrain,"120")
		reference.Powertrain = append(reference.Powertrain,"131")
		reference.Powertrain = append(reference.Powertrain,"175")
		reference.Powertrain = append(reference.Powertrain,"192")
		reference.Safety = append(reference.Safety,"137")
		reference.Safety = append(reference.Safety,"139")
		reference.Safety = append(reference.Safety,"140")
		reference.Telematics = append(reference.Telematics,"036")
		reference.Telematics = append(reference.Telematics,"047")
		reference.Telematics = append(reference.Telematics,"091")
	
	jsonAsBytes, _ = json.Marshal(reference)
	err = stub.PutState(REFERENCE_TABLES, jsonAsBytes)								
	if err != nil {
		fmt.Println("Error creating reference tables")
		return nil, err
	}

	// Create an array of contract ids to keep track of all contracts
	var contractIds []string
	contractIds = append(contractIds, TRAVEL_CONTRACT);
	contractIds = append(contractIds, FEEDBACK_CONTRACT);
	//contractIds = append(contractIds, REFERENCE_TABLES);
	
	jsonAsBytes, _ = json.Marshal(contractIds)
	err = stub.PutState("contractIds", jsonAsBytes)								
	if err != nil {
		fmt.Println("Error storing contract Ids on blockchain")
		return nil, err
	}
	
	return nil, nil
}

// ============================================================================================================================
// Run - Our entry point for Invocations - [LEGACY] obc-peer 4/25/2016
// ============================================================================================================================
func (t *SimpleChaincode) Run(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("run is running " + function)
	return t.Invoke(stub, function, args)
}

// ============================================================================================================================
// Invoke - Our entry point for Invocations
// ============================================================================================================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)


	
	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	} else if function == "transferPoints" {											//create a transaction
		return t.transferPoints(stub, args)
	} else if function == "updateEmbedded" {											//create a transaction
		return t.updateEmbedded(stub, args)
	} else if function == "addSmartContract" {											//create a transaction
		return t.addSmartContract(stub, args)
	} else if function == "incrementReferenceNumber" {											//create a transaction
		return t.incrementReferenceNumber(stub, args)
	} 
		
	fmt.Println("invoke did not find func: " + function)					//error

	return nil, errors.New("Received unknown function invocation")
}

// ============================================================================================================================
// Query - Our entry point for Queries
// ============================================================================================================================
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	
	if function == "getTxs" { return t.getTxs(stub, args[1]) }
	if function == "getOps" { return t.getOps(stub, args[1]) }
	if function == "getUserAccount" { return t.getUserAccount(stub, args[1]) }
	if function == "getSubsystem" { return t.getSubsystem(stub, args[1]) }
	if function == "getAllContracts" { return t.getAllContracts(stub) }
	if function == "getReferenceNumber" { return t.getReferenceNumber(stub) }
	if function == "getReferenceTables" { return t.getReferenceTables(stub) }
	
	fmt.Println("query did not find func: " + function)						//error

	return nil, errors.New("Received unknown function query")
}

// ============================================================================================================================
// Get Open Points member account from the blockchain
// ============================================================================================================================
func (t *SimpleChaincode) getUserAccount(stub shim.ChaincodeStubInterface, userId string)([]byte, error){
	
	fmt.Println("Start getUserAccount")
	fmt.Println("Looking for user with ID " + userId);

	//get the User index
	fdAsBytes, err := stub.GetState(userId)
	if err != nil {
		return nil, errors.New("Failed to get user account from blockchain")
	}

	return fdAsBytes, nil
	
}

// ============================================================================================================================
// Get Subsystem 'account' from the blockchain
// ============================================================================================================================
func (t *SimpleChaincode) getSubsystem(stub shim.ChaincodeStubInterface, SsId string)([]byte, error){
	
	fmt.Println("Start getSubsystem()")
	fmt.Println("Looking for subsystem with ID " + SsId);

	//get the User index
	fdAsBytes, err := stub.GetState(SsId)
	if err != nil {
		return nil, errors.New("Failed to get subsystem from blockchain")
	}

	return fdAsBytes, nil
	
}

// ============================================================================================================================
// Get all transactions that involve a particular user
// ============================================================================================================================
func (t *SimpleChaincode) getTxs(stub shim.ChaincodeStubInterface, userId string)([]byte, error){
	
	var res AllTransactions

	fmt.Println("Start find getTransactions")
	fmt.Println("Looking for " + userId);

	//get the AllTransactions index
	allTxAsBytes, err := stub.GetState("allTx")
	if err != nil {
		return nil, errors.New("Failed to get all Transactions")
	}

	var txs AllTransactions
	json.Unmarshal(allTxAsBytes, &txs)
	numTxs := len(txs.Transactions)

	for i := numTxs -1; i >= 0; i-- {
	    if txs.Transactions[i].From == userId{
			res.Transactions = append(res.Transactions, txs.Transactions[i])
		}

		if txs.Transactions[i].To == userId{
			res.Transactions = append(res.Transactions, txs.Transactions[i])
		}
		
		if (len(res.Transactions) >= NUM_TX_TO_RETURN) { break }
	}

	resAsBytes, _ := json.Marshal(res)

	return resAsBytes, nil
	
}

// ============================================================================================================================
// Get all operations that involve a particular subsystem
// ============================================================================================================================
func (t *SimpleChaincode) getOps(stub shim.ChaincodeStubInterface, SsId string)([]byte, error){
	
	var res AllOperations

	fmt.Println("Start getOps")
	fmt.Println("Looking for " + SsId);

	//get the AllOperations index
	allOpAsBytes, err := stub.GetState("allOp")
	if err != nil {
		return nil, errors.New("Failed to get all Operations")
	}

	var ops AllOperations
	json.Unmarshal(allOpAsBytes, &ops)
	numOps := len(ops.Operations)

	for i := numOps -1; i >= 0; i-- {
	    if ops.Operations[i].From == SsId{
			res.Operations = append(res.Operations, ops.Operations[i])
		}

		if ops.Operations[i].To == SsId{
			res.Operations = append(res.Operations, ops.Operations[i])
		}
		
		if (len(res.Operations) >= NUM_OP_TO_RETURN) { break }
	}

	resAsBytes, _ := json.Marshal(res)

	return resAsBytes, nil
	
}

// ============================================================================================================================
// Get the smart contract metadata from the blockchain
// ============================================================================================================================
func (t *SimpleChaincode) getAllContracts(stub shim.ChaincodeStubInterface)([]byte, error)  {

	contractIdsAsBytes, _ := stub.GetState("contractIds")
	var contractIds []string
	json.Unmarshal(contractIdsAsBytes, &contractIds)
	
	var allContracts []Contract
	for i := range contractIds{
		contractAsBytes, _ := stub.GetState(contractIds[i])
		var thisContract Contract
		json.Unmarshal(contractAsBytes, &thisContract)
		allContracts = append(allContracts, thisContract)
	}

	asBytes, _ := json.Marshal(allContracts)
	return asBytes, nil

}

func (t *SimpleChaincode) getReferenceTables(stub shim.ChaincodeStubInterface)([]byte, error)  {

	refTablesAsBytes, _ := stub.GetState(REFERENCE_TABLES)
	return refTablesAsBytes, nil

}

func (t *SimpleChaincode) incrementReferenceNumber(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {


	var refNumber int
	refNumberBytes, numErr := stub.GetState("refNumber")
	if numErr != nil {
		fmt.Println("Error Getting  ref number")
		return nil, numErr
	}
	
	json.Unmarshal(refNumberBytes, &refNumber)
	refNumber = refNumber + 1;
	refNumberBytes, _ = json.Marshal(refNumber)
	err := stub.PutState("refNumber", refNumberBytes)								
	if err != nil {
		fmt.Println("Error Creating updating ref number")
		return nil, err
	}

	return nil, nil
}

func (t *SimpleChaincode) getReferenceNumber(stub shim.ChaincodeStubInterface)([]byte, error)  {

	refNumberBytes, numErr := stub.GetState("refNumber")
	if numErr != nil {
		fmt.Println("Error Getting  ref number")
		return nil, numErr
	}
	
	return refNumberBytes, nil

}

// ============================================================================================================================
// Smart contract for giving user double points
// ============================================================================================================================
func travelContract(tx Transaction, stub shim.ChaincodeStubInterface) float64 {


	contractAsBytes, err := stub.GetState(TRAVEL_CONTRACT)
	if err != nil {
		return -99
	}
	var contract Contract
	json.Unmarshal(contractAsBytes, &contract)
	
	var pointsToTransfer float64
	pointsToTransfer = tx.Amount
	if (tx.Date.After(contract.StartDate) && tx.Date.Before(contract.EndDate)) {
	     pointsToTransfer = pointsToTransfer * 0.5
	}
 
 
  return pointsToTransfer
  
  
}


// ============================================================================================================================
// Smart contract for giving user points for completing feedback surveys
// ============================================================================================================================
func feedbackContract(tx Transaction, stub shim.ChaincodeStubInterface) float64 {
  

	contractAsBytes, err := stub.GetState(FEEDBACK_CONTRACT)
	if err != nil {
		return -99
	}
	var contract Contract
	json.Unmarshal(contractAsBytes, &contract)
	
	var pointsToTransfer float64
	pointsToTransfer = 0
	if (tx.Date.After(contract.StartDate) && tx.Date.Before(contract.EndDate)) {
	     pointsToTransfer = 1000
		 
		 if (tx.Activities > 0) {
			pointsToTransfer = pointsToTransfer + float64(tx.Activities)*100
		 }
	}
  
  
  return pointsToTransfer
  
  
}

func (t *SimpleChaincode) addSmartContract(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {


	// Create new smart contract based on user input
	var smartContract Contract
	
	discountRate, err := strconv.ParseFloat(args[4], 64)
	if err != nil {
		smartContract.Title= "Invalid Contract"
	}else{
		smartContract.DiscountRate = discountRate
	}
	
	
	smartContract.Id = args[0]
	smartContract.BusinessId  = "T5940872"
	smartContract.BusinessName = "Open Travel"
	smartContract.Title = args[1]
	smartContract.Description = ""
	smartContract.Conditions = append(smartContract.Conditions, args[2])
	smartContract.Conditions = append(smartContract.Conditions, args[3]) 
	smartContract.Icon = ""
	smartContract.Method = "travelContract"
	
	
	jsonAsBytes, _ := json.Marshal(smartContract)
	err = stub.PutState(smartContract.Id, jsonAsBytes)								
	if err != nil {
		fmt.Println("Error adding new smart contract")
		return nil, err
	}

	contractIdsAsBytes, _ := stub.GetState("contractIds")
	var contractIds []string
	json.Unmarshal(contractIdsAsBytes, &contractIds)
	
	
	var contractIdFound bool
	contractIdFound = false;
	for i := range contractIds{
		if (contractIds[i] == smartContract.Id)  {
			contractIdFound = true;
		}
	}
	
	if (!contractIdFound) {
		contractIds = append(contractIds, smartContract.Id);
	}
	
	
	jsonAsBytes, _ = json.Marshal(contractIds)
	err = stub.PutState("contractIds", jsonAsBytes)								
	if err != nil {
		fmt.Println("Error storing contract Ids on blockchain")
		return nil, err
	}

	return nil, nil

}


// ============================================================================================================================
// Transfer points between members of the Open Points Network
// ============================================================================================================================
func (t *SimpleChaincode) transferPoints(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	fmt.Println("Running transferPoints")
	currentDateStr := time.Now().Format(time.RFC822)
	startDate, _  := time.Parse(time.RFC822, currentDateStr)

	
	var tx Transaction
	tx.Date 		= startDate
	tx.To 			= args[0]
	tx.From 		= args[1]
	tx.Type 	    = args[2]
	tx.Description 	= args[3]
	tx.ContractId 	= args[4]
	activities, _  := strconv.Atoi(args[5])
	tx.Activities   = activities
	tx.StatusCode 	= 1
	tx.StatusMsg 	= "Transaction Completed"
	
	
	amountValue, err := strconv.ParseFloat(args[6], 64)
	if err != nil {
		tx.StatusCode = 0
		tx.StatusMsg = "Invalid Amount"
	}else{
		tx.Amount = amountValue
	}
	
	moneyValue, err := strconv.ParseFloat(args[7], 64)
	if err != nil {
		tx.StatusCode = 0
		tx.StatusMsg = "Invalid Amount"
	}else{
		tx.Money = moneyValue
	}
	
	
	// Get the current reference number and update it
	var refNumber int
	refNumberBytes, numErr := stub.GetState("refNumber")
	if numErr != nil {
		fmt.Println("Error Getting  ref number for transferring points")
		return nil, err
	}
	
	json.Unmarshal(refNumberBytes, &refNumber)
	tx.RefNumber 	= strconv.Itoa(refNumber)
	refNumber = refNumber + 1;
	refNumberBytes, _ = json.Marshal(refNumber)
	err = stub.PutState("refNumber", refNumberBytes)								
	if err != nil {
		fmt.Println("Error Creating updating ref number")
		return nil, err
	}
	
	// Determine point amount to transfer based on contract type
	if (tx.ContractId == TRAVEL_CONTRACT) {
		tx.Amount = travelContract(tx, stub)
	} else if (tx.ContractId == FEEDBACK_CONTRACT) {
		tx.Amount = feedbackContract(tx, stub)
	}  else {
	
		contractIdsAsBytes, _ := stub.GetState("contractIds")
		var contractIds []string
		json.Unmarshal(contractIdsAsBytes, &contractIds)
	
		for i := range contractIds{
			contractAsBytes, _ := stub.GetState(contractIds[i])
			var thisContract Contract
			json.Unmarshal(contractAsBytes, &thisContract)
			
			if (tx.ContractId == thisContract.Id) {
				tx.Amount = tx.Amount -  (tx.Amount * thisContract.DiscountRate);
			}
		}
	
	
	
	}
	
	
	// Get Receiver account from BC and update point balance
	rfidBytes, err := stub.GetState(tx.To)
	if err != nil {
		return nil, errors.New("transferPoints Failed to get Receiver from BC")
	}
	var receiver User
	fmt.Println("transferPoints Unmarshalling User Struct");
	err = json.Unmarshal(rfidBytes, &receiver)
	receiver.Balance = receiver.Balance  + tx.Amount
	receiver.Modified = currentDateStr
	receiver.NumTxs = receiver.NumTxs + 1
	tx.ToName = receiver.Name;
	
	
	//Commit Receiver to ledger
	fmt.Println("transferPoints Commit Updated receiver To Ledger");
	txsAsBytes, _ := json.Marshal(receiver)
	err = stub.PutState(tx.To, txsAsBytes)	
	if err != nil {
		return nil, err
	}
	
	// Get Sender account from BC nd update point balance
	rfidBytes, err = stub.GetState(tx.From)
	if err != nil {
		return nil, errors.New("transferPoints Failed to get Financial Institution")
	}
	var sender User
	fmt.Println("transferPoints Unmarshalling Sender");
	err = json.Unmarshal(rfidBytes, &sender)
	sender.Balance   = sender.Balance  - tx.Amount
	sender.Modified = currentDateStr
	sender.NumTxs = sender.NumTxs + 1
	tx.FromName = sender.Name;
	
	//Commit Sender to ledger
	fmt.Println("transferPoints Commit Updated Sender To Ledger");
	txsAsBytes, _ = json.Marshal(sender)
	err = stub.PutState(tx.From, txsAsBytes)	
	if err != nil {
		return nil, err
	}
	
	
	//get the AllTransactions index
	allTxAsBytes, err := stub.GetState("allTx")
	if err != nil {
		return nil, errors.New("transferPoints: Failed to get all Transactions")
	}

	//Update transactions arrary and commit to BC
	fmt.Println("SubmitTx Commit Transaction To Ledger");
	var txs AllTransactions
	json.Unmarshal(allTxAsBytes, &txs)
	txs.Transactions = append(txs.Transactions, tx)
	txsAsBytes, _ = json.Marshal(txs)
	err = stub.PutState("allTx", txsAsBytes)	
	if err != nil {
		return nil, err
	}
	
	
	return nil, nil

}

// ============================================================================================================================
// Update subsystem with new embedded operating software
// ============================================================================================================================
func (t *SimpleChaincode) updateEmbedded(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var err error

	fmt.Println("Running updateEmbedded()")
	currentDateStr := time.Now().Format(time.RFC822)
	startDate, _  := time.Parse(time.RFC822, currentDateStr)

	
	var op Operation
	op.Date		= startDate
	op.To 		= args[0]
	op.From 	= args[1]
	op.Description 	= args[2]
	op.Version	= args[3]
	op.StatusCode 	= 1
	op.StatusMsg 	= "Operation accepted"
	
	// Get the current reference number and update it
	var refNumber int
	refNumberBytes, numErr := stub.GetState("refNumber")
	if numErr != nil {
		fmt.Println("Error Getting  ref number for operation")
		return nil, err
	}
	
	json.Unmarshal(refNumberBytes, &refNumber)
	op.RefNumber 	= strconv.Itoa(refNumber)
	refNumber = refNumber + 1;
	refNumberBytes, _ = json.Marshal(refNumber)
	err = stub.PutState("refNumber", refNumberBytes)								
	if err != nil {
		fmt.Println("Error updating ref number")
		return nil, err
	}
	
	fmt.Printf("We're here...")
	refTablesAsBytes, _ := stub.GetState(REFERENCE_TABLES)
	refTables := fmt.Sprintf("%s",refTablesAsBytes)
	if (strings.Contains(refTables,op.Version)) {
		fmt.Printf("Hit")
	}

	// Get Receiver account from BC and update point balance
	rfidBytes, err := stub.GetState(op.To)
	if err != nil {
		return nil, errors.New("updateEmbedded failed to get receiver from BC")
	}

	var receiver Subsystem
	fmt.Println("updateEmbedded unmarshalling Subsystem struct");
	err = json.Unmarshal(rfidBytes, &receiver)
	receiver.osVersion = op.Version
	receiver.Modified = currentDateStr
	receiver.NumTxs = receiver.NumTxs + 1
	op.ToName = receiver.Name;
	
	
	//Commit Receiver to ledger
	fmt.Println("updateEmbedded commit updated receiver To ledger");
	opsAsBytes, _ := json.Marshal(receiver)
	err = stub.PutState(op.To, opsAsBytes)	
	if err != nil {
		return nil, err
	}
	
	//get the AllOperations index
	allOpAsBytes, err := stub.GetState("allOp")
	if err != nil {
		return nil, errors.New("updateEmbedded: Failed to get all Operations")
	}

	//Update operations array and commit to BC
	fmt.Println("updateEmbedded commit operation to ledger");
	var ops AllOperations
	json.Unmarshal(allOpAsBytes, &ops)
	ops.Operations = append(ops.Operations,op)
	opsAsBytes, _ = json.Marshal(ops)
	err = stub.PutState("allOp", opsAsBytes)	
	if err != nil {
		return nil, err
	}
	
	
	return nil, nil

}
>>>>>>> 05e3fef3ea2dfbbc37d7073c650a89cbbe7bd648
