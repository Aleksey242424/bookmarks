package main

import "fmt"

func main() {
	storage := map[string]string{}
	fmt.Println("1. Вывести все закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")
	for {
		var action int
		fmt.Scan(&action)
		if action == 1 {
			fmt.Println("Все закладки:")
			for key := range storage {
				fmt.Println(key)
			}
			fmt.Println("_____________________________________________________")
		} else if action == 2 {
			var key string
			var value string
			fmt.Println("Создание новой закладки")
			fmt.Print("Введите имя закладки: ")
			fmt.Scan(&key)
			fmt.Print("Введите значение закладки: ")
			fmt.Scan(&value)
			storage[key] = value
			fmt.Println("_____________________________________________________")
		} else if action == 3 {
			var key string
			fmt.Println("Удаление закладки")
			fmt.Print("Введите имя закладки которую хотите удалить: ")
			fmt.Scan(&key)
			delete(storage, key)
			fmt.Println("_____________________________________________________")
		} else if action == 4 {
			fmt.Println("Выход")
			break
		} else {
			fmt.Println("Незнакомое действие")
			fmt.Println("_____________________________________________________")
		}
	}
}
