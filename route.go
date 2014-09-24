package urlrouter

import (
	"regexp"
	"fmt"
	"strings"
)

type Route struct {
	Keys []string
	Regex *regexp.Regexp
	Pattern string
}

type Params map[string]string

func (route *Route) Params (url string) Params {
	match := route.Regex.FindAllStringSubmatch(url, -1)[0][1:]
	result := make(Params)

	for i := range match {
		if len(route.Keys) <= i {
			break
		}

		result[route.Keys[i]] = match[i]
	}

	return result
}

func NewRoute (pattern string) Route {
	regex, keys := PathToRegex(pattern)
	return Route{keys, regex, pattern}
}

func PathToRegex (path string) (*regexp.Regexp, []string) {
	pattern, _ := regexp.Compile(":([A-Za-z0-9]+)")
	matches := pattern.FindAllStringSubmatch(path, -1)
	keys := []string{}

	for i := range matches {
		keys = append(keys, matches[i][1])
	}

	str := fmt.Sprintf("^%s\\/?$", strings.Replace(path, "/", "\\/", -1))

	str = pattern.ReplaceAllString(str, "([^\\/]+)")
	str = strings.Replace(str, ".", "\\.", -1)

	regex, _ := regexp.Compile(str)

	return regex, keys
}
