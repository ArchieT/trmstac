package get

import (
	"net/http"
	"time"
	"bytes"
)

// Shot : Stali - [27]sta , Cza time.Time , Err error
type Shot struct{
	Stali [27]sta
	Cza time.Time
	Err error
}

// Download : returns a Shot
func Download() Shot {
	url := "http://trm24.pl/panel-trm/maps.jsp"
	cza := time.Now()
	response,err:=http.Get(url)
	if err!=nil {
		var zlalista [27]sta
		var o uint8
		for o=0;o<27;o++ {
			zlalista[o]=sta{0,0,0}
		}
		return Shot{zlalista,cza,err}
	}
	defer response.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	pagestr := buf.String()
	parsed, parserr := pars(&pagestr)
	//if parserr!=nil {
	star := Shot{parsed,cza,parserr}

//	star.giveme()

	return star
}
