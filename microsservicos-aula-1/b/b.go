package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-retryablehttp"
)

type Result struct {
	Status string
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9091", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	coupon := r.PostFormValue("coupon")
	ccNumber := r.PostFormValue("ccNumber")
	cpf := r.PostFormValue("cpf")
	result := Result{Status: "invalid"}

	if ccNumber == "1" {
		result.Status = "approved"
	}

	if len(coupon) != 0 {
		result.Status = "invalid"

		resultCoupon := makeHttpCall("http://localhost:9092", coupon, cpf)

		if resultCoupon.Status == "invalid" {
			result.Status = "invalid coupon"
		}
		if resultCoupon.Status == "invalid02" {
			result.Status = "cpf não aplicável ao cupon"
		}

		if resultCoupon.Status == "invalid03" {
			result.Status = "cpf não tem cupon"
		}

		if resultCoupon.Status == "approved" {
			result.Status = "approved"
		}
		if resultCoupon.Status == "Servidor fora do ar!" {
			result.Status = "Servidor fora do ar!"
		}
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error processing json")
	}

	fmt.Fprintf(w, string(jsonData))
}

func makeHttpCall(urlMicroservice string, coupon string, cpf string) Result {

	values := url.Values{}
	values.Add("coupon", coupon)
	values.Add("cpf", cpf)

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 5

	res, err := retryClient.PostForm(urlMicroservice, values)
	if err != nil {
		result := Result{Status: "Servidor fora do ar!"}
		return result
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error processing result")
	}

	result := Result{}

	json.Unmarshal(data, &result)

	return result

}
