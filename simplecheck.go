package main

import (
	"fmt"
	"github.com/ArchieT/trmstac/get"
	"github.com/ArchieT/trmstac/stadata"
)

func main(){
	a := get.Download()
//	fmt.Println(a)
	fmt.Println("Liczba rowerów na stacjach TRM: ", a.Cza)
	for ib,b := range a.Stali {
		fmt.Print(" ", stadata.List[ib].Stastr, " | ")
		for i:=b.Row;i>0;i-- {
			//fmt.Print("█")
			fmt.Print("▉")
		}
		for i:=b.Wol;i>0;i-- {
			fmt.Print("▒")
		}
		fmt.Println(" ",b.Row,"/",b.Row+b.Wol,"(",b.Wol," empty)")
	}
	fmt.Println(" ")
}
