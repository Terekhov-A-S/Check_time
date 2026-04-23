package main

import "fmt"

func main() {

	// Получаем входные данные и присваиваем их переменным
	time := getInput()
	checkTime(time)
}

// Функция возвращает полученные значения, не принимает параметров
func getInput() int64 {
	var getTime int64
	fmt.Print("Введите текущее время (без минут и секунд): ")
	fmt.Scanln(&getTime)
	return getTime
}

func checkTime(checkInput int64) {
	switch checkInput {
	case 6, 7, 8, 9, 10, 11:
		fmt.Println("На данный момент на улице утро.")
	case 12, 13, 14, 15, 16, 17:
		fmt.Println("На данный момент на улице день.")
	case 18, 19, 20, 21, 22:
		fmt.Println("На данный момент на улице вечер.")
	case 23, 24, 0, 1, 2, 3, 4, 5:
		fmt.Println("На данный момент на улице ночь.")
	default:
		fmt.Println("Вы ввели недопустимое время из параллельной вселенной")
	}
}
