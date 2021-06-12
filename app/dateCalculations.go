package app

import "fmt"

//GetDate отвечает за получение дня рождения от юзера
func getDate() dateValues {
	for {
		fmt.Printf("\nВведите дату рождения в формате %s:\n", userDate.userDateString)
		return userInput.scanDate()
	}
}

//ChooseDate отвечает за выбор юзером типа времени
func chooseDate() {
	for {
		fmt.Println("Формат даты:\n\t1. DD.MM.YEAR\n\t2. MM.DD.YEAR\n\t3. DD Месяц YEAR")
		response := userInput.scanInt()
		userDate.userDate = response
		switch response {
		case 1:
			userDate.userDateString = "DD.MM.YEAR"
			return
		case 2:
			userDate.userDateString = "MM.DD.YEAR"
			return
		case 3:
			userDate.userDateString = "DD Месяц YEAR"
			return
		default:
			fmt.Println(basicError)
			userDate.userDateString = ""
		}
	}
}
