# API-Client

## Client connect to openweathermap.org API and we GET information about weather in city: temperature, temperature feels like, humidity, pressure, wind speed, wind degrees.

### https://api.openweathermap.org/data/2.5/weather?q="your city"&appid="your API-Key"

## Then we decode with function Unmarshal() the JSON into a Go value
### err = json.Unmarshal(body, &data)

