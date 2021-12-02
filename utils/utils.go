package utils

import "regexp"

func findInText(target, patern string) []string {
	paterntext := regexp.MustCompile(patern)
	cleartext := paterntext.FindAllString(target, -1)
	return cleartext
}
