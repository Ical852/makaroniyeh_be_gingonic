package product

type GetByCatInput struct {
	ID int `uri:"id" binding:"required"`
}

type GetByNameInput struct {
	Name string `uri:"name" binding:"required"`
}