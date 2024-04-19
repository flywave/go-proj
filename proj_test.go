package proj

import (
	"fmt"
	"math"
	"testing"

	"github.com/flywave/go3d/float64/vec3"
)

func TestLatlongToMerc(t *testing.T) {
	ll, err := NewProj("+proj=latlong +datum=WGS84")
	defer ll.Close()
	if err != nil {
		t.Error(err)
	}

	merc, err := NewProj("+proj=merc +ellps=clrk66 +lat_ts=33")
	defer merc.Close()
	if err != nil {
		t.Error(err)
	}

	x, y, err := Transform2(ll, merc, DegToRad(-16), DegToRad(20.25))
	if err != nil {
		t.Error(err)
	} else {
		s := fmt.Sprintf("%.2f %.2f", x, y)
		s1 := "-1495284.21 1920596.79"
		if s != s1 {
			t.Errorf("LatlongToMerc = %v, want %v", s, s1)
		}
	}

	x1 := []float64{-16, -10, 0, 30.4}
	y1 := []float64{20.25, 25, 0, 40.8}
	for i := range x1 {
		x1[i] = DegToRad(x1[i])
		y1[i] = DegToRad(y1[i])
	}
	x2, y2, err := Transform2lst(ll, merc, x1, y1)
	if err != nil {
		t.Error(err)
	} else {
		s := fmt.Sprintf("[%.2f %.2f] [%.2f %.2f] [%.2f %.2f] [%.2f %.2f]", x2[0], y2[0], x2[1], y2[1], x2[2], y2[2], x2[3], y2[3])
		s1 := "[-1495284.21 1920596.79] [-934552.63 2398930.20] [0.00 0.00] [2841040.00 4159542.20]"
		if s != s1 {
			t.Errorf("LatlongToMerc = %v, want %v", s, s1)
		}
	}
}

func TestInvalidErrorProblem(t *testing.T) {
	ll, err := NewProj("+proj=latlong +datum=WGS84")
	defer ll.Close()
	if err != nil {
		t.Error(err)
	}

	merc, err := NewProj("+proj=geocent +datum=WGS84 +units=m +no_defs")
	defer merc.Close()
	if err != nil {
		t.Error(err)
	}

	_, _, _, err = Transform3(ll, merc, DegToRad(3000), DegToRad(500), 0)
	if err == nil {
		t.Error("err should not be nil")
	}

	merc2, err := NewProj("+proj=merc +ellps=clrk66 +lat_ts=33")
	defer merc2.Close()
	if err != nil {
		t.Error(err)
	}
}

const (
	SRS_3857 = "+proj=merc +a=6378137 +b=6378137 +lat_ts=0.0 +lon_0=0.0 +x_0=0.0 +y_0=0 +k=1.0 +units=m +nadgrids=@null +no_defs +over"
	SRS_4326 = "+proj=longlat +ellps=WGS84 +datum=WGS84 +no_defs +over"
)

func TestProj(t *testing.T) {
	pj1, _ := NewProj(SRS_4326)
	pj2, _ := NewProj(SRS_3857)
	trans, _ := NewTransformation(pj1, pj2)
	xs, ys, _, _ := trans.Transform([]float64{DegToRad(117), DegToRad(118)}, []float64{DegToRad(36), DegToRad(37)}, nil)
	fmt.Println(xs, ys)
}

func TestUp(t *testing.T) {
	x, y, z, _ := Lonlat2Ecef(118, 36, 30)
	org := vec3.T{x, y, z}
	org.Normalize()
	mt := GetUpRotation(x, y, z)
	v := vec3.T{0, -1, 0}
	v1 := mt.MulVec3(&v)
	v1.Normalize()
	cos := vec3.Dot(&org, &v1)
	el := math.Acos(cos) * Rad2Deg
	fmt.Println(el)
}
