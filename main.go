package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

var products = []string{
	"Apple AirPods Pro",
	"Apple iPhone 11 Pro Max",
	"Apple iPhone 11",
	"Apple iPhone XS Max",
	"Apple iPhone XS",
	"Apple iPhone XR",
	"Apple iPhone X",
	"Apple iPhone 8 Plus",
	"Apple iPhone 8",
	"Huawei P30 Pro",
	"Huawei P30",
	"Huawei Mate 20 Pro",
	"Huawei Mate 20",
	"Huawei P20 Pro",
	"Huawei P20",
	"Huawei Mate 10 Pro",
	"Huawei Mate 10",
	"LG G8X ThinQ",
	"LG G8S ThinQ",
	"LG V50S ThinQ",
	"LG V50 ThinQ",
	"LG V40 ThinQ",
	"LG V35 ThinQ",
}

func main() {
	engine := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./static")

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main")
	})

	app.Get("/products", func(c *fiber.Ctx) error {
		search := c.Query("q")

		data := []string{}

		if search != "" {
			for _, product := range products {
				if strings.Contains(strings.ToLower(product), strings.ToLower(search)) {
					data = append(data, product)
				}
			}
		} else {
			data = products
		}

		return c.Render("fragments/products", fiber.Map{
			"Products": data,
		})
	})

	app.Listen(":8080")
}
