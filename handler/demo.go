package handler

import (
	"demo/utils/util_1"
	"demo/utils/util_2"
	"demo/utils/util_3"
	"demo/utils/util_4"
)

// use Logic func
func One(cond string) string {
	r := util_1.Logic(cond)
	return r + " demo"
}

// use func GetUtils to get struct
func Two(cond string) string {
	r := util_2.GetUtils().Logic(cond)
	return r + " demo"
}

// use global var UtilsLogic
func Three(cond string) string {
	r := util_3.UtilsLogic(cond)
	return r + " demo"
}

// use func GetUtils to get struct
func Four(cond string) string {
	r := util_4.Logic(cond)
	return r + " demo"
}
