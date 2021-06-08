package version2

import "fmt"

//GetDate отвечает за получение дня рождения от юзера
func GetDate() DateValues {
	for {
		fmt.Printf("\nВведите дату рождения в формате %s:\n", UserDate.userDateString)
		return UserInput.scanDate()
	}
}

//ChooseDate отвечает за выбор юзером типа времени
func ChooseDate() {
	for {
		fmt.Println("Формат даты:\n\t1. DD.MM.YEAR\n\t2. MM.DD.YEAR\n\t3. DD Месяц YEAR")
		response := UserInput.scanInt()
		UserDate.userDate = response
		switch response {
		case 1:
			UserDate.userDateString = "DD.MM.YEAR"
		case 2:
			UserDate.userDateString = "MM.DD.YEAR"
		case 3:
			UserDate.userDateString = "DD Месяц YEAR"
		}
		return
	}
}
