package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BjornGudmundsson/Reboot/insurances"
	"github.com/BjornGudmundsson/Reboot/users"
)

func SignupForm(w http.ResponseWriter, req *http.Request) {
	j := json.NewDecoder(req.Body)
	var js users.User
	e := j.Decode(&js)
	if e != nil {
		fmt.Println(req.Body)
		fmt.Println(e.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ph := js.Number
	pw := js.PW
	h := sha256.Sum256([]byte(pw))
	hex := hex.EncodeToString(h[:])
	u := users.User{
		PW:     hex,
		Number: ph,
	}
	e = users.WriteUserToDB(u)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
	}
}

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
	pw := js.PW
	h := sha256.Sum256([]byte(pw))
	hex := hex.EncodeToString(h[:])
	u := users.User{
		PW:     hex,
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
		w.WriteHeader(http.StatusOK)
	}
}

func main() {
	http.HandleFunc("/loginForm", LoginForm)
	http.HandleFunc("/addInsurance", insurances.AcceptInsurance)
	http.HandleFunc("/myInsurances", insurances.GetMyInsurances)
	http.ListenAndServe(":8084", nil)
}
