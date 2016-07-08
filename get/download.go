package get

import (
	"bytes"
	"errors"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func Download() (d Downloaded, err error) {
	url := "http://trm24.pl/panel-trm/maps.jsp"
	d.Time = time.Now()
	response, err := http.Get(url)
	if err == nil {
		defer response.Body.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		d.Content = buf.String()
	}
	return
}

func (d *Downloaded) ParseSta() (staout []Sta, err error) {
	staout, _, err = pars(&(d.Content), false)
	return
}

func (d *Downloaded) ParseStaWithLoc() ([]Sta, []LocSta, error) { return pars(&(d.Content), true) }

type UnzipStaLs struct {
	StaL     []Sta
	LocStaL  []LocSta
	StaDataL []StaData
}

type AllSta struct {
	Sta
	LocSta
	StaData
}

func (d *Downloaded) ParseAll() (uz UnzipStaLs, slocerr, dataerr error) {
	uz.StaL, uz.LocStaL, slocerr = d.ParseStaWithLoc()
	uz.StaDataL, dataerr = d.ParseData()
	return
}

func inteqall(a ...int) bool {
	for i := range a {
		for j := range a {
			if j != i && a[j] != a[i] {
				return false
			}
		}
	}
	return true
}

func genfromplacetolen(from, tolen int) fromplacetolen {
	return fromplacetolen{from: from, to: tolen - 1, c: from, ok: true, psd: false}
}

type fromplacetolen struct {
	from, to, c int
	ok, psd     bool
}

func (fptl *fromplacetolen) next() {
	fptl.c += 1
	if !fptl.psd {
		if fptl.c > fptl.to {
			fptl.c = 0
			fptl.psd = true
		}
	} else if fptl.c >= fptl.from {
		fptl.ok = false
	}
}

func (uz *UnzipStaLs) Zip() (as []AllSta, err error) {
	if !inteqall(len(uz.StaL), len(uz.LocStaL), len(uz.StaDataL)) {
		err = errors.New("not equal lenghts")
	}
	as = make([]AllSta, len(uz.StaL))
	var wg sync.WaitGroup
	placeret := func(ylen int, yname string, givenum func(int) uint8, place func(int)) func(uint8) {
		ourforreturn := func(n uint8) {
			defer wg.Done()
			ni := int(n - 1)
			for o := genfromplacetolen(ni, ylen); o.ok; o.next() {
				if givenum(o.c) == n {
					place(o.c)
					return
				}
			}
			err = errors.New(yname + " found no " + strconv.Itoa(int(n)))
		}
		return ourforreturn
	}
	giveplacestal := func(ind int) uint8 { return uz.StaL[ind].Num }
	giveplaceloc := func(ind int) uint8 { return uz.LocStaL[ind].Num }
	giveplacedata := func(ind int) uint8 { return uz.StaDataL[ind].Num }
	helperplacestal := func(ind int) { as[ind].Sta = uz.StaL[ind] }
	helperplaceloc := func(ind int) { as[ind].LocSta = uz.LocStaL[ind] }
	helperplacedata := func(ind int) { as[ind].StaData = uz.StaDataL[ind] }
	placestal := placeret(len(uz.StaL), "placestal", giveplacestal, helperplacestal)
	placeloc := placeret(len(uz.LocStaL), "placeloc", giveplaceloc, helperplaceloc)
	placedata := placeret(len(uz.StaDataL), "placedata", giveplacedata, helperplacedata)
	ulen := uint8(len(uz.StaL))
	for i := uint8(1); i <= ulen; i++ {
		wg.Add(3)
		go placestal(i)
		go placeloc(i)
		go placedata(i)
	}
	wg.Wait()
	return
}
