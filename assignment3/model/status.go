package model

type Status struct {
	Status StatusBody `json:"status"`
}

type StatusBody struct {
	Water uint `json:"water"`
	Wind  uint `json:"wind"`
}
