package util_1

import (
	"fmt"
	"time"
)

// chỉ dùng func -> ko mock được khi test

func Logic(cond string) string {
	res := "default"

	switch cond {
	case "A":
		res = A()
	case "B":
		res = B()
	case "C":
		// long sleep here
		time.Sleep(time.Second * 2)
		res = "C after 2s"
	}
	fmt.Println(res)
	return res
}

func A() string {
	return "A"
}

// maybe http request
func B() string {
	time.Sleep(time.Second * 2)
	return "B after 2s"
}

func Logic2(cond string) string {
	return cond
}
