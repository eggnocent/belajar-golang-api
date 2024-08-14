package book

type BookInput struct {
	Title string `json:"title" binding:"required"`
	Stok  int    `json:"stok" binding:"required"`
}
