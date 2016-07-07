package get

import (
	"errors"
	"regexp"
	"strconv"
)

type StaData struct {
	Num  uint8
	Addr string
}

var rdall = regexp.MustCompile(`<a href="javascript:google\.maps\.event\.trigger\(gmarkers[(?P<gmarkersindex>\d{1,2})],'click'\);"><b>Stacja nr\. (?P<stacnumber>)\d{1,2}\. (?P<address>[^\a\f\t\n\r\v\<\>]{5,}?)</b></a> ?<br>`)

func (d *Downloaded) ParseData() (lista []StaData, err error) {
	return parsdata(&(d.Content))
}

func parsdata(skad *string) (lista []StaData, err error) {
	resall := rdall.FindAllStringSubmatch(*skad, -1)
	lista = make([]StaData, 0, 30)
	for j := range resall {
		var nasz StaData
		ri := make(map[string]int, 3)
		for i, name := range rdall.SubexpNames() {
			if i == 0 {
				continue
			}
			ri[name] = i
		}
		gmarkers, gerr := strconv.Atoi(resall[j][ri["gmarkersindex"]])
		stacnum, serr := strconv.Atoi(resall[j][ri["stacnumber"]])
		nasz.Num = uint8(stacnum)
		nasz.Addr = resall[j][ri["address"]]
		lista = append(lista, nasz)
		if gerr != nil {
			err = gerr
			return
		}
		if serr != nil {
			err = serr
			return
		}
		if gmarkers+1 != stacnum {
			err = errors.New("gmarkers = " + strconv.Itoa(gmarkers) + " and stacnum = " + strconv.Itoa(stacnum) + " are not the same")
			return
		}
	}
	return
}