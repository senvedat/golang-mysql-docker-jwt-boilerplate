package models

type Armament struct {
	Title string `json:"title"`
	Qty   string `json:"qty"`
}

type Spaceship struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	Class    string     `json:"class"`
	Armament []Armament `json:"armament"`
	Crew     int        `json:"crew"`
	Image    string     `json:"image"`
	Value    float64    `json:"value"`
	Status   string     `json:"status"`
}
