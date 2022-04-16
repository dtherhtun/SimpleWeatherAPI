# Weather API with Go

## About

This is a simple weather API writing with Golang for Gobootcamp and Using [OpenWeatherMap](https://openweathermap.org)

## Go Version

```
go1.17
```

## How To Usage
Before run or build the main.go, please add API key from OpenWeatherMap in line number 18.
Example:
```
const key="abcdefghijk123456789"
```

```
go run main.go
```
or
```
go build .
./SimpleWeatherAPI
```

Open the following URL in Browser and check the weather status using `location` parameter by city or city,countrycode.

- http://localhost:8080/weather?location=<city,countrycode>
- Example - http://localhost:8080/weather?location=Rangoon,mm