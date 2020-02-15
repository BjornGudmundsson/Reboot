package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BjornGudmundsson/Reboot/users"
)

func SignupForm(w http.ResponseWriter, req *http.Request) {
	e := req.ParseForm()
	if e != nil {
		w.Write([]byte("Could not parse form"))
		return
	}
	j := json.NewDecoder(req.Body)
	var js users.User
	e = j.Decode(&js)
	if e != nil {
		fmt.Println(req.Body)
		fmt.Println(e.Error())
		w.Write([]byte("Bjo"))
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
		w.Write([]byte(e.Error()))
	} else {
		w.Write([]byte("wrote the user to the database"))
	}
}

func LoginForm(w http.ResponseWriter, req *http.Request) {
	e := req.ParseForm()
	if e != nil {
		w.Write([]byte("Could not parse form"))
		return
	}
	j := json.NewDecoder(req.Body)
	var js users.User
	e = j.Decode(&js)
	if e != nil {
		panic(e)
	}
	ph := js.Number
	pw := js.PW
	h := sha256.Sum256([]byte(pw))
	hex := hex.EncodeToString(h[:])
	u := users.User{
		PW:     hex,
		Number: ph,
	}
	e = users.LoginUser(u)
	if e != nil {
		w.Write([]byte("Could not be authenticated"))
	} else {
		w.Write([]byte("You are logged in"))
	}
}

func main() {
	http.HandleFunc("/signupForm", SignupForm)
	http.HandleFunc("/loginForm", LoginForm)
	http.ListenAndServe(":8084", nil)
}
