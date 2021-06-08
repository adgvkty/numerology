package version2

import (
	"fmt"
	"os"
	"time"
)

//Подгрузка ресурсов в карту
var resources map[int]DestinyText = LoadResources()

//Переменная для хранения стандратной ошибки
var basicError string = "Вы ввели что-то не то"

//UserInputInterface - интерфейс для вводимых юзером данных
type UserInputInterface interface {
	scanInt() int8
	scanDate() DateValues
}

//MagicCalculationsInterface - интерфейс описывающий магические калькуляции
type MagicCalculationsInterface interface {
	calcMagicNumbers(DateValues) magicNumbers
	calcMagicStrings(magicNumbers, DateValues) magicStrings
}

//CompatibilityCalculationsInterface - интерфейс описывающий магические калькуляции
type CompatibilityCalculationsInterface interface {
	calcComp(magicStrings, magicStrings) int
}

//CompCalcilations структура данных через которую считается совместимость
var CompCalcilations CompatibilityCalculationsInterface = Compitability{}

//MagicCalculations структура данных через которую считают маг. числа
var MagicCalculations MagicCalculationsInterface = MagicValues{}

//UserInput сруктура данных через которую осуществляется ввод
var UserInput UserInputInterface = ConsoleInput{}

//DateFormat структура для хранения
//Выбранного юзером типа времени
type DateFormat struct {
	userDate       int8
	userDateString string
}

//DateValues хранит числовые значения дня рождения юзера
type DateValues struct {
	day, month, year int
}

//UserDate является структурой для хранения
//типа выбраного времени юзером
var UserDate DateFormat = DateFormat{}

//Core выполняет роль главной функции программы
func Core() {

	for {
		mainMenu()
		userDateValues := dateMenu()
		userMagicNumbers := MagicCalculations.calcMagicNumbers(userDateValues)
		userMagicStrings := MagicCalculations.calcMagicStrings(userMagicNumbers, userDateValues)
		destinyMenu(userMagicNumbers.destinyNum)
		time.Sleep(2 * time.Second)
		compatibilityMenu(userMagicStrings)
	}
}

//Меню совместимости - обработка ответа юзера и вызов расчетов
func compatibilityMenu(uMS magicStrings) {
	for {
		fmt.Println("\nЖелаете ли Вы узнать совместимость со своей второй половинкой?\n\t1. Да\n\t2. Нет")
		response := UserInput.scanInt()
		if response == 2 {
			return
		}
		pD := dateMenu()
		pMN := MagicCalculations.calcMagicNumbers(pD)
		pMS := MagicCalculations.calcMagicStrings(pMN, pD)
		comp := CompCalcilations.calcComp(uMS, pMS)
		if comp >= 8 {
			fmt.Println("Вы превосходно подходите друг-другу!!")
		} else if comp >= 4 && comp <= 7 {
			fmt.Println("Вы неплохо подходите друг-другу!")
		} else {
			fmt.Println("Ваша совместимость не очень..")
		}
		return
	}
}

//Меню судьбы - вывод текстов судьбы
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

//Меню даты - выбор юзером даты и ввод дня рождения
func dateMenu() DateValues {
	for {
		if UserDate.userDate == 0 {
			ChooseDate()
		}
		response := GetDate()
		if response.day == 0 && response.month == 0 && response.year == 0 {
			continue
		}
		return response
	}
}

//Главное меню - обработка выборов юзера
func mainMenu() {
	for {
		fmt.Println("\nДобро пожаловать в Нумеролого-о-Метр!\n\t1. Начать работу\n\t2. Завершить работу")
		response := UserInput.scanInt()
		switch response {
		case 0:
			continue
		case 1:
			return
		case 2:
			os.Exit(0)
		}
	}
}

//ErrorCheck глобальная функция для проверки чего либо на ошибку
func ErrorCheck(err error) bool {
	if err != nil {
		fmt.Println(basicError)
		return true
	}
	return false
}
