package customer

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/pgtype"
)

// Database table
type customer struct {
	FirstName pgtype.Text `json:"name" `
	ID        pgtype.UUID `json:"id" `
	Info      pgtype.JSON `json:"info" `
}

type createCustomer struct {
	Name string `json:"name" `
	Info string `json:"info" `
}

func Create(c *fiber.Ctx) error {

	customer := new(createCustomer)
	if err := c.BodyParser(customer); err != nil {
		return err
	}

	fmt.Println(customer)
	// Doesnt work yet first need to implement companies
	//if _, err := database.DB.Exec(context.Background(), `insert into public.customer (name, info) values ($1, $2)`, customer.Name, customer.Info); err != nil {
	//	return err
	//}
	return c.Status(200).JSON(fiber.Map{"success": "customer created"})
}

func GetAllCustomer(c *fiber.Ctx) error {
	var customer []customer
	// Doesnt work yet first need to implement companies
	//if err := database.DB.QueryRow(context.Background(),`select * from public.customer`).Scan(&customer); err != nil {
	//	return err
	//}
	return c.Status(200).JSON(fiber.Map{"customers": customer})
}
