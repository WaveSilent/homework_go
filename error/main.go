package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type User struct {
	Name     string
	Ballance int
}

func Pay(user *User, usd int) error {
	if user.Ballance-usd >= 0 {
		user.Ballance -= usd

		return nil
	}

	return errors.New("Недостаточно средств на балансе")
}

type Car struct {
	Armor int
}

func (c *Car) Gas() (error, int) {
	kmch := rand.Intn(100)

	if c.Armor-10 <= 0 {
		return errors.New("Машина не может разогнаться, так как броня закончилась"), 0
	} else {
		c.Armor -= 10
		return nil, kmch
	}
}

func main() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("Была паника:", p)
		}
	}()

	b := 0
	a := 1 / b

	fmt.Println(a)

	/*car := Car{
		Armor: 25,
	}

	pp.Println("Car до:", car)

	err, kmch := car.Gas()

	if err == nil {
		fmt.Println("Машина разогналась до:", kmch, "км/ч")
		pp.Println("Car после:", car)
	} else {
		fmt.Println("Что то произошло. Причина:", err.Error())
	}

	err, kmch = car.Gas()

	if err == nil {
		fmt.Println("Машина разогналась до:", kmch, "км/ч")
		pp.Println("Car после:", car)
	} else {
		fmt.Println("Что то произошло. Причина:", err.Error())
	}
	err, kmch = car.Gas()
	if err == nil {
		fmt.Println("Машина разогналась до:", kmch, "км/ч")
		pp.Println("Car после:", car)
	} else {
		fmt.Println("Что то произошло. Причина:", err.Error())
	}
	*/
}
