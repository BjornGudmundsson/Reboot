package insurances

import "strconv"

type Insurance struct {
	Name      string
	Payment   uint32
	Signature []byte
	Desc      string
	ID        uint32
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
