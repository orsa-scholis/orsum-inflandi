package main

import (
	"fmt"
	"strconv"
)

var validators = map[string]func(string) error{
	"required": required,
	"int":      checkInt,
}

var validatorsWithParam = map[string]func(string, string) error{
	"gt": gt,
	"lt": lt,
}

func required(toCheck string) (err error) {
	if len(toCheck) == 0 {
		err = fmt.Errorf("the param '%s' does not match rule required", toCheck)
	}
	return
}

func gt(toCheck string, minL string) (err error) {
	minLength, err := strconv.Atoi(minL)
	if err != nil {
		return
	}
	if len(toCheck) < minLength {
		err = fmt.Errorf("the param '%v' does not match rule gt:%v", toCheck, minL)
	}
	return
}

func lt(toCheck string, maxL string) (err error) {
	maxLength, err := strconv.Atoi(maxL)
	if err != nil {
		return
	}
	if len(toCheck) > maxLength {
		err = fmt.Errorf("the param '%v' does not match rule lt:%v", toCheck, maxL)
	}
	return
}

func checkInt(toCheck string) (err error) {
	_, err = strconv.Atoi(toCheck)
	if err != nil {
		err = fmt.Errorf("the param '%v' does not match rule int", toCheck)
	}
	return
}
