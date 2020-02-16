package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BjornGudmundsson/Reboot/insurances"
	"github.com/BjornGudmundsson/Reboot/users"
)

func init() {
	e := users.CheckIfUserExists("8446063")
	if e != nil {
		users.WriteUserToDB("8446063")
	}
}

func GetCookie(phone string) *http.Cookie {
	return &http.Cookie{
		Name:   "Reboot",
		Value:  phone,
		Domain: "stuff.localhost.com",
	}
}

func LoginForm(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	if req.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	http.

	j := json.NewDecoder(req.Body)
	var js users.User
	e := j.Decode(&js)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ph := js.Number
	users.WriteUserToDB(ph)
	fmt.Println(users.DB)
	http.SetCookie(w, GetCookie(ph))
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/loginForm", LoginForm)
	http.HandleFunc("/addInsurance", insurances.AcceptInsurance)
	http.HandleFunc("/myInsurances", insurances.GetMyInsurances)
	http.HandleFunc("/search", insurances.SearchForInsurance)
	http.ListenAndServe(":8084", nil)
}
