package handler

import (
	"backend/model"
	"backend/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetBook(c *fiber.Ctx) error {
	id64, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}
	id := uint(id64)
	book, err := h.bookStore.GetByID(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusOK).JSON(newBookResponse(book))
}

func (h *Handler) GetBooks(c *fiber.Ctx) error {
	limit, err := strconv.Atoi(c.Params("limit"))
	if err != nil {
		limit = 20
	}
	books, count, err := h.bookStore.List(limit)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(nil))
	}
	return c.Status(http.StatusOK).JSON(newBookListResponse(books, count))
}

func (h *Handler) DeleteBook(c *fiber.Ctx) error {
	id64, err := strconv.ParseUint(c.Params("id"), 10, 32)
	id := uint(id64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}
	b, err := h.bookStore.GetByID(id)
	if err != nil {
		log.Println("[FATAL]")
		log.Fatal(err)
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if b == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound())
	}
	if err := h.bookStore.Delete(b); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusOK).JSON(map[string]interface{}{"result": "ok"})
}

func (h *Handler) CreateBook(c *fiber.Ctx) error {
	var b model.Book
	req := &bookCreateRequest{}
	if err := req.bind(c, &b, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	b.AuthorID = 1
	err := h.bookStore.Create(&b)

	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusCreated).JSON(newBookResponse(&b))
}

func (h *Handler) UpdateBook(c *fiber.Ctx) error {
	id64, err := strconv.ParseUint(c.Params("id"), 10, 32)
	id := uint(id64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}
	b, err := h.bookStore.GetByID(id)
	if err != nil {
		log.Println("[FATAL]")
		log.Fatal(err)
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	req := &bookUpdateRequest{}
	req.populate(b)
	if err = req.bind(c, b, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	if err = h.bookStore.Update(b); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusOK).JSON(newBookResponse(b))
}
