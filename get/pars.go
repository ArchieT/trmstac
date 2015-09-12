package get

import (
	"regexp"
	"strconv"
)

func pars(skad *string) int {
	refi := regexp.MustCompile("w na stacji.*szt")
	rese := regexp.MustCompile("th.*szt")
	recy := regexp.MustCompile("\D")
	if row,err := strconv.Atoi(recy.ReplaceAllString(rese.FindString(refi.FindString(skad)),"")); err==nil {
		return row
	}
}
