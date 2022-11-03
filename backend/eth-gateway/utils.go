package eth_gateway

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func getContractAddressFromJSON() string {
	// Open the truffle generated json file
	jsonFile, err := os.Open("./smart-contracts/build/contracts/HoldContract.json")

	if err != nil {
		fmt.Println(err)
		return ""
	}

	// defer the closing of our json file
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}

	json.Unmarshal([]byte(byteValue), &result)

	addressJsonBytes, err := json.Marshal(result["networks"].(map[string]interface{})["5777"].(map[string]interface{})["address"])
	addressJsonString := string(addressJsonBytes)

	return addressJsonString
}

func getABIFromJSON() string {
	// Open the truffle generated json file
	jsonFile, err := os.Open("./smart-contracts/build/contracts/HoldContract.json")

	if err != nil {
		fmt.Println(err)
		return ""
	}

	// defer the closing of our json file
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}

	json.Unmarshal([]byte(byteValue), &result)

	addressJsonBytes, err := json.Marshal(result["abi"])
	addressJsonString := string(addressJsonBytes)

	return addressJsonString
}

func getBytecodeFromJSON() string {
	// Open the truffle generated json file
	jsonFile, err := os.Open("./smart-contracts/build/contracts/HoldContract.json")

	if err != nil {
		fmt.Println(err)
		return ""
	}

	// defer the closing of our json file
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}

	json.Unmarshal([]byte(byteValue), &result)

	addressJsonBytes, err := json.Marshal(result["bytecode"])
	addressJsonString := string(addressJsonBytes)

	return addressJsonString
}

func CreateAbiBytecode() {

	abi := getABIFromJSON()

	err := ioutil.WriteFile("/Users/david/Documents/GitHub/HoldContract-go/smart-contracts/build/contracts/HoldContract.abi", []byte(abi), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

	bytecode := getBytecodeFromJSON()

	bytecode = strings.Replace(bytecode, "\"", "", -1)
	bytecode = strings.Replace(bytecode, "0x", "", -1)

	err = ioutil.WriteFile("/Users/david/Documents/GitHub/HoldContract-go/smart-contracts/build/contracts/HoldContract.bin", []byte(bytecode), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}
