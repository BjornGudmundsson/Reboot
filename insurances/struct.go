package insurances

import (
	"errors"
	"strconv"
)

type Insurance struct {
	Name      string
	Payment   int
	Signature []byte
	Desc      string
	ID        uint32
}

type InsuranceBuy struct {
	Name string
	Min  int
	Max  int
}

func (i Insurance) String() string {
	ret := ""
	ret += "Name:" + i.Name + "-"
	ret += "Payment:" + strconv.Itoa(int(i.Payment)) + "-"
	ret += "Desc:" + i.Desc
	return ret
}

func GetInsuranceFromId(id uint32) Insurance {
	return Insurance{
		Desc: "This is a demo",
	}
}

func FindInsuranceThatMatches(buy InsuranceBuy) (Insurance, error) {
	for _, b := range availableInsurance {
		if b.Name == buy.Name && (buy.Min <= b.Payment && buy.Max >= b.Payment) {
			return b, nil
		}
	}
	return Insurance{}, errors.New("could not find a matching insurance")
}
