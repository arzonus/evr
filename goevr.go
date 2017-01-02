package goevr

import (
	"strings"
	"strconv"
)

type EVR struct {
	Epoch   string
	Version string
	Release string
}

func New(str string) EVR {
	e, v, r := Parse(str)
	return EVR{
		Epoch:   e,
		Version: v,
		Release: r,
	}
}

func (evr EVR) compare(aevr EVR) (result int) {
	result = compareFragments(evr.Epoch, aevr.Epoch)
	if result != 0 {
		return
	}

	result = compareFragments(evr.Version, aevr.Version)
	if result != 0 {
		return
	}

	result = compareFragments(evr.Release, aevr.Release)
	return
}

// evr greater than aevr
func (evr EVR) GT(aevr EVR) bool {
	return evr.compare(aevr) > 0
}

func (evr EVR) GTE(aevr EVR) bool {
	return evr.compare(aevr) >= 0
}

func (evr EVR) LT(aevr EVR) bool {
	return evr.compare(aevr) < 0
}

func (evr EVR) LTE(aevr EVR) bool {
	return evr.compare(aevr) <= 0
}

func (evr EVR) EQ(aevr EVR) bool {
	return evr.compare(aevr) == 0
}

func (evr EVR) NE(aevr EVR) bool {
	return evr.compare(aevr) != 0
}


type segment struct {
	Elements string
	IsNum bool
}


func GT(a, b string) bool {
	return compareFragments(a, b) > 0
}

func GTE(a, b string) bool {
	return compareFragments(a, b) >= 0
}

func LT(a, b string) bool {
	return compareFragments(a, b) < 0
}

func LTE(a, b string) bool {
	return compareFragments(a, b) <= 0
}

func EQ(a, b string) bool {
	return compareFragments(a, b) == 0
}

func NE(a, b string) bool {
	return compareFragments(a, b) != 0
}

// separate string for segments with num and letters
func getSegments(str string) (segs []segment){

	var seg segment
	var count = len(str)
	_, err := strconv.Atoi(string(str[0]))
	seg.Elements = string(str[0])
	seg.IsNum = err == nil


	if count <= 1 {
		segs = append(segs, seg)
		return
	}

	var _segs []segment

	for i := 1; i < count; i++ {

		_, err := strconv.Atoi(string(str[i]))

		if (err != nil && !seg.IsNum) ||
			 (err == nil && seg.IsNum) {
			seg.Elements += string(str[i])
			continue
		}

		_segs = getSegments(str[i:count])
		break
	}

	if seg.IsNum {
		seg.Elements = strings.TrimLeft(seg.Elements, "0")
	}
	segs = append(segs, seg)

	if len(_segs) >= 0 {
		segs = append(segs, _segs...)
	}

	return
}

// comparing fragments
// algortihm -> http://blog.jasonantman.com/2014/07/how-yum-and-rpm-compare-versions/
func compareFragments(a, b string) int {

	if a == b {
		return 0
	}

	lenA := len(a)
	lenB := len(b)

	if lenB == 0 && lenA > lenB {
		return 1
	} else
	if lenA == 0 && lenB > lenA {
		return -1
	}

	var i int
	if string(a[i]) == "~" && string(b[i]) != "~" {
		return -1
	} else
	if string(a[i]) != "~" && string(b[i]) == "~" {
		return 1
	} else
	if string(a[i]) == "~" && string(b[i]) == "~" {
		i++
	}

	aSegs := getSegments(a[i:lenA])
	bSegs := getSegments(b[i:lenB])

	lenASegs := len(aSegs)
	lenBSegs := len(bSegs)

	var count int

	if lenASegs > lenBSegs {
		count = lenBSegs
	} else {
		count = lenASegs
	}

	for i := 0; i < count; i++ {

		if aSegs[i].IsNum != bSegs[i].IsNum {
			if aSegs[i].IsNum && !bSegs[i].IsNum {
				return 1
			} else
			if !aSegs[i].IsNum && bSegs[i].IsNum {
				return -1
			}
		}

		if aSegs[i].IsNum {

			aNum, _ := strconv.Atoi(aSegs[i].Elements)
			bNum, _ := strconv.Atoi(bSegs[i].Elements)

			if aNum > bNum {
				return 1
			} else
			if aNum < bNum {
				return -1
			}

			continue
		}

		res := strings.Compare(aSegs[i].Elements, bSegs[i].Elements)
		if res == 0 {
			continue
		}

		return res
	}

	return 0
}

// Parse string to epoch, version and release
func Parse(str string) (e, v, r string) {
	e, vr := getEAndVR(str)
	v, r = getVAndR(vr)
	return
}

// Parse string to epoch and version with release string
func getEAndVR(str string) (e, vr string) {
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

// Parse string to version and release string
func getVAndR(str string) (v, r string) {

	vrs := strings.Split(str, "-")
	count := len(vrs)

	switch {
	case count == 1:
		v = vrs[0]
	case count == 2:
		v = vrs[0]
		r = vrs[1]
	case count > 2:
		// case for multiple '-' symbol.
		// standard describe n-e:v-r.a
		// so, release is a string after last '-'
		for i := 0; i < count-1; i++ {
			v += vrs[i] + "-"
		}
		v = strings.TrimRight(v, "-")
		r = vrs[count - 1]
	}

	return
}
