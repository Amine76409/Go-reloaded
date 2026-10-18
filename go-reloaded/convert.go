package main

import "strconv"

func ConvertNumber(word string, base int) string {
	number, err := strconv.ParseInt(word, base, 64)
	if err != nil {
		return word
	}
	return strconv.FormatInt(number, 10)
}
