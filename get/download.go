package get

import (
	"bytes"
	"net/http"
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
