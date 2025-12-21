package home

import (
	"boy/go-fiber-templ/pkg/tadaptor"
	"boy/go-fiber-templ/views"

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
	// api := h.router.Group("/")
	h.router.Get("/", h.home)
	h.router.Get("/404", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	component := views.Main()

	return tadaptor.Render(c, component)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	h.customLogger.Info().
		Bool("VAl", false).
		Int("ID", 10).
		Msg("Info")
	return c.SendString("Error\n")
}
