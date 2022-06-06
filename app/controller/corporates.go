package controller

type Corporate struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type CreateCorporatesRequest struct {
	Corporates []Corporate `json:"corporates" binding:"required"`
}
