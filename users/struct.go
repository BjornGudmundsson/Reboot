package users

import (
	"bufio"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"os"
	"strings"
)

func Exists(fn string) bool {
	if _, err := os.Stat(fn); err == nil {
		return true

	} else if os.IsNotExist(err) {
		return false

	} else {
		if err != nil {
			return false
		}
	}
	return false
}

type User struct {
	Number string
	Key    string
	K      *ecdsa.PrivateKey
}

func (u User) String() string {
	return u.Number + " " + u.Key
}

func WriteUserToDB(n string) error {
	k, eK := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if eK != nil {
		return eK
	}
	t, e1 := k.X.MarshalText()
	if e1 != nil {
		return e1
	}
	h := hex.EncodeToString(t)
	u := User{
		Number: n,
		Key:    h,
		K:      k,
	}
	DB[n] = u
	//exists := CheckIfUserExists(u.Number)
	//if exists == nil {
	//	return errors.New("Already exists")
	//}
	//f, e := os.OpenFile("db.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	//if e != nil {
	//	return e
	//}
	//defer f.Close()
	//_, eW := f.WriteString(u.String() + "\n")
	//if eW != nil {
	//	return eW
	//}
	return nil
}

func CheckIfUserExists(number string) error {
	f, e := os.OpenFile("db.txt", os.O_RDONLY, os.ModeAppend)
	if e != nil {
		return e
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l := scanner.Text()
		spl := strings.Split(l, " ")
		nbr := spl[0]
		if nbr == number {
			return nil
		}
	}
	return errors.New("This user does not exist")
}

func LoginUser(u string) error {
	f, e := os.OpenFile("db.txt", os.O_RDONLY, os.ModeAppend)
	if e != nil {
		return e
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l := scanner.Text()
		spl := strings.Split(l, " ")
		nbr := spl[0]

		if nbr == u {
			return nil
		}
	}
	return errors.New("This user could not be logged in")
}
