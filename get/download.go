package get

import (
	"net/http"
	"time"
	"bytes"
)

func Download() (*[27]Sta,time.Time,error) {
	url := "http://trm24.pl/panel-trm/maps.jsp"
	cza := time.Now()
	response,err:=http.Get(url)
	if err!=nil {
		return nil,cza,err
	}
	defer response.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	pagestr := buf.String()
	parsed, parserr := pars(&pagestr)
	//if parserr!=nil {
	//star := Shot{parsed,cza,parserr}
//	star.giveme()
	return parsed,cza,parserr
}
