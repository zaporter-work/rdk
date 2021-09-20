package main

import (
	"github.com/kellydunn/golang-geo"

	"github.com/adrianmo/go-nmea"
)

func ToPoint(a nmea.GLL) *geo.Point {
	return geo.NewPoint(a.Latitude, a.Longitude)
}

func Bearing(a, b nmea.GLL) float64 {
	aa := geo.NewPoint(a.Latitude, a.Longitude)
	bb := geo.NewPoint(b.Latitude, b.Longitude)

	return aa.BearingTo(bb)
}
