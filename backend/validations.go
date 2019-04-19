package main

import (
	"fmt"
	"strconv"
)

var validators = map[string]func(string) error{
	"required": required,
}

var validatorsWithParam = map[string]func(string, string) error{
	"min": min,
	"max": max,
}

func required(toCheck string) (err error) {
	if len(toCheck) == 0 {
		err = fmt.Errorf("the param '%s' does not match rule required", toCheck)
	}
	return
}

func min(toCheck string, minL string) (err error) {
	minLength, err := strconv.ParseInt(minL, 10, 16)
	if err != nil {
		return
	}
	if len(toCheck) < int(minLength) {
		err = fmt.Errorf("the param '%v' does not match rule min:%v", toCheck, minL)
	}
	return
}

func max(toCheck string, maxL string) (err error) {
	maxLength, err := strconv.ParseInt(maxL, 10, 16)
	if err != nil {
		return
	}
	if len(toCheck) > int(maxLength) {
		err = fmt.Errorf("the param '%v' does not match rule max:%v", toCheck, maxL)
	}
	return
}
