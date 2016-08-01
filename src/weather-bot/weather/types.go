package weather

type Resp struct {
	Base  string `json:"base"`
	Cod   int `json:"cod"`
	Coord *Coord `json:"coord"`
	Main *Main `json:"main"`
	Name string `json:"name"`
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