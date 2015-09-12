package get

import (
	"net/http"
	"strconv"
	"time"
	"bytes"
)

type sta struct {
	row int
	sta int
	stastr string
	cza int64
}

func download(stacja int) sta {
	stais := strconv.Itoa(stacja)
	var stastr string
	if stacja<1 || stacja>26 {
		panic
	} else if stacja<10 {
		stastr = "00" + stais + "TOR"
	} else {
		stastr = "0" + stais + "TOR"
	}
	var url string
	url = "http://trm24.pl/panel-trm/" + stastr + ".jsp"

	cza := time.Now().Unix()

	response,err:=http.Get(url)
	if err!=nil {
		return sta{-1,stacja,stastr,cza}
	}
	defer response.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	pagestr := buf.String()

	star := sta{pars(&pagestr),stacja,stastr,cza}

	star.giveme()

	return star
}
