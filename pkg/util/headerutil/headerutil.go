package headerutil

import (
	"strconv"
	"strings"
)

// AcceptedLanguage struct
// Language and q-factor weighting
type AcceptedLanguage struct {
	Lang string
	Q    float64
}

// ParseAcceptLanguage (acceptLanguage)
// Use to parse browser 'Accept-Language' header including accepted language and q-factor weighting
// e. g. 'fr-CH, fr;q=0.9, en;q=0.8, de;q=0.7, *;q=0.5'
func ParseAcceptLanguage(acceptLanguage string) []AcceptedLanguage {
	var acceptedLanguages []AcceptedLanguage
	for _, langs := range strings.Split(acceptLanguage, ",") {
		langQ := strings.Split(strings.Trim(langs, " "), ";")
		if len(langQ) == 1 {
			acceptedLanguages = append(acceptedLanguages, AcceptedLanguage{Lang: langQ[0], Q: 1.0})
		} else {
			q, err := strconv.ParseFloat(strings.Split(langQ[1], "=")[1], 64)
			if err != nil {
				q = 0.0
			}
			acceptedLanguages = append(acceptedLanguages, AcceptedLanguage{Lang: langQ[0], Q: q})
		}
	}
	return acceptedLanguages
}

func AcceptLanguage(acceptLanguage string) string {
	var lang string
	var q float64 = 0
	language := ParseAcceptLanguage(acceptLanguage)
	for _, v := range language {
		if v.Q > q {
			q = v.Q
			lang = v.Lang
		}
	}
	return lang
}
