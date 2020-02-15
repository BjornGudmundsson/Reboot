package users

import (
	"bufio"
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
}

func (u User) String() string {
	return u.Number + " " + "PW"
}

func WriteUserToDB(u User) error {
	exists := CheckIfUserExists(u.Number)
	if exists == nil {
		return errors.New("Already exists")
	}
	f, e := os.OpenFile("db.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if e != nil {
		return e
	}
	defer f.Close()
	_, eW := f.WriteString(u.String() + "\n")
	if eW != nil {
		return eW
	}
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

func LoginUser(u User) error {
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

		if nbr == u.Number {
			return nil
		}
	}
	return errors.New("This user could not be logged in")
}
