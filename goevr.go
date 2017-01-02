package goevr

import (
	"strings"
)

type EVR struct {
	Epoch string
	Version string
	Release string
}

func Parse(str string) (e, v, r string) {
	e, vr := GetEAndVR(str)
	v, r = GetVAndR(vr)
	return
}


func GetEAndVR(str string) (e, vr string) {
	evr := strings.SplitN(str, ":", 2)
	if len(evr) == 1 {
		e = "0"
		vr = evr[0]
		return
	}

	e = evr[0]
	vr = evr[1]

	return
}

func GetVAndR(str string) (v, r string) {

	vrs := strings.Split(str, "-")
	count := len(vrs)

	switch {
	case count == 1:
		v = vrs[0]
	case count == 2:
		v = vrs[0]
		r = vrs[1]
	case count > 2:

		for i := 0; i < count-1; i++ {
			v+=vrs[i]+"-"
		}
		v = strings.TrimRight(v, "-")
		r = vrs[count-1]
	}

	return
}
