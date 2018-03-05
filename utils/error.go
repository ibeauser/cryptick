package utils

import (
	"fmt"
)

func CheckErr(comment string, err error) {
	if err != nil {
		// panic(err)
		// log.Fatal(err)
		fmt.Println(comment, ", ", err)
	}
}
