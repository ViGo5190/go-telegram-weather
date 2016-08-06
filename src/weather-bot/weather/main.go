package weather

import (
	"net/http"
	"strconv"
	"encoding/json"
	"log"
)

func floatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func TempToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 1, 64)
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func GetWeather(weatherToken string, lat, lon float64) (*Resp,error) {
	latStr := floatToString(lat)
	lonStr := floatToString(lon)
	requestUrl := "http://api.openweathermap.org/data/2.5/weather?lat=" + latStr + "&lon=" + lonStr + "&appid=" + weatherToken + "&units=metric"

	resp := new(Resp)
	err := getJson(requestUrl, resp);

	//if (err != nil) {
	//	log.Panic(err)
	//}
	//log.Println(resp.Base)
	//log.Println(resp.Cod)
	//log.Println(resp.Coord)
	//log.Println(resp.Coord.Lat)
	//log.Println(resp.Coord.Lon)
	//
	//log.Println(resp.Main)
	log.Println(resp.Main.Temp)
	//log.Println(resp.Main.Pressure)
	//log.Println(resp.Main.Humidity)

	//log.Println(resp.Name)

	return resp, err
}