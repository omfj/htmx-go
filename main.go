package main

import (
	"context"
	"database/sql"
	"strconv"
	"strings"

	"github.com/omfj/htmx-go/product"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"

	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	queries := product.New(db)

	engine := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Static("/", "./static")

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main")
	})

	app.Get("/products", func(c *fiber.Ctx) error {
		search := c.Query("q")

		data := []product.Product{}

		products, err := queries.ListProducts(ctx)

		if err != nil {
			return err
		}

		if search != "" {
			for _, product := range products {
				if strings.Contains(strings.ToLower(product.Name), strings.ToLower(search)) {
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

	app.Get("/products/:id", func(c *fiber.Ctx) error {
		idQuery := c.Params("id")

		id, err := strconv.Atoi(idQuery)

		if err != nil {
			return err
		}

		product, err := queries.GetProduct(ctx, int32(id))

		if err != nil {
			return err
		}

		return c.Render("product", fiber.Map{
			"Product": product,
		}, "layouts/main")
	})

	app.Listen(":8080")
}
