package book

type BookInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Discount    int    `json:"discount"`
	Rating      int    `json:"rating"`
}
