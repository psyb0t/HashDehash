package utils

import "strings"

func GetPathParams(path string) []string {
	return strings.Split(path[1:], "/")[1:]
}
