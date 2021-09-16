package proj

import (
	"math"

	pj "github.com/flywave/go-proj"
	dmat4 "github.com/flywave/go3d/float64/mat4"
	dvec3 "github.com/flywave/go3d/float64/vec3"
)

const (
	Deg2Rad = math.Pi / 180
	Rad2Deg = 180 / math.Pi
)

func Lonlat2Ecef(lon, lat, z float64) (float64, float64, float64, error) {
	wgs84Proj, _ := pj.NewProj("+proj=longlat +datum=WGS84 +no_defs")
	ecefProj, _ := pj.NewProj("+proj=geocent +datum=WGS84 +units=m +no_defs")
	return pj.Transform3(wgs84Proj, ecefProj, lon*Deg2Rad, lat*Deg2Rad, z)
}

func Ecef2Lonlat(x, y, z float64) (float64, float64, float64, error) {
	wgs84Proj, _ := pj.NewProj("+proj=longlat +datum=WGS84 +no_defs")
	ecefProj, _ := pj.NewProj("+proj=geocent +datum=WGS84 +units=m +no_defs")
	lon, lat, z, err := pj.Transform3(ecefProj, wgs84Proj, x, y, z)
	return lon * Rad2Deg, lat * Rad2Deg, z, err
}

func Lonlat2Merc(lon, lat float64) (float64, float64, error) {
	wgs84Proj, _ := pj.NewProj("+proj=longlat +datum=WGS84 +units=m +no_defs")
	mercProj, _ := pj.NewProj("+proj=merc +a=6378137 +b=6378137 +lat_ts=0.0 +lon_0=0.0 +x_0=0.0 +y_0=0.0 +k=1.0 +units=m +nadgrids=@null +wktext +no_defs +over")
	return pj.Transform2(wgs84Proj, mercProj, lon*Deg2Rad, lat*Deg2Rad)
}

func Merc2Lonlat(x, y float64) (float64, float64, error) {
	wgs84Proj, _ := pj.NewProj("+proj=longlat +datum=WGS84 +units=m +no_defs")
	mercProj, _ := pj.NewProj("+proj=merc +a=6378137 +b=6378137 +lat_ts=0.0 +lon_0=0.0 +x_0=0.0 +y_0=0.0 +k=1.0 +units=m +nadgrids=@null +wktext +no_defs +over")
	lon, lat, err := pj.Transform2(mercProj, wgs84Proj, x, y)
	return lon * Rad2Deg, lat * Rad2Deg, err
}

func GetUpRotation(x, y, z float64) *dmat4.T {
	eye := dvec3.T{x, y, z}
	target := dvec3.T{0, 0, 0}
	up := dvec3.T{0, 0, 1}

	_z := dvec3.Sub(&eye, &target)

	if _z.LengthSqr() == 0 {
		// eye and target are in the same position
		_z[2] = 1
	}

	_z.Normalize()
	_x := dvec3.Cross(&up, &_z)

	if _x.LengthSqr() == 0 {
		// up and z are parallel
		if math.Abs(float64(up[2])) == 1 {
			_z[0] += 0.0001
		} else {
			_z[2] += 0.0001
		}

		_z.Normalize()
		_x = dvec3.Cross(&up, &_z)
	}

	_x.Normalize()
	_y := dvec3.Cross(&_z, &_x)

	te := dmat4.Ident
	te[0][0] = _x[0]
	te[0][1] = _y[0]
	te[0][2] = _z[0]
	te[1][0] = _x[1]
	te[1][1] = _y[1]
	te[1][2] = _z[1]
	te[2][0] = _x[2]
	te[2][1] = _y[2]
	te[2][2] = _z[2]
	te.Transpose()
	return &te
}
