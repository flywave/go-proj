package proj

/*
#include "proj_c_api.h"

#cgo CFLAGS:-I ./lib

#cgo linux LDFLAGS: -L ./lib/linux  -Wl,--start-group -lproj -ldl -lm -lcproj -Wl,--end-group
#cgo darwin LDFLAGS: -L /usr/lib -lc++ -L ./lib/darwin  -lproj -lcproj -lm
#cgo darwin,arm LDFLAGS: -L /usr/lib -lc++ -L ./lib/darwin_arm  -lproj -lcproj -lm
#cgo windows LDFLAGS: -L ./lib/windows -lproj -lcproj  -fPIC
*/
import "C"
import (
	"errors"
	"math"
	"runtime"
	"sync"
	"unsafe"
)

var (
	mu sync.Mutex
)

type Proj struct {
	pj     C.projPJ
	opened bool
}

func NewProj(definition string) (*Proj, error) {
	cs := C.CString(definition)
	defer C.free(unsafe.Pointer(cs))
	proj := &Proj{opened: false}

	mu.Lock()
	defer mu.Unlock()

	proj.pj = C.pj_init_plus(cs)

	if proj.pj == nil {
		return proj, errors.New(C.GoString(C.get_err()))
	}

	proj.opened = true
	runtime.SetFinalizer(proj, (*Proj).Close)

	return proj, nil
}

func (p *Proj) Close() {
	if p.opened {
		C.pj_free(p.pj)
		p.opened = false
	}
}

func (p *Proj) IsLatLong() bool {
	return C.pj_is_latlong(p.pj) != 0
}

func (p *Proj) IsGeoCent() bool {
	return C.pj_is_geocent(p.pj) != 0
}

func (p *Proj) GetDef() string {
	return C.GoString(C.pj_get_def(p.pj, 0))
}

func (p *Proj) CompareDatums(c *Proj) bool {
	return int(C.pj_compare_datums(p.pj, c.pj)) > 0
}

func transform(srcpj, dstpj *Proj, x, y, z []float64) ([]float64, []float64, []float64, error) {
	if !(srcpj.opened && dstpj.opened) {
		return []float64{}, []float64{}, []float64{}, errors.New("projection is closed")
	}

	var x1, y1, z1 []C.double

	ln := len(x)
	if len(y) < ln {
		ln = len(y)
	}
	if z != nil && len(z) < ln {
		ln = len(z)
	}

	if ln == 0 {
		return []float64{}, []float64{}, []float64{}, nil
	}

	x1 = make([]C.double, ln)
	y1 = make([]C.double, ln)
	if z != nil {
		z1 = make([]C.double, ln)
	}
	for i := 0; i < ln; i++ {
		x1[i] = C.double(x[i])
		y1[i] = C.double(y[i])
		if z != nil {
			z1[i] = C.double(z[i])
		}
	}

	var e *C.char
	if z != nil {
		e = C.transform(srcpj.pj, dstpj.pj, C.long(ln), &x1[0], &y1[0], &z1[0])
	} else {
		e = C.transform(srcpj.pj, dstpj.pj, C.long(ln), &x1[0], &y1[0], nil)
	}

	if e != nil {
		return []float64{}, []float64{}, []float64{}, errors.New(C.GoString(e))
	}

	var x2, y2, z2 []float64
	x2 = make([]float64, ln)
	y2 = make([]float64, ln)
	if z != nil {
		z2 = make([]float64, ln)
	}
	for i := 0; i < ln; i++ {
		x2[i] = float64(x1[i])
		y2[i] = float64(y1[i])
		if z != nil {
			z2[i] = float64(z1[i])
		}
	}
	return x2, y2, z2, nil
}

func Transform2(srcpj, dstpj *Proj, x, y float64) (float64, float64, error) {
	xx, yy, _, err := transform(srcpj, dstpj, []float64{x}, []float64{y}, nil)
	if err != nil {
		return math.NaN(), math.NaN(), err
	}
	return xx[0], yy[0], err
}

func Transform3(srcpj, dstpj *Proj, x, y, z float64) (float64, float64, float64, error) {
	xx, yy, zz, err := transform(srcpj, dstpj, []float64{x}, []float64{y}, []float64{z})
	if err != nil {
		return math.NaN(), math.NaN(), math.NaN(), err
	}
	return xx[0], yy[0], zz[0], err
}

func Transform2lst(srcpj, dstpj *Proj, x, y []float64) ([]float64, []float64, error) {
	xx, yy, _, err := transform(srcpj, dstpj, x, y, nil)
	return xx, yy, err
}

func Transform3lst(srcpj, dstpj *Proj, x, y, z []float64) ([]float64, []float64, []float64, error) {
	return transform(srcpj, dstpj, x, y, z)
}

// Longitude and latitude in degrees
func Fwd(proj *Proj, long, lat float64) (x, y float64, err error) {
	if !proj.opened {
		return math.NaN(), math.NaN(), errors.New("projection is closed")
	}
	x1 := C.double(long)
	y1 := C.double(lat)
	e := C.fwd(proj.pj, &x1, &y1)
	if e != nil {
		return math.NaN(), math.NaN(), errors.New(C.GoString(e))
	}
	return float64(x1), float64(y1), nil
}

// Longitude and latitude in degrees
func Inv(proj *Proj, x, y float64) (long, lat float64, err error) {
	if !proj.opened {
		return math.NaN(), math.NaN(), errors.New("projection is closed")
	}
	x2 := C.double(x)
	y2 := C.double(y)
	e := C.inv(proj.pj, &x2, &y2)
	if e != nil {
		return math.NaN(), math.NaN(), errors.New(C.GoString(e))
	}
	return float64(x2), float64(y2), nil
}

func DegToRad(deg float64) float64 {
	return deg / 180.0 * math.Pi
}

func RadToDeg(rad float64) float64 {
	return rad / math.Pi * 180.0
}

type Transformation struct {
	src *Proj
	dst *Proj
}

func NewTransformation(src, dst *Proj) (Transformation, error) {
	return Transformation{src, dst}, nil
}

func (t Transformation) Transform(xs, ys, zs []float64) ([]float64, []float64, []float64, error) {
	return transform(t.src, t.dst, xs, ys, zs)
}
