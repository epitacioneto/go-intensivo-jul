package main

import (
	"database/sql"
	"fmt"

	"github.com/epitacioneto/go-intensivo-jul/internal/infra/database"
	"github.com/epitacioneto/go-intensivo-jul/internal/usecases"
	_ "github.com/mattn/go-sqlite3"
)

type Car struct {
	Model string
	Color string
}

// metodo

func (c Car) Start() {
	println(c.Model + " has been started")
}

// func (c Car) ChangeColor(color string) {
// 	c.Color = color // duplicando o valor de c.color na memória - cópia do color original
// 	println("New color: " + c.Color)
// }

func (c *Car) ChangeColor(color string) {
	c.Color = color
	println("New color: " + c.Color)
}

// funcao

// func soma(x, y int) int {
// 	return x + y
// }

func main() {

	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}

	orderRepository := database.NewOrderRepository(db)
	usecase := usecases.NewCalculateFinalPrice(orderRepository)

	input := usecases.OrderInput{
		ID:    "123",
		Price: 10.0,
		Tax:   1.0,
	}

	output, err := usecase.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)

	// order, err := entities.NewOrder("1", -10, 1)
	// if err != nil {
	// 	println(err.Error())
	// } else {
	// 	println(order.ID)
	// }
	// car := Car{ // declarando e atribuindo a variável car
	// 	Model: "Ferrari",
	// 	Color: "Red",
	// }

	// //car.Model = "Fiat" // alterando o valor do atributo model
	// car.Start()
	// car.ChangeColor("Blue")
	// //println(car.Model)
	// println(car.Color) //
}
