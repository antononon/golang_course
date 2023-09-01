package main

import (
	"errors"
	"fmt"
)

func main() {
	// message, err := enterTheClub(7)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(message)

	fmt.Println(prediction("чкт"))
}

func enterTheClub(age int) (string, error) {
	if age >= 18 && age < 45 {
		return "Входи, только аккуратно", nil
	} else if age >= 45 && age < 65 {
		return "Вам точно понравится эта музыка?", nil
	} else if age >= 65 {
		return "Это уже слишком для вас", errors.New("you are too old")

	}
	return "Тебе нет 18-ти", errors.New("you are too yong")
}

func prediction(dayOfWeek string) (string, error) {
	switch dayOfWeek {
	case "пн":
		return "Хорошего тебе начала недели", nil

	case "вт":
		return "прекрасного тебе вторника", nil

	case "ср":
		return "Замечательной тебе среды", nil

	case "чт":
		return "Четверг - это маленькая пятница!", nil

	default:
		return "Не валидный день недели", errors.New("invalid day of the week")

	}
}
