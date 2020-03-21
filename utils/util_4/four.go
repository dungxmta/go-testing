package util_4

import (
	"time"
)

// gán func A, B làm thuộc tính của struct -> mock được A, B khi test

type Utils struct {
	A func() string
	B func() string
}

var u *Utils

func GetUtils() *Utils {
	if u == nil {
		u = &Utils{
			A: A,
			B: B,
		}
	}
	return u
}

func A() string {
	return "A"
}

func B() string {
	time.Sleep(time.Second * 2)
	return "B after 2s"
}

func Logic(cond string) string {
	res := "default"

	switch cond {
	case "A":
		res = GetUtils().A()
	case "B":
		res = GetUtils().B()
	case "C":
		// long sleep here
		time.Sleep(time.Second * 2)
		res = "C after 2s"
	}
	return res
}
