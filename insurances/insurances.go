package insurances

import "github.com/BjornGudmundsson/Reboot/users"

var insurances map[string]Insurance

var availableInsurance []Insurance

func init() {
	insurances = make(map[string]Insurance, 0)
	insurance1.Desc = "This is a perfectly fine description"
	insurance2.Desc = "This is not a perfectly fine description"
	insurance3.Desc = "WTF is this desc"
	insurance1.Payment = 1500
	insurance2.Payment = 2500
	insurance3.Payment = 3500
	insurance1.Name = "Heilsutrygging"
	insurance2.Name = "Sjukratrygging"
	insurance3.Name = "Bilatrygging"
	insurance1.ID = 1
	insurance2.ID = 2
	insurance3.ID = 3
	availableInsurance = []Insurance{insurance1, insurance2, insurance3}
}

func AddInsurance(u users.User, i Insurance, r int) {
	if _, ok := insurances[u.Number]; !ok {
		insurances[u.Number] = i
	}
}

func GetInsurances(u users.User) []Insurance {
	ret := make([]Insurance, 0)
	for key, val := range insurances {
		if key == u.Number {
			ret = append(ret, val)
		}
	}
	return ret
}
