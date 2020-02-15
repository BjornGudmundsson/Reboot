package insurances

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/BjornGudmundsson/Reboot/users"
)

func GetMyInsurances(w http.ResponseWriter, req *http.Request) {
	c, e := req.Cookie("Reboot")
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	phone := c.Value
	u := users.GetUser(phone)
	arr := GetInsurances(u)
	s := ""
	for _, i := range arr {
		s += i.String() + ";"
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(s))
}

func AcceptInsurance(w http.ResponseWriter, req *http.Request) {
	c, e := req.Cookie("Reboot")
	if e != nil {
		fmt.Println("Bjo")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	b := make([]byte, 100)
	req.Body.Read(b)
	phone := c.Value
	id, e := strconv.Atoi(string(b[:2]))
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	i := GetInsuranceFromId(uint32(id))
	AddInsurance(users.GetUser(phone), i, 0)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	fmt.Println(insurances)
}

func SearchForInsurance(w http.ResponseWriter, req *http.Request) {
	_, e := req.Cookie("Reboot")
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var buy InsuranceBuy
	j := json.NewDecoder(req.Body)
	e = j.Decode(&buy)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	i, eI := FindInsuranceThatMatches(buy)
	if eI != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Write([]byte(i.String()))
	w.WriteHeader(http.StatusOK)
}
