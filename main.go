package main

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"fmt"
)

type person struct {
	Nome 		string `json:"nome"`
	Email 		string `json:"email"`
	Sexo 		string `json:"sexo"`
	Idade 	    int    `json:"idade"`
	Endereco 	string `json:"endereco"`
	Telefone 	string `json:"telefone"`
	Cpf 	    string `json:"cpf"`
}

type response struct {
	Data []person
}

func omdbAPI(url string) ([]byte, error) {
	var resp *http.Response
	var err error

	if resp, err = http.Get(url); err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var (
			body []byte
			err error
      		resp []person
		)

		if body, err = ioutil.ReadFile("database.json"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

   	    if err = json.Unmarshal(body,&resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if err = json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err.Error())
	}


}
