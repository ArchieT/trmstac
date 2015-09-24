package get

import (
	"fmt"
	"regexp"
	"strconv"
)

type sta struct {num,row,wol int}

//rall := regexp.MustCompile(`Stacja nr\s \d+\s+</br>\s+Dostępne rowery: \d+\s+</br>\s+Wolne sloty \d+ ', \d+\.\d+ , \d+\.\d+ , 'http:`)
var rall = regexp.MustCompile(`Stacja nr\s \d+\s+</br>\s+Dostępne rowery: \d+\s+</br>\s+Wolne sloty \d+ ',`)
var rsta = regexp.MustCompile(`Stacja nr\s \d+\s+</br>\s+Dostępne rowery:`)
var rrow = regexp.MustCompile(`</br>\s+Dostępne rowery: \d+\s+</br>\s+Wolne sloty`)
var rwol = regexp.MustCompile(`</br>\s+Wolne sloty \d+ ',`)
//rloc := regexp.MustCompile(`', \d+\.\d+ , \d+\.\d+ , 'http:`)
//rlat := regexp.MustCompile(`', \d+\.\d+ , `)
//rlon := regexp.MustCompile(` , \d+\.\d+ , 'http:`)
var rint = regexp.MustCompile(`\d+`)
//rflo := regexp.MustCompile(`\d+\.\d+`)

func singpars(frag *string, nlista *[27]sta, enasz *error) {
	fmt.Println(*frag)
	ressta := rsta.FindString(*frag)
	fmt.Println(ressta)
	resrow := rrow.FindString(*frag)
	fmt.Println(resrow)
	reswol := rwol.FindString(*frag)
	fmt.Println(reswol)
	//resloc := rloc.FindString(frag)
	//reslat := rlat.FindString(resloc)
	//reslon := rlon.FindString(resloc)
	resintsta := rint.FindString(ressta)
	resintrow := rint.FindString(resrow)
	resintwol := rint.FindString(reswol)
	osta, erronintsta := strconv.Atoi(resintsta)
	orow, erronintrow := strconv.Atoi(resintrow)
	owol, erronintwol := strconv.Atoi(resintwol)
	nsta := sta{osta,orow,owol}
	nlista[osta-1] = nsta
	switch {
	case erronintsta!=nil:
		*enasz=erronintsta
	case erronintrow!=nil:
		*enasz=erronintrow
	case erronintwol!=nil:
		*enasz=erronintwol
	}
}


func pars(skad *string) ([27]sta, error) {
	var lista [27]sta
	var errnasz error
	resall := rall.FindAllString(*skad,-1)
		for j,_ := range resall {go singpars(&resall[j], &lista, &errnasz)}
		return lista, errnasz
}
