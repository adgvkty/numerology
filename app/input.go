package version2

import (
	"fmt"
	"strings"
)

//consoleInput структура вводимых в консоль данных
type consoleInput struct {
	UserInput int8
}

//Метод скана интового инпута
func (c consoleInput) scanInt() int8 {
	_, err := fmt.Scan(&c.UserInput)
	if errorCheck(err) {
		return 0
	}
	return c.UserInput
}

//Метод скана вводимой строки с датой
func (c consoleInput) scanDate() dateValues {

	var day, month, year int
	var monthString string

	switch userDate.userDate {
	case 1:
		_, err := fmt.Scanf("\n%d.%d.%d", &day, &month, &year)
		if errorCheck(err) {
			return dateValues{}
		}
		return dateValues{day: day, month: month, year: year}
	case 2:
		_, err := fmt.Scanf("\n%d.%d.%d", &day, &month, &year)
		if errorCheck(err) {
			return dateValues{}
		}
		return dateValues{day: day, month: month, year: year}
	case 3:
		_, err := fmt.Scan(&day, &monthString, &year)
		if errorCheck(err) {
			return dateValues{}
		}
		monthString = strings.ToLower(monthString)
		switch monthString {
		case "январь", "января":
			month = 1
		case "февраль", "февраля":
			month = 2
		case "март", "марта":
			month = 3
		case "апрель", "апреля":
			month = 4
		case "май", "мая":
			month = 5
		case "июнь", "июня":
			month = 6
		case "июль", "июля":
			month = 7
		case "август", "августа":
			month = 8
		case "сентябрь", "сентября":
			month = 9
		case "октябрь", "октября":
			month = 10
		case "ноябрь", "ноября":
			month = 11
		case "декабрь", "декабря":
			month = 12
		}
		return dateValues{day: day, month: month, year: year}
	default:
		return dateValues{}
	}
}
