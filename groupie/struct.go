package groupie

type PageDataArtice struct {
	All          Artist
	allLocation  ResponseData
	allDates     ResponseData2
	allrelations ResponseData3
}

type PageData struct {
	All          []Artist
}

type Artist struct {
	ID           int         `json:"id"`
	Image        string      `json:"image"`
	Name         string      `json:"name"`
	Members      []string    `json:"members"`
	CreationDate interface{} `json:"creationDate"`
	FirstAlbum   string      `json:"firstAlbum"`
}

type ResponseData struct {
	Index []Index `json:"index"`
}

type Index struct {
	ID      int      `json:"id"`
	TheData []string `json:"locations"`
}

type ResponseData2 struct {
	Index2 []Index2 `json:"index"`
}

type Index2 struct {
	ID      int      `json:"id"`
	TheData []string `json:"Dates"`
}

type ResponseData3 struct {
	Index3 []Index3 `json:"index"`
}

type Index3 struct {
	ID      int         `json:"id"`
	TheData interface{} `json:"datesLocations"`
}
