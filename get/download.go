package get

import (
	"net/http"
	"strconv"
	"time"
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
		error
	}
	else if stacja<10 {
		stastr = "00" + stais + "TOR"
	}
	else {
		stastr = "0" + stais + "TOR"
	}
	var url string
	url = "http://trm24.pl/panel-trm/" + stastr + ".jsp"

	cza = time.Now().Unix()

	response,get:=http.Get(url)
	if err!=nil {
		return
	}
	defer response.Body.Close()

	sta := ile{pars(response.Body),stacja,stastr,cza}

	sta.giveme()

	return sta
}
