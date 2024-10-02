package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func request(apiUrl string) {
    resp, err := http.Get(apiUrl)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    var result map[string]interface{}
    err = json.Unmarshal(body, &result)
    if err != nil {
        log.Fatal("Ошибка при декодировании JSON:", err)
    }

    hourly, ok := result["hourly"].(map[string]interface{})
    if !ok {
        log.Fatal("Не удалось извлечь данные hourly")
    }

    times, ok := hourly["time"].([]interface{})
    if !ok {
        log.Fatal("Не удалось извлечь временные метки")
    }
    temperatures, ok := hourly["temperature_2m"].([]interface{})
    if !ok {
        log.Fatal("Не удалось извлечь данные по температуре")
    }

    var targetDate string

    fmt.Println("Введите дату и время в формате ГГГГ-ММ-ДД")
    fmt.Scan(&targetDate)

    getWeatherForDate(times, temperatures, targetDate)
}
