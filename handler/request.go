package handler

import (
	"backend/model"

	"github.com/gofiber/fiber/v2"
)

type bookCreateRequest struct {
	Book struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
	} `json:"book"`
}

func (r *bookCreateRequest) bind(c *fiber.Ctx, b *model.Book, v *Validator) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	b.Name = r.Book.Name
	b.Description = r.Book.Description
	return nil
}

type bookUpdateRequest struct {
	Book struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
	} `json:"book"`
}

func (r *bookUpdateRequest) populate(b *model.Book) {
	r.Book.Name = b.Name
	r.Book.Description = b.Description
}

func (r *bookUpdateRequest) bind(c *fiber.Ctx, b *model.Book, v *Validator) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	b.Name = r.Book.Name
	b.Description = r.Book.Description
	return nil
}
