package stadata


type Vertex struct {
	Lat, Long float64
}


type sdata struct {
	Num int
	Stastr, Staaddrloc string
	Staloc Vertex
	Rowslot int
}

var List = [27]sdata{
	sdata{1,`001TOR`,`Rynek Staromiejski`,Vertex{53.010747 , 18.603927},15},
	sdata{2,`002TOR`,`Plac św. Katarzyny`,Vertex{53.013731 , 18.612896},15},
	sdata{3,`003TOR`,`Plac Rapackiego`,Vertex{53.009713 , 18.600316},25},
	sdata{4,`004TOR`,`ul. Bulwar Filadelfijski - Brama Klasztorna`,Vertex{53.008216 , 18.603825},15},
	sdata{5,`005TOR`,`ul. Szosa Chełmińska - Targowisko Miejskie`,Vertex{53.01801 , 18.600215},15},
	sdata{6,`006TOR`,`ul. Gagarina - Biblioteka Uniwersytecka`,Vertex{53.020541 , 18.570548},15},
	sdata{7,`007TOR`,`ul. Broniewskiego - Tesco`,Vertex{53.013509 , 18.571054},15},
	sdata{8,`008TOR`,`ul. Gen. Józefa Hallera - basen`,Vertex{52.995533 , 18.617013},15},
	sdata{9,`009TOR`,`ul. Szosa Chełmińska - Polo Market`,Vertex{53.042359 , 18.580434},15},
	sdata{10,`010TOR`,`PKP Toruń Główny`,Vertex{53.001739 , 18.614953},15},
	sdata{11,`011TOR`,`ul. Dziewulskiego - Komisariat Policji`,Vertex{53.027362 , 18.665074},15},
	sdata{12,`012TOR`,`ul. Konstytucji 3 Maja - Pawilon Maciej`,Vertex{53.023813 , 18.676055},15},
	sdata{13,`013TOR`,`ul. Dąbrowskiego - Dworzec autobusowy`,Vertex{53.015798 , 18.607102},15},
	sdata{14,`014TOR`,`ul. Wały gen. Sikorskiego - Urząd Miasta`,Vertex{53.013286 , 18.603526},15},
	sdata{15,`015TOR`,`ul. Gen. Józefa Bema - Tor-Tor`,Vertex{53.018912 , 18.592017},15},
	sdata{16,`016TOR`,`ul. Przysiecka - Barbarka`,Vertex{53.054141 , 18.541683},15},
	sdata{17,`017TOR`,`ul. Bażyńskich - basen`,Vertex{53.021249 , 18.612961},15},
	sdata{18,`018TOR`,`PKP Toruń Wschodni`,Vertex{53.026017 , 18.633976},15},
	sdata{19,`019TOR`,`ul. Kościuszki / ul. Świętopełka`,Vertex{53.023938 , 18.617852},15},
	sdata{20,`020TOR`,`ul. Mickiewicza / ul. Tujakowskiego`,Vertex{53.01202 , 18.596131},15},
	sdata{21,`021TOR`,`ul. Gagarina - Od Nowa`,Vertex{53.018791 , 18.580527},15},
	sdata{22,`022TOR`,`ul. Rydygiera / ul. Donimirskiego`,Vertex{53.029478 , 18.653368},15},
	sdata{23,`023TOR`,`ul. Kolankowskiego / ul. Kosynierów`,Vertex{53.028687 , 18.692435},15},
	sdata{24,`024TOR`,`ul. Św. Klemensa / ul. Św. Józefa`,Vertex{53.022939 , 18.583693},15},
	sdata{25,`025TOR`,`ul. Legionów - Rondo Czadcy`,Vertex{53.033301 , 18.599847},15},
	sdata{26,`026TOR`,`ul. Zółkiewskiego - Atrium Copernicus`,Vertex{53.024561 , 18.638458},20},
	sdata{27,`027TOR`,`ul. Grudziądzka - Kometa`,Vertex{53.034255 , 18.613291},20},
}
