package main

import (
	"os"
	"encoding/json"
)

type Conta struct {
	Numero	int `json:"num"`
	Saldo	int `json:"sal"`
}

func main() {

	// object to json
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	// json to object
	jsonString := []byte(`{"Numero":2, "Saldo": 200}`)
	var conta2 Conta
	err = json.Unmarshal(jsonString, &conta2)
	if err != nil {
		panic(err)
	}
	println(conta2.Saldo)
}