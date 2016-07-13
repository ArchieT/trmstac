package main

import "fmt"
import "github.com/ArchieT/trmstac/get"

func main() {
	s, e := get.DownloadString()
	if e != nil {
		fmt.Println(len(s), e, "\n", s)
	} else {
		fmt.Println("lenght: ", len(s))
		fmt.Println(s)
	}
}
