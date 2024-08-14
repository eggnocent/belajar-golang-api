package book

type BookRequest struct {
	Title string `json:"title" binding:"required"`
	Stok  int    `json:"stok" binding:"required"`
}
