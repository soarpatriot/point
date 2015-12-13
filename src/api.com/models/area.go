package models

type Area struct {
	Id        int
	Latitude  float64
	Longitude float64
	Label     string
	Code      string
	Price     string
	Points    []*Point
}
