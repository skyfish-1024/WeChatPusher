package model

type WeatherResp struct {
	Code       string  `json:"code"`
	UpdateTime string  `json:"updateTime"`
	FxLink     string  `json:"fxLink"`
	Daily      []Daily `json:"daily"`
	Refer      Refer   `json:"refer"`
}
type Daily struct {
	FxDate         string `json:"fxDate"`
	Sunrise        string `json:"sunrise"`
	Sunset         string `json:"sunset"`
	Moonrise       string `json:"moonrise"`
	Moonset        string `json:"moonset"`
	MoonPhase      string `json:"moonPhase"`
	MoonPhaseIcon  string `json:"moonPhaseIcon"`
	TempMax        string `json:"tempMax"`
	TempMin        string `json:"tempMin"`
	IconDay        string `json:"iconDay"`
	TextDay        string `json:"textDay"`
	IconNight      string `json:"iconNight"`
	TextNight      string `json:"textNight"`
	Wind360Day     string `json:"wind360Day"`
	WindDirDay     string `json:"windDirDay"`
	WindScaleDay   string `json:"windScaleDay"`
	WindSpeedDay   string `json:"windSpeedDay"`
	Wind360Night   string `json:"wind360Night"`
	WindDirNight   string `json:"windDirNight"`
	WindScaleNight string `json:"windScaleNight"`
	WindSpeedNight string `json:"windSpeedNight"`
	Humidity       string `json:"humidity"`
	Precip         string `json:"precip"`
	Pressure       string `json:"pressure"`
	Vis            string `json:"vis"`
	Cloud          string `json:"cloud"`
	UvIndex        string `json:"uvIndex"`
}
type Refer struct {
	Sources []string `json:"sources"`
	License []string `json:"license"`
}
