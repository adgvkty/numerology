package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

//Карта, которая хранит в себе ресурсы
//LoadResources() вызывается из другого файла
var resources map[int]DestinyText = LoadResources()

//Текст базовой ошибки
var basicError string = "Вы ввели что-то не то"

//Структура для хранения числовых значений даты дня рождения
type dateValues struct {
	day, month, year int
}

//Метод, который возвращает заполненный тип magicValues
//последовательно заполняя встроенные
func (date dateValues) calcMagic() magicValues {
	magicNumbers := date.calcMagicNumbers()
	magicStrings := magicStrings{birthdayString: date.dvTOstring(),
		magicNumString: magicNumbers.mnTOstring()}
	return magicValues{magicNumbers: magicNumbers, magicStrings: magicStrings}
}

//Метод для перевода дня рождения в строку
func (date dateValues) dvTOstring() string {
	return conc([]string{toStr(date.day),
		toStr(date.month),
		toStr(date.year)})
}

//Метод для подсчета магических чисел
func (date dateValues) calcMagicNumbers() magicNumbers {
	resMN := magicNumbers{}
	resMN.nOne = digitSum(date.day) + digitSum(date.month) + digitSum(date.year)
	resMN.nTwo = digitSum(resMN.nOne)
	resMN.nThree = resMN.nOne - date.day/10*2
	resMN.nFour = digitSum(resMN.nThree)
	resMN.destinyNum = digitSum(digitSum(date.day) + digitSum(date.month) + digitSum(date.year))
	for resMN.destinyNum >= 10 {
		resMN.destinyNum = digitSum(resMN.destinyNum)
	}
	return resMN
}

//Составной тип для хранения маг. чисел и строк
type magicValues struct {
	magicNumbers
	magicStrings
}

//Тип для магических чисел и числа судьбы
type magicNumbers struct {
	nOne, nTwo, nThree, nFour, destinyNum int
}

//Метод, который перевод магические числа в строку
func (mN magicNumbers) mnTOstring() string {
	return conc([]string{toStr(mN.nOne),
		toStr(mN.nTwo),
		toStr(mN.nThree),
		toStr(mN.nFour)})
}

//Тип для магических строк
type magicStrings struct {
	birthdayString string
	magicNumString string
}

//Переменная для хранения значения типа времени
var dateFormat int8

//Переменная для хранения строки с типом времени
var dateFormatString string

//Core - главная функция, отвечает за вызов всех меню
//С соответствующими аргументами
func main() {
	for {
		mainMenu()
		userDate := dateMenu()
		userMagicValues := userDate.calcMagic()
		destinyMenu(userMagicValues.destinyNum)
		time.Sleep(2 * time.Second)
		compatibilityMenu(userMagicValues)
	}
}

//Меню проверки совместимости
func compatibilityMenu(uMV magicValues) {
	res := askComp()
	if res == 2 {
		return
	}
	partnerDate := dateMenu()
	pMV := partnerDate.calcMagic()
	compVal := compatibilityCalc(uMV.magicNumString,
		uMV.birthdayString, pMV.magicNumString, pMV.birthdayString)
	fmt.Printf("\nВаша совместимость - %d. ", compVal)
	if compVal >= 8 {
		fmt.Println("Вы превосходно подходите друг-другу!!")
	} else if compVal >= 4 && compVal <= 7 {
		fmt.Println("Вы неплохо подходите друг-другу!")
	} else {
		fmt.Println("Ваша совместимость не очень..")
	}

}

//Функция подсчета числового значения совместимости
//на основе сравнения магических строк юзера и его партнера
func compatibilityCalc(uMS, uBS, pMS, pBS string) int {
	var comp int

	for i := 1; i <= 9; i++ {
		iStr := toStr(i)
		iNumberUser := strings.Count(uMS, iStr) + strings.Count(uBS, iStr)
		iNumberPartner := strings.Count(pMS, iStr) + strings.Count(pBS, iStr)
		if iNumberUser > iNumberPartner {
			comp += iNumberPartner
		} else {
			comp += iNumberUser
		}

	}
	return comp
}

//Обработка ответа юзера на вопрос о совместимости
func askComp() int8 {
	var response int8
	for {
		fmt.Println("\nЖелаете ли Вы узнать совместимость со своей второй половинкой?\n\t1. Да\n\t2. Нет")
		_, err := fmt.Scan(&response)
		if err != nil {
			fmt.Println(basicError)
			continue
		}
		return response
	}
}

//Меню вывода текстов судьбы
func destinyMenu(d int) {
	fmt.Println("Наши сионские мудрецы считают Ваше магические число..")
	time.Sleep(2 * time.Second)
	fmt.Printf("\nЭто число равняется - %v\nСейчас они расскажут Вам всё, что они знают..\n", d)
	time.Sleep(2 * time.Second)
	destiny := resources[d]
	fmt.Printf("%s %s %s %s",
		destiny.destinyChar,
		destiny.destinyPros,
		destiny.destinyCons,
		destiny.destinyPurpose,
	)
}

//Меню выбора типа даты и обработки ввода дня рождения
func dateMenu() dateValues {
	for {
		if dateFormat == 0 {
			chooseDate()
		}
		return getDate()
	}
}

//Обработка введенной даты дня рождения
func getDate() dateValues {
	var day, month, year int
	var monthString string

	for {
		fmt.Printf("\nВведите дату рождения в формате %s:\n", dateFormatString)
		switch dateFormat {
		case 1:
			_, err := fmt.Scanf("\n%d.%d.%d", &day, &month, &year)
			if err != nil {
				fmt.Println(basicError)
				continue
			}
			return dateValues{day: day, month: month, year: year}
		case 2:
			_, err := fmt.Scanf("\n%d.%d.%d", &day, &month, &year)
			if err != nil {
				fmt.Println(basicError)
				continue
			}
			return dateValues{day: day, month: month, year: year}
		case 3:
			_, err := fmt.Scan(&day, &monthString, &year)
			if err != nil {
				fmt.Println(basicError)
				continue
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
		}
	}
}

//Выбор юзером типа даты
func chooseDate() {
	var response int8

	fmt.Println("Формат даты:\n\t1. DD.MM.YEAR\n\t2. MM.DD.YEAR\n\t3. DD Месяц YEAR")
	for {
		_, err := fmt.Scan(&response)
		if err != nil {
			fmt.Println(basicError)
			continue
		}
		dateFormat = response
		switch response {
		case 1:
			dateFormatString = "DD.MM.YEAR"
		case 2:
			dateFormatString = "MM.DD.YEAR"
		case 3:
			dateFormatString = "DD Месяц YEAR"
		}
		return
	}
}

//Главное меню и обработка ввода юзера в нём
func mainMenu() {
	var response int8
	for {
		fmt.Println("\nДобро пожаловать в Нумеролого-о-метр!\n\t1. Начать работу\n\t2. Справка\n\t3. Завершить работу")
		_, err := fmt.Scan(&response)
		if err != nil {
			fmt.Println(basicError)
			continue
		}
		switch response {
		case 1:
			return
		case 2:
			readme()
		case 3:
			os.Exit(0)
		}
	}
}

func readme() {
	var response int8
	for {
		fmt.Println("Добро пожаловать в справку!\n\t1. Как пользоваться\n\t2. О программе\n\t3. Выйти в главное меню")
		_, err := fmt.Scan(&response)
		if err != nil {
			fmt.Println(basicError)
			continue
		}
		switch response {
		case 1:
			fmt.Println("Выбирайте пункты в меню с помощью ввода цифр,\nно не расстраивайтесь если Вы ввели что-либо\nнеправильно, программа обязательно попросит\nВас попробовать ещё раз.\nНа некоторые операции (подсчёт магических чисел\nили высчитывание судьбы) может потребоваться\nнесколько секунд - это нормально.")
		case 2:
			fmt.Println("Данная программа была спонсирована госдепом Украины.\nЕё запуск - согласие на обработку личных данных.")
		case 3:
			return
		}

	}
}

//Конкатенации строк из массива в одну
func conc(array []string) string {
	var res string

	for _, v := range array {
		res += v
	}
	return res
}

//Конвертирует int в string
func toStr(num int) string {
	return strconv.Itoa(num)
}

//Подсчет суммы цифр числа
//Результат может быть >= 10
func digitSum(digit int) int {
	var result = 0
	for ; digit > 0; digit /= 10 {
		result += digit % 10
	}
	return result
}
