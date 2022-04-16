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


## Sample Output
```
$ curl http://localhost:8080/weather?location=Rangoon
{"coord":{"lon":96.1561,"lat":16.8053},"weather":[{"id":802,"main":"Clouds","description":"scattered clouds","icon":"03d"}],"base":"stations","main":{"temp":309.13,"feels_like":316.13,"temp_min":309.13,"temp_max":309.13,"pressure":1004,"humidity":52},"visibility":8000,"wind":{"speed":4.12,"deg":250},"clouds":{"all":40},"dt":1650099855,"sys":{"type":1,"id":9322,"country":"MM","sunrise":1650064758,"sunset":1650109848},"timezone":23400,"id":1298824,"name":"Yangon","cod":200}
```
