package insurances

import (
	"errors"
	"strconv"
)

var insurance1 Insurance
var insurance2 Insurance
var insurance3 Insurance

type Insurance struct {
	Name      string
	Payment   int
	Signature []byte
	Desc      string
	ID        uint32
}

type InsuranceBuy struct {
	Name string
}

func (i Insurance) String() string {
	ret := "{"
	ret += "\"Name\":" + "\"" + i.Name + "\"" + ","
	ret += "\"Payment\":" + "\n" + strconv.Itoa(int(i.Payment)) + "\n" + ","
	ret += "\"Desc\":" + "\"" + i.Desc + "\n" + "}"
	return ret
}

func GetInsuranceFromId(id uint32) Insurance {
	for _, i := range availableInsurance {
		if i.ID == id {
			return i
		}
	}
	return Insurance{}
}

func FindInsuranceThatMatches(buy InsuranceBuy) (Insurance, error) {
	for _, b := range availableInsurance {
		if b.Name == buy.Name {
			return b, nil
		}
	}
	return Insurance{}, errors.New("could not find a matching insurance")
}
