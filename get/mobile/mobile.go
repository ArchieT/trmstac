package mobile

import "github.com/ArchieT/trmstac/get"
import "strconv"

const (
	IDXROW int = iota
	IDXWOL
	IDXNUM
	IDXLAT
	IDXLON
	IDXADDR
)

var UnzipStaLs get.UnzipStaLs

func GiveUnzipStaLs() get.UnzipStaLs { return UnzipStaLs }

var AllSta []get.AllSta

//type MobAS get.AllSta

func GiveAllSta() []get.AllSta { return AllSta }

func GiveAllStaLen() int              { return len(AllSta) }
func GiveAllStaItem(i int) get.AllSta { return AllSta[i] }

func GiveASpRow(i int) uint8   { return GiveAllStaItem(i).Sta.Row }
func GiveASpWol(i int) uint8   { return GiveAllStaItem(i).Sta.Wol }
func GiveASpNum(i int) uint8   { return GiveAllStaItem(i).Sta.Num }
func GiveASpLat(i int) float64 { return GiveAllStaItem(i).LocSta.Location.Lat }
func GiveASpLon(i int) float64 { return GiveAllStaItem(i).LocSta.Location.Lon }
func GiveASpAddr(i int) string { return GiveAllStaItem(i).StaData.Addr }

func ourfmtloc(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func ourfmtuint8(u uint8) string {
	return strconv.Itoa(int(u))
}

type sixstring [6]string

func (ss sixstring) Give(i int) string {
	return ss[i]
}

type GiveByInt interface {
	Give(int) string
}

func GiveASpSTRTUP(i int) GiveByInt {
	var r sixstring
	o := GiveAllStaItem(i)
	r[IDXROW] = ourfmtuint8(o.Row)
	r[IDXWOL] = ourfmtuint8(o.Wol)
	r[IDXNUM] = ourfmtuint8(o.Sta.Num)
	r[IDXLAT] = ourfmtloc(o.Lat)
	r[IDXLON] = ourfmtloc(o.Lon)
	r[IDXADDR] = o.Addr
	return r
}

//func GiveMobASItem(i int) MobAS       { return MobAS(GiveAllStaItem(i)) }

var WewnString string

func GiveWewnString() string  { return WewnString }
func TakeWewnString(s string) { WewnString = s }

func ParseAll() string {
	var jeden, drugi error
	UnzipStaLs, jeden, drugi = get.ParseAll(&WewnString)
	if jeden != nil {
		return jeden.Error()
	}
	if drugi != nil {
		return drugi.Error()
	}
	return ""
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
