package version2

import (
	"strconv"
	"strings"
)

//Тип для хранения магических чисел
type magicNumbers struct {
	nOne, nTwo, nThree, nFour, destinyNum int
}

//Метод подсчета магических чисел
func (m magicNumbers) calcMagicNumbers(date DateValues) magicNumbers {
	m.nOne = digitSum(date.day) + digitSum(date.month) + digitSum(date.year)
	m.nTwo = digitSum(m.nOne)
	m.nThree = m.nOne - date.day/10*2
	m.nFour = digitSum(m.nThree)
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
func (m magicStrings) calcMagicStrings(mN magicNumbers, date DateValues) magicStrings {
	str := []string{toStr(date.day), toStr(date.month), toStr(date.year)}
	m.birthdayString = strings.Join(str, "")
	str = []string{toStr(mN.nOne), toStr(mN.nTwo), toStr(mN.nThree), toStr(mN.nFour)}
	m.magicString = strings.Join(str, "")
	return m
}

//MagicValues структура для хранения всех магических значений
type MagicValues struct {
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
