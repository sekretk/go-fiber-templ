package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
	}
	api := h.router.Group("/api")
	api.Get("/hello", h.home)
	api.Get("/error", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	// // "{{.Count}} - users"
	// tmpl := template.Must(template.ParseFiles("./html/page.html"))
	// // if err != nil {
	// // 	return fiber.NewError(fiber.StatusBadRequest, "template error")
	// // }
	// var tpl bytes.Buffer
	// if err := tmpl.Execute(&tpl, data); err != nil {
	// 	return fiber.NewError(fiber.StatusBadRequest, "data error")
	// }
	// c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	// return c.Send(tpl.Bytes())
	// data := struct{ Count int }{Count: 1}
	// return c.Render("page", data)
	return c.Render("page", fiber.Map{
		"Count":   5,
		"IsAdmin": false,
		"CanUse":  true,
	})
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	h.customLogger.Info().
		Bool("VAl", false).
		Int("ID", 10).
		Msg("Info")
	return c.SendString("Error\n")
}
