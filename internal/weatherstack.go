package internal

import "time"

/*
* These structs are the response from weatherstack
* Only WeatherData has two extra fields for the cache
* servieces
 */
type WeatherData struct {
	Historical       map[string]History `json:"historical"`
	Location         Location           `json:"location"`
	Request          Request            `json:"request"`
	Current          Current            `json:"current"`
	RequestTime      time.Time
	IsCachedResponse bool `json:"isCachedResponse"`
}

type Request struct {
	Type     string `json:"type"`
	Query    string `json:"query"`
	Language string `json:"language"`
	Unit     string `json:"unit"`
}

type Location struct {
	Name           string `json:"name"`
	Country        string `json:"country"`
	Region         string `json:"region"`
	Lat            string `json:"lat"`
	Lon            string `json:"lon"`
	TimezoneID     string `json:"timezone_id"`
	Localtime      string `json:"localtime"`
	UTCOffset      string `json:"utc_offset"`
	LocaltimeEpoch int64  `json:"localtime_epoch"`
}

type Current struct {
	WindDir             string   `json:"wind_dir"`
	ObservationTime     string   `json:"observation_time"`
	WeatherDescriptions []string `json:"weather_descriptions"`
	WeatherIcons        []string `json:"weather_icons"`
	WindDegree          int      `json:"wind_degree"`
	WindSpeed           int      `json:"wind_speed"`
	WeatherCode         int      `json:"weather_code"`
	Temperature         int      `json:"temperature"`
	Pressure            int      `json:"pressure"`
	Precip              float64  `json:"precip"`
	Humidity            int      `json:"humidity"`
	Cloudcover          int      `json:"cloudcover"`
	Feelslike           int      `json:"feelslike"`
	UVIndex             int      `json:"uv_index"`
	Visibility          int      `json:"visibility"`
}

type History struct {
	Date      string   `json:"date"`
	Astro     Astro    `json:"astro"`
	Hourly    []Hourly `json:"hourly"`
	DateEpoch int64    `json:"date_epoch"`
	MinTemp   int      `json:"mintemp"`
	MaxTemp   int      `json:"maxtemp"`
	AvgTemp   int      `json:"avgtemp"`
	TotalSnow int      `json:"totalsnow"`
	SunHour   float64  `json:"sunhour"`
	UVIndex   int      `json:"uv_index"`
}

type Astro struct {
	Sunrise          string `json:"sunrise"`
	Sunset           string `json:"sunset"`
	Moonrise         string `json:"moonrise"`
	Moonset          string `json:"moonset"`
	MoonPhase        string `json:"moon_phase"`
	MoonIllumination int    `json:"moon_illumination"`
}

type Hourly struct {
	WindDir             string   `json:"wind_dir"`
	Time                string   `json:"time"`
	WeatherDescriptions []string `json:"weather_descriptions"`
	WeatherIcons        []string `json:"weather_icons"`
	DewPoint            int      `json:"dewpoint"`
	FeelsLike           int      `json:"feelslike"`
	WindDegree          int      `json:"wind_degree"`
	WindSpeed           int      `json:"wind_speed"`
	Precip              float64  `json:"precip"`
	Humidity            int      `json:"humidity"`
	Visibility          int      `json:"visibility"`
	Pressure            int      `json:"pressure"`
	Cloudcover          int      `json:"cloudcover"`
	HeatIndex           int      `json:"heatindex"`
	Temperature         int      `json:"temperature"`
	WindChill           int      `json:"windchill"`
	WindGust            int      `json:"windgust"`
	WeatherCode         int      `json:"weather_code"`
	ChanceOfRain        int      `json:"chanceofrain"`
	ChanceOfRemDry      int      `json:"chanceofremdry"`
	ChanceOfWindy       int      `json:"chanceofwindy"`
	ChanceOfOvercast    int      `json:"chanceofovercast"`
	ChanceOfSunshine    int      `json:"chanceofsunshine"`
	ChanceOfFrost       int      `json:"chanceoffrost"`
	ChanceOfHighTemp    int      `json:"chanceofhightemp"`
	ChanceOfFog         int      `json:"chanceoffog"`
	ChanceOfSnow        int      `json:"chanceofsnow"`
	ChanceOfThunder     int      `json:"chanceofthunder"`
	UVIndex             int      `json:"uv_index"`
}
