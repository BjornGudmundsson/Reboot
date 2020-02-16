package insurances

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/BjornGudmundsson/Reboot/users"
)

func GetMyInsurances(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "True")
	if req.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	var phone string
	c, e := req.Cookie("Reboot")
	if e != nil {
		phone = "8446063"
	} else {
		phone = c.Value
	}
	u, _ := users.GetUser(phone)
	arr := GetInsurances(u)
	s := ""
	for _, i := range arr {
		s = i.String()
	}
	w.Write([]byte(s))
}

func AcceptInsurance(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "True")
	if req.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	var phone string
	c, e := req.Cookie("Reboot")
	if e != nil {
		phone = "8446063"
	} else {
		phone = c.Value
	}
	b := make([]byte, 100)
	req.Body.Read(b)
	id, e := strconv.Atoi(string(b[:1]))
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	i := GetInsuranceFromId(uint32(id))
	if i.ID == uint32(0) {
		w.WriteHeader(http.StatusNoContent)
	}
	u, e := users.GetUser(phone)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	AddInsurance(u, i, id)
	w.WriteHeader(http.StatusOK)
}

func SearchForInsurance(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "True")
	if req.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	/*var phone string
	c, e := req.Cookie("Reboot")
	if e != nil {
		phone = "8446063"
	} else {
		phone = c.Value
	}*/
	var buy InsuranceBuy
	j := json.NewDecoder(req.Body)
	e := j.Decode(&buy)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	i, eI := FindInsuranceThatMatches(buy)
	if eI != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Write([]byte(i.String()))
}
