package entities

type ApplyDTO struct {
	Date      string `json:"date"`
	Name      string `json:"name"`
	Hakbun    int    `json:"hakbun"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Portfolio string `json:"portfolio"`
	Reason    string `json:"reason"`
}

type Apply struct {
	Uuid      string
	Date      string
	Name      string
	Hakbun    int
	Phone     string
	Email     string
	Portfolio string
	Reason    string
}
