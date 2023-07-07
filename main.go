package main

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var Inmemory = make(map[int]User)

var CountID = 0

type User struct {
	Firstname string
	Lastname  string
}

func DeleteUsers(c *fiber.Ctx) error {
	c.AllParams()
	user := Inmemory[1]
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	if user.Firstname == user.Firstname {
		delete(Inmemory, 1)
		fmt.Println(Inmemory)
		return c.SendString("Usuário excluído com sucesso")
	}
	fmt.Println(Inmemory[1])
	return c.Status(fiber.StatusOK).JSON(user)
}

func HandleCreateUser(c *fiber.Ctx) error {
	c.AllParams()
	user := User{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	Inmemory[CountID+1] = user
	return c.Status(fiber.StatusOK).JSON(user)
}

func HandleShowUser(c *fiber.Ctx) error {
	c.AllParams()
	user := Inmemory[1]
	if user.Firstname == "" {
		return errors.New("User Not Found")
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func hello(c *fiber.Ctx) error {
	c.Send([]byte("Hello World!"))
	return nil
}

func main() {

	app := fiber.New()

	app.Get("/", hello)
	app.Get("/users/:id", HandleShowUser)
	app.Post("/users/:id", HandleCreateUser)
	app.Delete("/users/:id", DeleteUsers)

	app.Listen(":3000")
}

// Finalizar com update e list (listar) e delete.
// arrumar: pegar dados da request (passar pra struct) ok
// ID pela request
// DOC fiber ensina.
//ajustar padrão URL do fiber na func main(). ok
// ajustar names Handlers ok
