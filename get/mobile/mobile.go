package mobile

import "github.com/ArchieT/trmstac/get"

var UnzipStaLs get.UnzipStaLs

func GiveUnzipStaLs() get.UnzipStaLs { return UnzipStaLs }

var AllSta []get.AllSta

func GiveAllSta() []get.AllSta { return AllSta }

var WewnString string

func GiveWewnString() string  { return WewnString }
func TakeWewnString(s string) { WewnString = s }

func ParseAll() error {
	var jeden, drugi error
	UnzipStaLs, jeden, drugi = get.ParseAll(&WewnString)
	if jeden != nil {
		return jeden
	}
	return drugi
}

func ZipUzS() (err error) { AllSta, err = UnzipStaLs.Zip(); return }

const THE_URL = get.THE_URL

func GoHTTPDownloadStringFromURL(url string) (err error) {
	WewnString, err = get.DownloadStringFromURL(url)
	return
}

func GoHTTPDownloadString() (err error) { WewnString, err = get.DownloadString(); return }
