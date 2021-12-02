package Utils

import "regexp"

func FindInText(target, patern string) []string {
	paterntext := regexp.MustCompile(patern)
	cleartext := paterntext.FindAllString(target, -1)
	return cleartext
}
