package main

import (
	"fmt"
)

type storageMap = map[string]string

func main() {
	storage := storageMap{}
	printInfo()
Menu:
	for {
		var action int
		fmt.Scan(&action)
		switch action {
		case 1:
			printAllBookmarks(storage)
		case 2:
			storage = addNewBookmarks(storage)
		case 3:
			storage = deleteBookmarks(storage)
		case 4:
			fmt.Println("Выход")
			break Menu
		default:
			fmt.Println("Незнакомое действие")
			fmt.Println("_____________________________________________________")
		}
	}
}

func printInfo() {
	fmt.Println("1. Вывести все закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")
}

func printAllBookmarks(storage storageMap) {
	fmt.Println("Все закладки:")
	for key, value := range storage {
		fmt.Println(key, value)
	}
	fmt.Println("_____________________________________________________")
}

func addNewBookmarks(storage storageMap) storageMap {
	var key string
	var value string
	fmt.Println("Создание новой закладки")
	fmt.Print("Введите имя закладки: ")
	fmt.Scan(&key)
	fmt.Print("Введите значение закладки: ")
	fmt.Scan(&value)
	storage[key] = value
	fmt.Println("_____________________________________________________")
	return storage
}

func deleteBookmarks(storage storageMap) storageMap {
	var key string
	fmt.Println("Удаление закладки")
	fmt.Print("Введите имя закладки которую хотите удалить: ")
	fmt.Scan(&key)
	delete(storage, key)
	fmt.Println("_____________________________________________________")
	return storage
}
