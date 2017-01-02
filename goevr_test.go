package goevr

import (
	"testing"
	"reflect"
)

func TestGetEAndVRWithoutEpoch(t *testing.T) {

	const (
		evrWithoutEpoch = `1.12.1+dfsg-19+deb8u2`
		epoch           = `0`
		versionrelease  = `1.12.1+dfsg-19+deb8u2`
	)

	e, vr := getEAndVR(evrWithoutEpoch)

	if e != epoch {
		t.Fatal("versionrelease are not equals: ", e, epoch)
	}

	if versionrelease != vr {
		t.Fatal("epoch are not equals: ", vr, versionrelease)
	}

}

func TestGetEAndVRWithEpoch(t *testing.T) {

	const (
		evrWithEpoch   = "20:1.12.1+dfsg-19+deb8u2"
		epoch          = `20`
		versionrelease = `1.12.1+dfsg-19+deb8u2`
	)

	e, vr := getEAndVR(evrWithEpoch)

	if e != epoch {
		t.Fatal("versionrelease are not equals: ", e, epoch)
	}

	if versionrelease != vr {
		t.Fatal("epoch are not equals: ", vr, versionrelease)
	}

}

func TestGetVAndRWithoutRelease(t *testing.T) {
	const (
		vr      = "1.12.1"
		version = "1.12.1"
		release = ""
	)

	v, r := getVAndR(vr)
	if v != version {
		t.Fatal("versions are not equals: ", v, version)
	}
	if r != release {
		t.Fatal("releases are not equals: ", r, release)
	}
}

func TestGetVAndRWithRelease(t *testing.T) {
	const (
		vr      = "1.12.1+dfsg-19+deb8u2"
		version = "1.12.1+dfsg"
		release = "19+deb8u2"
	)

	v, r := getVAndR(vr)
	if v != version {
		t.Fatal("versions are not equals: ", v, version)
	}
	if r != release {
		t.Fatal("releases are not equals: ", r, release)
	}
}

func TestGetVAndRWithReleaseAndDoubleVersion(t *testing.T) {
	const (
		vr      = "1.12.1+dfsg-19+deb8u2-19+deb8u2"
		version = "1.12.1+dfsg-19+deb8u2"
		release = "19+deb8u2"
	)

	v, r := getVAndR(vr)
	if v != version {
		t.Fatal("versions are not equals: ", v, version)
	}
	if r != release {
		t.Fatal("releases are not equals: ", r, release)
	}
}

func TestParseDefault(t *testing.T) {
	const (
		evr     = "1:1.12.1+dfsg-19+deb8u2"
		epoch   = "1"
		version = "1.12.1+dfsg"
		release = "19+deb8u2"
	)

	e, v, r := Parse(evr)

	if e != epoch {
		t.Fatal("epoches are not equals: ", e, epoch)
	}
	if v != version {
		t.Fatal("versions are not equals: ", v, version)
	}
	if r != release {
		t.Fatal("releases are not equals: ", r, release)
	}
}

func TestParseWithoutEpoch(t *testing.T) {
	const (
		evr     = "1.12.1+dfsg-19+deb8u2"
		epoch   = "0"
		version = "1.12.1+dfsg"
		release = "19+deb8u2"
	)

	e, v, r := Parse(evr)

	if e != epoch {
		t.Fatal("epoches are not equals: ", e, epoch)
	}
	if v != version {
		t.Fatal("versions are not equals: ", v, version)
	}
	if r != release {
		t.Fatal("releases are not equals: ", r, release)
	}
}

func TestParseWithoutRelease(t *testing.T) {
	const (
		evr     = "1.12.1+dfsg"
		epoch   = "0"
		version = "1.12.1+dfsg"
		release = ""
	)

	e, v, r := Parse(evr)

	if e != epoch {
		t.Fatal("epoches are not equals: ", e, epoch)
	}
	if v != version {
		t.Fatal("versions are not equals: ", v, version)
	}
	if r != release {
		t.Fatal("releases are not equals: ", r, release)
	}
}

func TestParseWithBigVersion(t *testing.T) {
	const (
		evr     = "1.12.1+dfsg-19+deb8u2-19+deb8u2"
		epoch   = "0"
		version = "1.12.1+dfsg-19+deb8u2"
		release = "19+deb8u2"
	)

	e, v, r := Parse(evr)

	if e != epoch {
		t.Fatal("epoches are not equals: ", e, epoch)
	}
	if v != version {
		t.Fatal("versions are not equals: ", v, version)
	}
	if r != release {
		t.Fatal("releases are not equals: ", r, release)
	}
}

func TestParseEmptyString(t *testing.T) {
	const (
		evr     = ""
		epoch   = "0"
		version = ""
		release = ""
	)

	e, v, r := Parse(evr)

	if e != epoch {
		t.Fatal("epoches are not equals: ", e, epoch)
	}
	if v != version {
		t.Fatal("versions are not equals: ", v, version)
	}
	if r != release {
		t.Fatal("releases are not equals: ", r, release)
	}
}

func TestGetSegments(t *testing.T) {
	const (
		evr     = "1.12bdfd"
	)

	var segments = []segment{
		segment{
			Elements: "1",
			IsNum: true,
		},
		segment{
			Elements: ".",
			IsNum: false,
		},
		segment{
			Elements: "12",
			IsNum: true,
		},
		segment{
			Elements: "bdfd",
			IsNum: false,
		},

	}

	if !reflect.DeepEqual(getSegments(evr), segments) {
		t.Fatal("segments are not equals: ", segments, getSegments(evr))
	}
}

func TestCompareFragments(t *testing.T) {

	const (
		a = "1.8.7"
		b = "1.8.7"
	)

	if compareFragments(a, b) != 0 {
		t.Fatal("a, b are not equals: ", a, b)
	}
}

func TestCompareGTTrue(t *testing.T) {
	const (
		a = "1.8.8"
		b = "1.8.7"
	)

	if !GT(a, b) {
		t.Fatal("a is not GT b", a, b)
	}
}

func TestCompareGTETrue(t *testing.T) {
	const (
		a = "1.8.7"
		b = "1.8.7"
	)

	if !GTE(a, b) {
		t.Fatal("a is not GTE b", a, b)
	}
}
func TestCompareLTTrue(t *testing.T) {
	const (
		a = "1.8.6"
		b = "1.8.7"
	)

	if !LT(a, b) {
		t.Fatal("a is not LT b", a, b)
	}
}
func TestCompareLTETrue(t *testing.T) {
	const (
		a = "1.8.7"
		b = "1.8.7"
	)

	if !LTE(a, b) {
		t.Fatal("a is not LTE b", a, b)
	}
}
func TestCompareEQTrue(t *testing.T) {
	const (
		a = "1.8.7"
		b = "1.8.7"
	)

	if !EQ(a, b) {
		t.Fatal("a is not EQ b", a, b)
	}
}
func TestCompareNETrue(t *testing.T) {
	const (
		a = "1.8.3"
		b = "1.8.7"
	)

	if !NE(a, b) {
		t.Fatal("a is not NE b", a, b)
	}
}