package insurances

import "github.com/BjornGudmundsson/Reboot/users"

var insurances map[string]Insurance

func init() {
	insurances = make(map[string]Insurance, 0)
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
