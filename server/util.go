package main

import "math"

const EarthRadiusMeters = 6367500

func Radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func LatLngToXYZ(lat, lng float64) (x, y, z float64) {
	lat, lng = Radians(lat), Radians(lng)
	x = EarthRadiusMeters * math.Cos(lat) * math.Cos(lng)
	y = EarthRadiusMeters * math.Cos(lat) * math.Sin(lng)
	z = EarthRadiusMeters * math.Sin(lat)
	return
}
