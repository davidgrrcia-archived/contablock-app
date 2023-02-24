package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
)

const (
	zaraAccountantYamlPath = "./config/zara-accountant.yaml"
	backendURL             = "http://localhost:4000/api/projects/"
)

type YamlFile struct {
	Accounts []Account  `yaml:"accounts"`
	Business []Business `yaml:"business"`
	Balances []Balances `yaml:"balances"`
}

type Account struct {
	Project   string  `yaml:"project" json:"project"`
	ID        string  `yaml:"id" json:"id"`
	Name      string  `yaml:"name" json:"name"`
	Level     string  `yaml:"level" json:"level"`
	BalanceID *string `yaml:"balanceId" json:"balanceId"`
}

type Business struct {
	Project            string `yaml:"project"`
	ID                 string `yaml:"id"`
	Name               string `yaml:"name"`
	SocialDenomination string `yaml:"socialDenomination"`
	RegisteredOffice   string `yaml:"registeredOffice"`
	Cif                string `yaml:"cif"`
	Logo               string `yaml:"logo"`
	EthAddress         string `yaml:"ethAddress"`
}

type Balances struct {
	Project  string   `yaml:"project"`
	ID       string   `yaml:"id"`
	Type     string   `yaml:"type"`
	Kind     *string  `yaml:"kind"`
	Parent   *string  `yaml:"parent"`
	Child    *string  `yaml:"child"`
	Accounts []string `yaml:"accounts"`
}

func main() {
	yamlFile := readYamlFile(zaraAccountantYamlPath)
	addAccounts(yamlFile)
}

func readYamlFile(path string) YamlFile {
	data, err := ioutil.ReadFile(zaraAccountantYamlPath)
	if err != nil {
		log.Fatal(err)
	}

	var yamlFile YamlFile
	err = yaml.Unmarshal(data, &yamlFile)
	if err != nil {
		log.Fatal(err)
	}

	return yamlFile
}

func addAccounts(yamlFile YamlFile) {
	for _, account := range yamlFile.Accounts {
		requestBody, err := json.Marshal(account)
		if err != nil {
			log.Fatal(err)
		}

		path := backendURL + account.Project + "/accounts"

		response, err := http.Post(path, "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			log.Fatal(err)
		}

		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(body))
	}

}
