package version2

import (
	"fmt"
	"strings"
)

//ConsoleInput структура вводимых в консоль данных
type ConsoleInput struct {
	UserInput     int8
	UserDateInput int8
}

//Метод скана интового инпута
func (c ConsoleInput) scanInt() int8 {
	_, err := fmt.Scan(&c.UserInput)
	if ErrorCheck(err) {
		return 0
	}
	return c.UserInput
}

//Метод скана вводимой строки с датой
func (c ConsoleInput) scanDate() DateValues {
	var day, month, year int
	var monthString string
	switch UserDate.userDate {
	case 1:
		_, err := fmt.Scanf("\n%d.%d.%d", &day, &month, &year)
		if ErrorCheck(err) {
			return DateValues{}
		}
		return DateValues{day: day, month: month, year: year}
	case 2:
		_, err := fmt.Scanf("\n%d.%d.%d", &day, &month, &year)
		if ErrorCheck(err) {
			return DateValues{}
		}
		return DateValues{day: day, month: month, year: year}
	case 3:
		_, err := fmt.Scan(&day, &monthString, &year)
		if ErrorCheck(err) {
			return DateValues{}
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
		return DateValues{day: day, month: month, year: year}
	default:
		return DateValues{}
	}
}
