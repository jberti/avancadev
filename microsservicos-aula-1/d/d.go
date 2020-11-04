package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Coupon struct {
	Code string
}

type TakenCoupons struct {
	Coupon []Coupon
}

func (t TakenCoupons) IsTaken(aCode string) bool{

	for _, item := range t.Coupon {
		if aCode == item.Code {
			return true
		}
	}
	return false
}

type Result struct {
	Status bool
}

var takenCoupons TakenCoupons

func main(){
	coupon := Coupon{
		Code: "abc",
	}
	takenCoupons.Coupon = append(takenCoupons.Coupon, coupon)

	http.HandleFunc("/", home)
	http.ListenAndServe(":9093", nil)
}

func home(w http.ResponseWriter, r *http.Request){

	coupon := r.PostFormValue("coupon")
	var x = ""

	if takenCoupons.IsTaken(coupon){
		x = "used"
	} else{
		x = "not-used"
	}

	err, _ := json.Marshal(x)
	if err != nil {
		log.Fatal("Error converting json")
	}

	fmt.Fprintf(w, "Cheguei no serviço D e não sei o que fazer :D ")
}
