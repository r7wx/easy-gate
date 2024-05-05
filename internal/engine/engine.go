package engine

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"

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

	var styleData []byte

	if status.Theme.CustomCss != "" {
		data, _ := os.ReadFile(status.Theme.CustomCss)
		styleData = data
	} else {
		data, _ := staticFS.ReadFile("static/styles/style.css")
		styleData = data
	}

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

	app.Get("/style.css", func(c *fiber.Ctx) error {

		tmpl, err := template.New("").Parse(string(styleData))
		if err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}

		c.Set("Content-type", "text/css")
		return tmpl.Execute(c, fiber.Map{
			"Background": status.Theme.Background,
			"Foreground": status.Theme.Foreground,
		})
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
			"Title": status.Title,
			"Data":  data,
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
