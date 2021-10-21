package app

import "strings"

// Compitability структура для хранения данных о совместимости
type Compitability struct {
	Compitability int
}

// Метод подсчета совместимости по маг. строкам
func (c Compitability) calcComp(uMS magicStrings, pMS magicStrings) int {
	var comp int

	for i := 1; i <= 9; i++ {
		iStr := toStr(i)
		iNumberUser := strings.Count(uMS.magicString, iStr) + strings.Count(uMS.birthdayString, iStr)
		iNumberPartner := strings.Count(pMS.magicString, iStr) + strings.Count(pMS.birthdayString, iStr)
		if iNumberUser > iNumberPartner {
			comp += iNumberPartner
		} else {
			comp += iNumberUser
		}

	}
	return comp
}
