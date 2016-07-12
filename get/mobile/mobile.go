package mobile

import "github.com/ArchieT/trmstac/get"

func ParseAll(s *string) (uz get.UnzipStaLs, oneerr error) {
	var jeden, drugi error
	uz, jeden, drugi = get.ParseAll(s)
	if jeden != nil {
		oneerr = jeden
	} else {
		oneerr = drugi
	}
	return
}

func ZipUzS(uz *(get.UnzipStaLs)) (as []get.AllSta, err error) { return uz.Zip() }

const THE_URL = get.THE_URL

func GoHTTPDownloadStringFromURL(url string) (s string, err error) {
	return get.DownloadStringFromURL(url)
}

func GoHTTPDownloadString() (s string, err error) { return get.DownloadString() }
