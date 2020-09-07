package models

//	Smartphone model  for smartphone bean
type Smartphone struct {
	Id            int64
	Name          string
	Price         int
	CountryOrigin string
	Os            string
}

// CreateSmartphoneCMD command to create Smartphones
type CreateSmartphoneCMD struct {
	Name          string `json:"name"`
	Price         int    `json:"price"`
	CountryOrigin string `json:"contryorigin"`
	Os            string `json:"os"`
}
