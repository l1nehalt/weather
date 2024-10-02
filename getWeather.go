package main

import (
	"fmt"
	"log"
	"time"
)

func getWeatherForDate(times []interface{}, temperatures []interface{}, targetDate string) {
	for i, t := range times {
		timestamp, ok := t.(string)
		if !ok {
			continue
		}

		parsedTime, err := time.Parse("2006-01-02T15:04", timestamp)
		if err != nil {
			log.Println("Ошибка при парсинге времени:", err)
			continue
		}

		if parsedTime.Format("2006-01-02") == targetDate {
			fmt.Printf("Время: %s, Температура: %.2f°C\n", parsedTime, temperatures[i].(float64))
		}
	}
}
