package users

import "os"
import "fmt"

func init() {
	if !Exists("db.txt") {
		fmt.Println("bjorn")
		_, e := os.Create("db.txt")
		if e != nil {
			panic(e)
		}	
	}
}