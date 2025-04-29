package weather

// 见文档：https://lbs.amap.com/api/webservice/guide/api-advanced/weatherinfo
type AllInfo struct {
	Status    string `json:"status"`
	Count     string `json:"count"`
	Info      string `json:"info"`
	Infocode  string `json:"infocode"`
	Forecasts []struct {
		City       string `json:"city"`
		Adcode     string `json:"adcode"`
		Province   string `json:"province"`
		Reporttime string `json:"reporttime"`
		Casts      []struct {
			Date           string `json:"date"`
			Week           string `json:"week"`
			Dayweather     string `json:"dayweather"`
			Nightweather   string `json:"nightweather"`
			Daytemp        string `json:"daytemp"`
			Nighttemp      string `json:"nighttemp"`
			Daywind        string `json:"daywind"`
			Nightwind      string `json:"nightwind"`
			Daypower       string `json:"daypower"`
			Nightpower     string `json:"nightpower"`
			DaytempFloat   string `json:"daytemp_float"`
			NighttempFloat string `json:"nighttemp_float"`
		} `json:"casts"`
	} `json:"forecasts"`
}
type BaseInfo struct {
	Status   string `json:"status"`
	Count    string `json:"count"`
	Info     string `json:"info"`
	Infocode string `json:"infocode"`
	Lives    []struct {
		Province         string `json:"province"`
		City             string `json:"city"`
		Adcode           string `json:"adcode"`
		Weather          string `json:"weather"`
		Temperature      string `json:"temperature"`
		Winddirection    string `json:"winddirection"`
		Windpower        string `json:"windpower"`
		Humidity         string `json:"humidity"`
		Reporttime       string `json:"reporttime"`
		TemperatureFloat string `json:"temperature_float"`
		HumidityFloat    string `json:"humidity_float"`
	} `json:"lives"`
}
