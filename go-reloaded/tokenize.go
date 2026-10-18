package main

import "regexp"

var tokenRegex = regexp.MustCompile(`\((?:hex|bin|up|low|cap)(?:,\s*\d+)?\)|\S+`)
var markerRegex = regexp.MustCompile(`^\((?:hex|bin|up|low|cap)(?:,\s*\d+)?\)$`)

func Tokenize(line string) []string {
	return tokenRegex.FindAllString(line, -1)
}

func IsMarker(token string) bool {
	return markerRegex.MatchString(token)
}
