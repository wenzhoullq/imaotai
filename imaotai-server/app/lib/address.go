package lib

import "math"

func CalDis(lat1, lng1, lat2, lng2 float64) float64 {
	const EarthRadius = 6371000
	φ1 := lat1 * math.Pi / 180 // 转换纬度 to rad
	φ2 := lat2 * math.Pi / 180
	Δφ := (lat2 - lat1) * math.Pi / 180
	Δλ := (lng2 - lng1) * math.Pi / 180

	a := math.Sin(Δφ/2)*math.Sin(Δφ/2) +
		math.Cos(φ1)*math.Cos(φ2)*
			math.Sin(Δλ/2)*math.Sin(Δλ/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := EarthRadius * c // 单位：米
	return distance
}
