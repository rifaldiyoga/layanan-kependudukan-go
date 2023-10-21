package keluarga

type GetKeluargaDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateKeluargaInput struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}
