package main

import (
	"encoding/json"
	"net/http"

	"github.com/BjornGudmundsson/Reboot/insurances"
	"github.com/BjornGudmundsson/Reboot/users"
)

func GetCookie(phone string) *http.Cookie {
	return &http.Cookie{
		Name:  "Reboot",
		Value: phone,
	}
}

func LoginForm(w http.ResponseWriter, req *http.Request) {
	j := json.NewDecoder(req.Body)
	var js users.User
	e := j.Decode(&js)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ph := js.Number

	u := users.User{
		Number: ph,
	}
	e = users.CheckIfUserExists(ph)
	if e != nil {
		users.WriteUserToDB(u)
	}
	e = users.LoginUser(u)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		http.SetCookie(w, GetCookie(ph))
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
	}
}

func main() {
	http.HandleFunc("/loginForm", LoginForm)
	http.HandleFunc("/addInsurance", insurances.AcceptInsurance)
	http.HandleFunc("/myInsurances", insurances.GetMyInsurances)
	http.HandleFunc("/search", insurances.SearchForInsurance)
	http.ListenAndServe(":8084", nil)
}
