package util_2

import (
	"time"
)

// khai báo stuct đại diện cho pkg -> ko mock được khi test

type Utils struct{}

var u *Utils

func GetUtils() *Utils {
	if u == nil {
		u = &Utils{}
	}
	return u
}

func (obj *Utils) Logic(cond string) string {
	res := "default"

	switch cond {
	case "A":
		res = obj.A()
	case "B":
		res = obj.B()
	case "C":
		// long sleep here
		time.Sleep(time.Second * 2)
		res = "C after 2s"
	}
	return res
}

func (obj *Utils) A() string {
	return "A"
}

func (obj *Utils) B() string {
	time.Sleep(time.Second * 2)
	return "B after 2s"
}
