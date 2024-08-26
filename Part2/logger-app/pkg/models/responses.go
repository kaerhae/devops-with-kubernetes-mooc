package models

type ReadResponse struct {
	Hash    string
	Content string
}

type ErrorResponse struct {
	Status  int
	Message string
}

type PongResponse struct {
	Pongs int `json:"pongs"`
}
