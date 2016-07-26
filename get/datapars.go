package get

import (
	"errors"
	"strconv"
)

type StaData struct {
	Num  uint8  `json:"num" bson:"num"`
	Addr string `json:"addr" bson:"addr"`
}

func (d *Downloaded) ParseInfoIntoAddrList() (lista []string, err error) {
	return ParseInfoIntoAddrList(&(d.Content))
}
func ParseInfoIntoAddrList(s *string) (lista []string, err error) {
	par, err := ParseData(s)
	if err != nil {
		return
	}
	return ParseStaDataIntoAddrList(par)
}

func ParseStaDataIntoAddrList(par []StaData) (lista []string, err error) {
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
	return ParseData(&(d.Content))
}

func ParseData(skad *string) (lista []StaData, err error) {
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
