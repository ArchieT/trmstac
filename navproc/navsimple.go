package navproc

type FTStaEntrInterface interface {
	TimeSec() uint16
}

type FTStaWFunc func(FTStaEntrInterface) uint16

//type FTStaEntrInterfaceMatrix [27][27]FTStaEntrInterface

func ShortestPath(g [27][27]FTStaEntrInterface, w FTStaWFunc, from,to uint8) (uint16,[]uint8) {
	var d [27]uint16
	var p [27]uint8
	q := make([]uint8, 27)
	var v uint8
	for v=0;v<27;v++ {
		d[v] = ^uint16(0)
		p[v] = ^uint8(0)
		q[v] = v
	}
	d[from] = 0
	for len(q)>0 {
		u,q := q[0],q[:1]
		for v=0;v<27;v++ {
			if d[v] > d[u] + w((g[u][v])) {
				d[v] = d[u] + w((g[u][v]))
				p[v] = u
				q = append(q,v)
			}
		}
	}
	poprzednicy:=make([]uint8,0,27)
	niemafrom := true
	poprzednicy = append(poprzednicy,to)
	for niemafrom {
		poprzednicy = append(poprzednicy,p[poprzednicy[len(poprzednicy)-1]])
		if poprzednicy[len(poprzednicy)-1] == from {
			niemafrom = false
		}
	}
	for i := len(poprzednicy)/2-1; i >= 0; i-- {
		opp := len(poprzednicy)-1-i
		poprzednicy[i], poprzednicy[opp] = poprzednicy[opp], poprzednicy[i]
	}
	return d[to],poprzednicy
}

func TimeSec(i FTStaEntrInterface) uint16 {
	return i.TimeSec()
}

func Oplata(i FTStaEntrInterface) uint16 {
	if i.TimeSec() < 20 {
		return uint16(0)
	} else if i.TimeSec() <60 {
		return uint16(1)
	} else if i.TimeSec() < 120 {
		return uint16(4)
	} else if i.TimeSec() < 180 {
		return uint16(9)
	} else if i.TimeSec() <240 {
		return uint16(16)
	} else if i.TimeSec() > 239 {
		return uint16((9+((i.TimeSec()-179)/60)+1)*7)
	}
	return uint16(65535)
}


