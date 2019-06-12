package entity

const IssuesURL = "https://api.delitime.ru/api/v1/cars?with=fuel,model"

type Car struct {
	ID    int       `json:"id"`
	Lat   float64   `json:"lat"`
	Lon   float64   `json:"lon"`
	Geom  []float64 `json:"-" gorm:"-"`
	Fuel  string    `json:"fuel"`
	Model `json:"model"`
}

func (Car) TableName() string {
	return "dely"
}

type Model struct {
	//	ID             int    `json:"-"`
	Name           string  `json:"name"`
	Year           int     `json:"year"`
	EngineCapacity float64 `json:"engine_capacity"`
	EnginePower    int     `json:"engine_power"`
	Transmission   string  `json:"transmission"`
	Equipment      string  `json:"equipment"`
	Img            string  `json:"img"`
	NameFull       string  `json:"name_full"`
	ImgThumb       string  `json:"img_thumb"`
}

type Dely struct {
	Cars    []Car `json:"cars"`
	Success bool  `json:"success"`
}
