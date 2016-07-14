package mobile

import "github.com/ArchieT/trmstac/get"

var UnzipStaLs get.UnzipStaLs

func GiveUnzipStaLs() get.UnzipStaLs { return UnzipStaLs }

var AllSta []get.AllSta

type MobAS get.AllSta

func GiveAllSta() []get.AllSta { return AllSta }

func GiveAllStaLen() int              { return len(AllSta) }
func GiveAllStaItem(i int) get.AllSta { return AllSta[i] }
func GiveMobASItem(i int) MobAS       { return MobAS(GiveAllStaItem(i)) }

var WewnString string

func GiveWewnString() string  { return WewnString }
func TakeWewnString(s string) { WewnString = s }

func ParseAll() string {
	var jeden, drugi error
	UnzipStaLs, jeden, drugi = get.ParseAll(&WewnString)
	if jeden != nil {
		return jeden.Error()
	}
	return drugi.Error()
}

func ZipUzS() string { var err error; AllSta, err = UnzipStaLs.Zip(); return err.Error() }

const THE_URL = get.THE_URL

/*
func GoHTTPDownloadStringFromURL(url string) string {
	var err error
	WewnString, err = get.DownloadStringFromURL(url)
	return err.Error()
}

func GoHTTPDownloadString() string {
	var err error
	WewnString, err = get.DownloadString()
	return err.Error()
}
*/
