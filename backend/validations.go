package main

import "strconv"

var validators = map[string]func(string) bool{
	"require": require,
}

var validatorsWithParam = map[string]func(string, string) bool{
	"min": min,
}

func require(toCheck string) bool {
	return len(toCheck) > 0
}

func min(toCheck string, minL string) bool {
	minLength, err := strconv.ParseInt(minL, 10, 16)
	if err != nil {
		return false
	}
	return len(toCheck) >= int(minLength)
}
