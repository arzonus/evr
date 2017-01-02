package goevr

import "testing"

func TestGetEAndVRWithoutEpoch(t *testing.T) {

	const (
		evrWithoutEpoch = `1.12.1+dfsg-19+deb8u2`
		epoch           = `0`
		versionrelease  = `1.12.1+dfsg-19+deb8u2`
	)

	e, vr := GetEAndVR(evrWithoutEpoch)

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

	e, vr := GetEAndVR(evrWithEpoch)

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

	v, r := GetVAndR(vr)
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

	v, r := GetVAndR(vr)
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

	v, r := GetVAndR(vr)
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