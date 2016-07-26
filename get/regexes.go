package get

import "regexp"

//Google Maps array row regex — bike counts and locations
var rall = regexp.MustCompile(`Stacja nr ? (?P<nrsta>\d{1,2}) ? ? ? ? ?</br> ? ? ? ? ? ?Dostępne rowery: (?P<dostrow>\d{1,2}) ? ? ? ?</br> ? ? ? ?Wolne sloty (?P<wolrow>\d{1,2}) ', (?P<lat>\d+\.\d+) , (?P<lon>\d+\.\d+) , 'http`)

//Clickable links for stations — addresses
var rdall = regexp.MustCompile(`<a href="javascript:google\.maps\.event\.trigger\(gmarkers\[(?P<gmarkersindex>\d{1,2})\],'click'\);"><b> ? ? ? ?Stacja nr\. (?P<stacnumber>\d{1,2})\. (?P<address>[^\a\f\t\n\r\v\<\>]{5,}?) {0,5}?</b> ? ? ? ?</a> ? ? ? ?<[Bb]r>`)
