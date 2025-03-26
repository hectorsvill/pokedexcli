package main

import (
	"regexp"
	"strings"
)

func removeDoubleSpace(str string) string {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(str, " ")
}

func cleanInput(text string) []string {
	str := strings.ToLower(text)
	str = strings.Trim(str, " ")
	str = removeDoubleSpace(str)
	return strings.Split(str, " ")
}
