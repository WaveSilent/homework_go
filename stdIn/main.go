// Пользовательский ввод

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	infoIngredients := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Введите команду: ")

		if ok := scanner.Scan(); !ok {
			fmt.Println("Ошибка ввода")
			return
		}

		text := scanner.Text()

		fields := strings.Fields(text)

		if fields[0] == "Выйти" {
			break
		} else if fields[0] == "добавить" {
			fmt.Println("Добавлено", fields[2], fields[1])

			num, _ := strconv.Atoi(fields[2])

			infoIngredients[string(fields[1])] += num
			fmt.Println()
		} else if fields[0] == "удалить" {
			num, _ := strconv.Atoi(fields[2])

			if infoIngredients[string(fields[1])]-num < 0 {
				fmt.Println("Отрицательное кол-во ингридиентов")
			} else {
				fmt.Println("Удалено", fields[2], fields[1])

				infoIngredients[string(fields[1])] -= num
			}
			fmt.Println()
		} else if fields[0] == "получить" {
			fmt.Println("Итоговое кол-во индгридиентов", fields[1], ":", infoIngredients[string(fields[1])])
		} else if fields[0] == "help" {
			fmt.Println("Список доступных команд:")
			fmt.Println()

			fmt.Println("help -- эта команда должна выводить список доступных команд и их описание")
			fmt.Println("добавить {ингредиент} {количество} -- эта команда должна добавлять какой-то ингредиент, и запоминать сколько и чего мы добавили")
			fmt.Println("удалить {ингредиент} {количество} -- эта команда должна удалять какой-то ингредиент, если он был до этого добавлен")
			fmt.Println("получить {ингредиент} -- эта команда должна выводить на экран итоговое количество ингредиента")
			fmt.Println("выйти -- эта команда дожна завершать выполнение программы")
			fmt.Println()
		} else {
			fmt.Println("Неизвестная команда")
			fmt.Println()
		}

	}

}
