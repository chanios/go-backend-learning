package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) Register(r *fiber.App) {
	v1 := r.Group("/api/v1")
	books := v1.Group("/books")
	books.Get("", h.GetBooks)
	books.Get("/:id", h.GetBook)
	books.Delete("/:id", h.DeleteBook)
	books.Post("", h.CreateBook)
	books.Put("/:id", h.UpdateBook)
}
