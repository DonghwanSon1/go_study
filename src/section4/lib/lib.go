package lib

import "fmt"

func init() {
	fmt.Println("lib Package > init Start!")
}

func CheckNum(c int32) bool {
	return c > 10
}
