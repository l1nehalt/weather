package main

import "fmt"

var apiUrlRyz = "https://api.open-meteo.com/v1/forecast?latitude=54.6269&longitude=39.6916&hourly=temperature_2m"
var apiUrlMos = "https://api.open-meteo.com/v1/forecast?latitude=55.7522&longitude=37.6156&hourly=temperature_2m"
var apiUrlLukh = "https://api.open-meteo.com/v1/forecast?latitude=54.9766&longitude=39.0444&hourly=temperature_2m"
var apiUrlKol = "https://api.open-meteo.com/v1/forecast?latitude=55.0794&longitude=38.7783&hourly=temperature_2m"

func main() {
	var town int
	flag := true
	for flag {
		printMenu()
		fmt.Println("Выберите номер города:")
		fmt.Scan(&town)
		switch town {
		case 1:
			request(apiUrlRyz)
		case 2:
			request(apiUrlMos)
		case 3:
			request(apiUrlLukh)
		case 4:
			request(apiUrlKol)
		case 5:
			flag = false
		default:
			fmt.Println("Такого города нет. Пожалуйста, выберите корректный номер.")
		}
	}
}
