package main

/*


import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (cfg *ApiConfig) GetPrice(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		fmt.Println("couldn't decode params")
		return
	}

	id, err := cfg.DB.GetItemByName(r.Context(), params.Body)

	if err != nil {
		fmt.Println("couldnt find the skin")
	}

	price, err := cfg.DB.GetPricebyId(r.Context(), id)

	if err != nil {
		fmt.Println("no price for this skin")
		return
	}

	fmt.Println(price)

}

dbURL := os.Getenv("DATABASE_URL")
	db, _ := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	type Item struct {
		Name string `json:"name"`
	}
	app.Post("/v1/api/price", func(c *fiber.Ctx) error {

		p := new(Item)

		if err := c.BodyParser(p); err != nil {
			return err
		}

		id, err := dbQueries.GetItemByName(c.Context(), p.Name)

		if err != nil {
			fmt.Println("couldnt find the skin")
		}

		price, err := dbQueries.GetPricebyId(c.Context(), id)

		if err != nil {
			panic("no price")
		}
		log.Info(id)
		log.Info(price)
		return c.JSON(price)
	})

	app.Listen(":8080")
*/
