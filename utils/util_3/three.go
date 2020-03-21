package util_3

import (
	"demo/utils/util_1"
	"fmt"
	"time"
)

// đặt biến toàn cục cho mỗi func -> có thể mock A, B khi test

var UtilsLogic = Logic

var UtilsA = util_1.A
var UtilsB = util_1.B

func Logic(cond string) string {
	res := "default"

	switch cond {
	case "A":
		res = UtilsA()
	case "B":
		res = UtilsB()
	case "C":
		// long sleep here
		time.Sleep(time.Second * 2)
		res = "C after 2s"
	}
	fmt.Println(res)
	return res
}
