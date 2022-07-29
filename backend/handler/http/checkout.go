package http

import (
	"github.com/gofiber/fiber/v2"
)

func (h *HttpHandler) RegisterCheckoutHandler() {
	h.f.Get("/checkout", h.getAllCheckout)
	h.f.Get("/checkout/status/:id", h.getCheckoutId)
	h.f.Post("/checkout", h.createCheckout)
	h.f.Post("/checkout/process/:id", h.moveToProcess)
	h.f.Post("/checkout/deliver/:id", h.moveToDeliver)
}

func (h *HttpHandler) getAllCheckout(c *fiber.Ctx) error {

	objs, err := h.u.GetAllCheckout()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).JSON(objs)

}

func (h *HttpHandler) getCheckoutId(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString("ID not valid")
	}

	objs, err := h.u.GetCheckoutById(uint(id))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).JSON(objs)

}

type createCheckoutBody struct {
	Name string `json:"name" xml:"name" form:"name"`
}

func (h *HttpHandler) createCheckout(c *fiber.Ctx) error {
	var obj createCheckoutBody

	if err := c.BodyParser(&obj); err != nil {
		return err
	}

	checkout, err := h.u.CreateCheckout(obj.Name)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(checkout)
}

func (h *HttpHandler) moveToProcess(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString("ID not valid")
	}

	obj, err := h.u.CheckoutGoToProcess(uint(id))
	if err != nil {
		return err
	}

	return c.Status(200).JSON(obj)
}

func (h *HttpHandler) moveToDeliver(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString("ID not valid")
	}

	obj, err := h.u.CheckoutGoToDeliver(uint(id))
	if err != nil {
		return err
	}

	return c.Status(200).JSON(obj)
}
