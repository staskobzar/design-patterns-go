package main

import (
	"fmt"
	"strings"
)

type WeatherData struct {
	temperature float32
	description string
}

type GeoLocationService struct{}

func (l *GeoLocationService) getCoordinates(city string) (float32, float32) {
	switch strings.ToLower(city) {
	case "paris":
		return 48.8566, 2.3522
	case "new york":
		return 40.7128, -74.0060
	default:
		return 0, 0
	}
}

type WeatherService struct{}

func (w *WeatherService) fetchWeather(_, _ float32) *WeatherData {
	return &WeatherData{temperature: 18.5, description: "Cloudy"}
}

type ReportService struct{}

func (*ReportService) format(city string, data *WeatherData) {
	fmt.Printf("[+] the weather in %s is %s, temperature %f\n", city, data.description, data.temperature)
}

type WeatherFacade struct {
	geoloc *GeoLocationService
	wtr    *WeatherService
	report *ReportService
}

func NewFacade() *WeatherFacade {
	return &WeatherFacade{
		geoloc: &GeoLocationService{},
		wtr:    &WeatherService{},
		report: &ReportService{},
	}
}

func (f *WeatherFacade) getWeather(city string) {
	lon, lat := f.geoloc.getCoordinates(city)
	wdata := f.wtr.fetchWeather(lon, lat)
	f.report.format(city, wdata)
}

func main() {
	fmt.Println("===== FACADE =====")

	f := NewFacade()
	f.getWeather("Paris")
	f.getWeather("New York")
}
