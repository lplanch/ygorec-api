package util

import "strings"

func ReallySplit(s string, sep string) []string {
	if len(s) == 0 {
		return []string{}
	}
	return strings.Split(s, sep)
}
