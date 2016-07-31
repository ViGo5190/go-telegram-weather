package weather

import (
	"net/http"
	"fmt"
	"strconv"
	"encoding/json"
)

type Resp struct {
	Base  string `json:"base"`
	Cod   int `json:"cod"`
	Coord *Coord `json:"coord"`
	Main *Main `json:"main"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp     float64 `json:"temp"`
	Pressure float64 `json:"pressure"`
	Humidity float64 `json:"humidity"`
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func GetWeather(weatherToken string, lat, lon float64) *Resp{
	latStr := FloatToString(lat)
	lonStr := FloatToString(lon)
	requestUrl := "http://api.openweathermap.org/data/2.5/weather?lat=" + latStr + "&lon=" + lonStr + "&appid=" + weatherToken + "&units=metric"

	resp := new(Resp)
	getJson(requestUrl, resp);

	fmt.Println(resp.Base)
	fmt.Println(resp.Cod)
	fmt.Println(resp.Coord)
	fmt.Println(resp.Coord.Lat)
	fmt.Println(resp.Coord.Lon)

	fmt.Println(resp.Main)
	fmt.Println(resp.Main.Temp)
	fmt.Println(resp.Main.Pressure)
	fmt.Println(resp.Main.Humidity)

	return resp
}