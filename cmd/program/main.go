package main

import (
	"fmt"

	"getitqec.com/server/user/pkg/commons"
)

func main() {
	// fmt.Println(model.GenerateOTP())
	// fmt.Println(model.GenerateOTP())
	// fmt.Println(model.GenerateOTP())
	// fmt.Println(model.GenerateOTP())
	// fmt.Println(model.GenerateOTP())
	// fmt.Println(model.GenerateOTP())
	// t := time.Time{}
	m := int64(857145600000)
	fmt.Println(m)
	t2 := commons.MilliToTime(m)
	fmt.Println(t2)
	// fmt.Println((-25) % 4)
}
