package engine

import (
	"embed"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/r7wx/easy-gate/internal/routine"
)

var (
	//go:embed template/*
	templateFS embed.FS
	//go:embed static/*
	staticFS embed.FS
)

// Engine - Easy Gate engine struct
type Engine struct {
	Routine *routine.Routine
}

// NewEngine - Create a new engine
func NewEngine(routine *routine.Routine) *Engine {
	return &Engine{routine}
}

// Serve - Serve application
func (e Engine) Serve() {
	status, _ := e.Routine.GetStatus()

	htmlEngine := html.NewFileSystem(http.FS(templateFS), ".html")
	app := fiber.New(fiber.Config{
		Views:                 htmlEngine,
		DisableStartupMessage: true,
	})

	app.Use(logger.New(logger.Config{
		Format:     "${time} ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "2006/01/02 15:04:05",
	}))

	app.Get("/favicon.ico", func(c *fiber.Ctx) error {
		data, err := staticFS.ReadFile("static/favicon.ico")
		if err != nil {
			return c.SendStatus(http.StatusNotFound)
		}

		c.Set("Content-type", "image/x-icon")
		return c.Send(data)
	})

	app.Get("/roboto-regular.ttf", func(c *fiber.Ctx) error {
		data, err := staticFS.ReadFile("static/font/roboto-regular.ttf")
		if err != nil {
			return c.SendStatus(http.StatusNotFound)
		}

		c.Set("Content-type", "font/ttf")
		return c.Send(data)
	})

	app.Get("/link.svg", func(c *fiber.Ctx) error {
		data, err := staticFS.ReadFile("static/link.svg")
		if err != nil {
			return c.SendStatus(http.StatusNotFound)
		}

		c.Set("Content-type", "image/svg+xml")
		return c.Send(data)
	})

	app.Get("/note.svg", func(c *fiber.Ctx) error {
		data, err := staticFS.ReadFile("static/note.svg")
		if err != nil {
			return c.SendStatus(http.StatusNotFound)
		}

		c.Set("Content-type", "image/svg+xml")
		return c.Send(data)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		status, err := e.Routine.GetStatus()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.SendString(err.Error())
		}

		addr := getAddr(status, c)
		data := getData(status, addr)

		return c.Render("template/index", fiber.Map{
			"Title":      status.Title,
			"Background": status.Theme.Background,
			"Foreground": status.Theme.Foreground,
			"Data":       data,
		})
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.Redirect("/")
	})

	if status.UseTLS {
		log.Println("Listening for connections on", status.Addr, "(HTTPS)")
		if err := app.ListenTLS(status.Addr, status.CertFile,
			status.KeyFile); err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Listening for connections on", status.Addr)
	if err := app.Listen(status.Addr); err != nil {
		log.Fatal(err)
	}
}
