package get

import (
	"net/http"
	"time"
	"bytes"
)

type shot struct{
	stali [27]sta
	cza int64
	err error
}

func download() shot {
	url := "http://trm24.pl/panel-trm/maps.jsp"
	cza := time.Now().Unix()
	response,err:=http.Get(url)
	if err!=nil {
		var zlalista [27]sta
		for o:=0;o<27;o++ {
			zlalista[o]=sta{-1,-1,-1}
		}
		return shot{zlalista,cza,err}
	}
	defer response.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	pagestr := buf.String()
	parsed, parserr := pars(&pagestr)
	//if parserr!=nil {
	star := shot{parsed,cza,parserr}

//	star.giveme()

	return star
}
