package insurances

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/BjornGudmundsson/Reboot/users"
)

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
	insurance1.Name = "Health insurance"
	insurance2.Name = "Car insurance"
	insurance3.Name = "Life insurance"
	insurance1.ID = 1
	insurance2.ID = 2
	insurance3.ID = 3
	availableInsurance = []Insurance{insurance1, insurance2, insurance3}
}

func AddInsurance(u users.User, i Insurance, t int) {
	k := u.K
	//fmt.Println("K: ", k)
	h := sha256.Sum256([]byte(strconv.Itoa(int(i.ID))))
	r, s, e := ecdsa.Sign(rand.Reader, k, h[:])
	fmt.Println(h)
	if e != nil {
		panic("e")
	}
	key := r.String() + ";" + s.String()
	//fmt.Println("Key: ", key)
	if _, ok := insurances[key]; !ok {
		insurances[key] = i
	}
}

func GetInsurances(u users.User) []Insurance {
	ret := make([]Insurance, 0)
	for sig, val := range insurances {
		h := sha256.Sum256([]byte(strconv.Itoa(int(val.ID))))
		spl := strings.Split(sig, ";")
		//b1, _ := hex.DecodeString(spl[0])
		//b2, _ := hex.DecodeString(spl[1])
		i1, i2 := new(big.Int), new(big.Int)
		r, _ := i1.SetString(spl[0], 10)
		s, _ := i2.SetString(spl[1], 10)
		fmt.Println("R1: ", r.String())
		fmt.Println("S1: ", s.String())
		//fmt.Println("Sig: ", sig)
		ver := ecdsa.Verify(&u.K.PublicKey, h[:], r, s)
		fmt.Println("Ver: ", ver)
		if ver {
			ret = append(ret, val)
		}
	}
	return ret
}
