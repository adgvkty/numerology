package version2

import (
	"fmt"
	"os"
	"time"
)

//Подгрузка ресурсов в карту
var resources map[int]DestinyText = loadResources()

//Переменная для хранения стандратной ошибки
var basicError string = "Вы ввели что-то не то"

//userInputInterface - интерфейс для вводимых юзером данных
type userInputInterface interface {
	scanInt() int8
	scanDate() dateValues
}

//magicCalculationsInterface - интерфейс описывающий магические калькуляции
type magicCalculationsInterface interface {
	calcMagicNumbers(dateValues) magicNumbers
	calcMagicStrings(magicNumbers, dateValues) magicStrings
}

//compatibilityCalculationsInterface - интерфейс описывающий магические калькуляции
type compatibilityCalculationsInterface interface {
	calcComp(magicStrings, magicStrings) int
}

//compCalculations структура данных через которую считается совместимость
var compCalculations compatibilityCalculationsInterface = Compitability{}

//magicCalculations структура данных через которую считают маг. числа
var magicCalculations magicCalculationsInterface = magicValues{}

//userInput сруктура данных через которую осуществляется ввод
var userInput userInputInterface = consoleInput{}

//dateFormat структура для хранения
//Выбранного юзером типа времени
type dateFormat struct {
	userDate       int8
	userDateString string
}

//dateValues хранит числовые значения дня рождения юзера
type dateValues struct {
	day, month, year int
}

//userDate является структурой для хранения
//типа выбраного времени юзером
var userDate dateFormat = dateFormat{}

//Core выполняет роль главной функции программы
func Core() {

	for {
		mainMenu()
		userDateValues := dateMenu()
		userMagicNumbers := magicCalculations.calcMagicNumbers(userDateValues)
		userMagicStrings := magicCalculations.calcMagicStrings(userMagicNumbers, userDateValues)
		destinyMenu(userMagicNumbers.destinyNum)
		time.Sleep(2 * time.Second)
		compatibilityMenu(userMagicStrings)
	}
}

//Меню совместимости - обработка ответа юзера и вызов расчетов
func compatibilityMenu(userMS magicStrings) {
	for {
		fmt.Println("\nЖелаете ли Вы узнать совместимость со своей второй половинкой?\n\t1. Да\n\t2. Нет")
		response := userInput.scanInt()
		if response == 2 {
			return
		}
		partnerDate := dateMenu()
		partnerMN := magicCalculations.calcMagicNumbers(partnerDate)
		partnerMS := magicCalculations.calcMagicStrings(partnerMN, partnerDate)
		comp := compCalculations.calcComp(userMS, partnerMS)
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
func dateMenu() dateValues {
	for {
		if userDate.userDate == 0 {
			chooseDate()
		}
		response := getDate()
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
		response := userInput.scanInt()
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
func errorCheck(err error) bool {
	if err != nil {
		fmt.Println(basicError)
		return true
	}
	return false
}
