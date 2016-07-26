package get

import (
	"strconv"
	"sync"
)

type Sta struct {
	Num uint8 `json:"stanum" bson:"stanum"`
	Row uint8 `json:"dostrow" bson:"dostrow"`
	Wol uint8 `json:"wolrow" bson:"wolrow"`
}

type Location struct {
	Lat float64 `json:"lat" bson:"lat"`
	Lon float64 `json:"lon" bson:"lon"`
}

type LocSta struct {
	Num      uint8 `json:"num" bson:"num"`
	Location `json:"location" bson:"location"`
}

var raliw = make(map[string]int, 30)

var NRSTA, DOSTROW, WOLROW, LATIND, LONIND int

func init() {
	for i, name := range rall.SubexpNames() {
		if i == 0 {
			continue
		}
		raliw[name] = i
	}
	NRSTA = raliw["nrsta"]
	DOSTROW = raliw["dostrow"]
	WOLROW = raliw["wolrow"]
	LATIND = raliw["lat"]
	LONIND = raliw["lon"]
}

func pars(skad *string, withloc bool) (lista []Sta, locs []LocSta, err error) {
	lista = make([]Sta, 0, 30)
	if withloc {
		locs = make([]LocSta, 0, 30)
	}
	resall := rall.FindAllStringSubmatch(*skad, -1)
	var wg sync.WaitGroup
	czekamy := make([]chan bool, len(resall)+1)
	for j := range resall {
		czekamy[j] = make(chan bool)
	}
	czekamy[len(resall)] = make(chan bool)
	go func() { czekamy[0] <- true }()
	for j := range resall {
		wg.Add(1)
		go func() { <-czekamy[len(resall)] }()
		go func(j int) {
			defer wg.Done()
			osta, erronintsta := strconv.Atoi(resall[j][NRSTA])
			orow, erronintrow := strconv.Atoi(resall[j][DOSTROW])
			owol, erronintwol := strconv.Atoi(resall[j][WOLROW])
			uosta := uint8(osta)
			nsta := Sta{uosta, uint8(orow), uint8(owol)}
			if withloc {
				var nloc LocSta
				nloc.Num = uosta
				var erronflolat, erronflolon error
				nloc.Location.Lat, erronflolat = strconv.ParseFloat(resall[j][LATIND], 64)
				nloc.Location.Lon, erronflolon = strconv.ParseFloat(resall[j][LONIND], 64)
				<-czekamy[uosta-1]
				locs = append(locs, nloc)
				switch {
				case err != nil:
				case erronflolat != nil:
					err = erronflolat
				case erronflolon != nil:
					err = erronflolon
				}
			}
			if !withloc {
				<-czekamy[uosta-1]
			}
			lista = append(lista, nsta)
			switch {
			case err != nil:
			case erronintsta != nil:
				err = erronintsta
			case erronintrow != nil:
				err = erronintrow
			case erronintwol != nil:
				err = erronintwol
			}
			czekamy[uosta] <- true
		}(j)

	}
	wg.Wait()
	return lista, locs, err
}
