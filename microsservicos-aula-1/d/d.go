package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Check(listaCupons []string, cupon string) string {

	for _, v := range listaCupons {
		if cupon == v {
			return "valid"
		}

	}
	return "invalid"
}

type Result struct {
	Status string
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9093", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	clientes := map[string][]string{
		"42730365845": {"abc", "cba"},
		"42730365945": {"abc", "cba"},
	}
	coupon := r.PostFormValue("coupon")
	cpf := r.PostFormValue("cpf")
	lista, ok := clientes[cpf]
	valid := Check(lista, coupon)

	result := Result{Status: "approved"}

	if valid == "invalid" {
		result.Status = "invalid"
	}

	if ok == false {
		result.Status = "invalid01"
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error processing json")
	}

	fmt.Fprintf(w, string(jsonData))
}
