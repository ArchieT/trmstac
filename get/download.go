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

func (d *Downloaded) ParseSta() ([]Sta, error) {
	return pars(&(d.Content))
}
