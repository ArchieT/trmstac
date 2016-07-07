package get

import (
	"github.com/ArchieT/trmstac/stadata"
	"regexp"
	"strconv"
	"sync"
)

type Sta struct{ Num, Row, Wol uint8 }

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

func pars(skad *string) (lista []Sta, err error) {
	lista = make([]Sta, 0, stadata.ILOSCSTA)
	resall := rall.FindAllString(*skad, -1)
	var wg sync.WaitGroup
	for j := range resall {
		wg.Add(1)
		go func(frag *string) {
			defer wg.Done()
			ressta := rsta.FindString(*frag)
			resrow := rrow.FindString(*frag)
			reswol := rwol.FindString(*frag)
			//resloc := rloc.FindString(frag)
			//reslat := rlat.FindString(resloc)
			//reslon := rlon.FindString(resloc)
			resintsta := rint.FindString(ressta)
			resintrow := rint.FindString(resrow)
			resintwol := rint.FindString(reswol)
			osta, erronintsta := strconv.Atoi(resintsta)
			orow, erronintrow := strconv.Atoi(resintrow)
			owol, erronintwol := strconv.Atoi(resintwol)
			nsta := Sta{uint8(osta), uint8(orow), uint8(owol)}
			lista = append(lista, nsta)
			switch {
			case erronintsta != nil:
				err = erronintsta
			case erronintrow != nil:
				err = erronintrow
			case erronintwol != nil:
				err = erronintwol
			}
		}(&resall[j])

	}
	wg.Wait()
	return lista, err
}
