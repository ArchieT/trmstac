package get

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type StaData struct {
	Num  uint8
	Addr string
}

var rdall = regexp.MustCompile(`<a href="javascript:google\.maps\.event\.trigger\(gmarkers\[(?P<gmarkersindex>\d{1,2})\],'click'\);"><b>Stacja nr\. (?P<stacnumber>\d{1,2})\. (?P<address>[^\a\f\t\n\r\v\<\>]{5,}?)</b></a> ?<br>`)

func (d *Downloaded) ParseInfoIntoAddrList() (lista []string, err error) {
	par, err := d.ParseData()
	if err != nil {
		return
	}
	lista = make([]string, 0, len(par))
	var ii uint8
	for i := 1; i <= len(par); i++ {
		ii = uint8(i)
		juz := false
		for _, j := range par {
			if !juz && j.Num == ii {
				juz = true
				lista = append(lista, j.Addr)
			}
		}
		if !juz {
			err = errors.New("nie ma " + strconv.Itoa(i))
			return
		}
	}
	return
}

func (d *Downloaded) ParseData() (lista []StaData, err error) {
	return parsdata(&(d.Content))
}

func parsdata(skad *string) (lista []StaData, err error) {
	resall := rdall.FindAllStringSubmatch(*skad, -1)
	fmt.Println(resall)
	lista = make([]StaData, 0, 30)
	for j := range resall {
		var nasz StaData
		ri := make(map[string]int, 3)
		fmt.Println(rdall.SubexpNames())
		for i, name := range rdall.SubexpNames() {
			if i == 0 {
				continue
			}
			ri[name] = i
		}
		fmt.Println(ri)
		gmarkers, gerr := strconv.Atoi(resall[j][ri["gmarkersindex"]])
		stacnum, serr := strconv.Atoi(resall[j][ri["stacnumber"]])
		nasz.Num = uint8(stacnum)
		nasz.Addr = resall[j][ri["address"]]
		fmt.Println(nasz)
		lista = append(lista, nasz)
		if gerr != nil {
			err = gerr
			fmt.Println("g", err)
			return
		}
		if serr != nil {
			err = serr
			fmt.Println("s", err)
			return
		}
		if gmarkers+1 != stacnum {
			err = errors.New("gmarkers = " + strconv.Itoa(gmarkers) + " and stacnum = " + strconv.Itoa(stacnum) + " are not the same")
			fmt.Println("neq", err)
			return
		}
	}
	return
}
