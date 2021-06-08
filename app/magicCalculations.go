package app

import (
	"strconv"
	"strings"
)

//Тип для хранения магических чисел
type magicNumbers struct {
	One, Two, Three, Four, destinyNum int
}

//Метод подсчета магических чисел
func (m magicNumbers) calcMagicNumbers(date dateValues) magicNumbers {
	m.One = digitSum(date.day) + digitSum(date.month) + digitSum(date.year)
	m.Two = digitSum(m.One)
	m.Three = m.One - date.day/10*2
	m.Four = digitSum(m.Three)

	m.destinyNum = digitSum(digitSum(date.day) + digitSum(date.month) + digitSum(date.year))
	for m.destinyNum >= 10 {
		m.destinyNum = digitSum(m.destinyNum)
	}
	return m
}

//Тип для хранения магических строк
type magicStrings struct {
	birthdayString string
	magicString    string
}

//Метод для подсчета магических строк
func (m magicStrings) calcMagicStrings(magicN magicNumbers, date dateValues) magicStrings {
	dateString := []string{toStr(date.day), toStr(date.month), toStr(date.year)}
	m.birthdayString = strings.Join(dateString, "")
	numbersString := []string{toStr(magicN.One), toStr(magicN.Two), toStr(magicN.Three), toStr(magicN.Four)}
	m.magicString = strings.Join(numbersString, "")
	return m
}

//magicValues структура для хранения всех магических значений
type magicValues struct {
	magicNumbers
	magicStrings
}

//Функция подсчета суммы цифр
func digitSum(digit int) int {
	var result = 0
	for ; digit > 0; digit /= 10 {
		result += digit % 10
	}
	return result
}

//Быстрый перевод в интов в строки
func toStr(num int) string {
	return strconv.Itoa(num)
}
