package common

import "strings"

func ColNotEmpty(row []string) int {
	count := 0
	for _, col := range row {
		if col != "" {
			count++
		}
	}
	return count
}

func ToCamelClass(source string) string {
	end := strings.Index(source, ".")
	a := source
	if end > -1 {
		a = source[:end]
	}
	s := strings.Split(a, "_")
	target := ""
	for _, v := range s {
		f := strings.ToUpper(v[0:1])
		e := v[1:]
		target += f + e
	}
	return target
}

func ToCamelVariable(source string) string {
	end := strings.Index(source, ".")
	a := source
	if end > -1 {
		a = source[:end]
	}
	s := strings.Split(a, "_")
	target := ""
	for i, v := range s {
		f := ""
		if i == 0 {
			f = v[0:1]
		} else {
			f = strings.ToUpper(v[0:1])
		}
		e := v[1:]
		target += f + e
	}
	return target
}
