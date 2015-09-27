package main

import (
	"bytes"
	"fmt"
	"github.com/ArchieT/trmstac/get"
	"github.com/ArchieT/trmstac/stadata"
	"strconv"
)

func main(){
	a := get.Download()
//	fmt.Println(a)
	fmt.Println("Liczba rowerów na stacjach TRM: ", a.Cza)
	for ib,b := range a.Stali {
		var buffer bytes.Buffer
		buffer.WriteString(" ")
		buffer.WriteString(stadata.List[ib].Stastr)
		buffer.WriteString(" | ")
		for i:=b.Row;i>0;i-- {
			//fmt.Print("█")
			buffer.WriteString("▉")
		}
		for i:=b.Wol;i>0;i-- {
			buffer.WriteString("▒")
		}
		buffer.WriteString(" ")
		buffer.WriteString(strconv.Itoa(b.Row))
		buffer.WriteString("/")
		buffer.WriteString(strconv.Itoa(b.Row+b.Wol))
		buffer.WriteString(" (")
		buffer.WriteString(strconv.Itoa(b.Wol))
		buffer.WriteString(" empty)")
		fmt.Println(buffer.String())
	}
	fmt.Println(" ")
}
