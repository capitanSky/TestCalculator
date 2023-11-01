package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin) // создается буферизованный читатель Reader для потока ввода os.Stdin
	fmt.Println("Добро пожаловать в калькулятор!!")
	fmt.Println("Введите выражение:")

	input, _ := reader.ReadString('\n') //чтение строки с помощью ReadString до символа новой строки '\n'
	input = strings.TrimSpace(input)    //удаление всех пробелов и символов перевода строки из строки с перезаписью исходной переменной input
	parts := strings.Split(input, " ")  //разбитие строки на подстроки с помошью разделителя " " и сохранение результа в срезе parts

	if len(parts) != 3 { // проверка, если количество элементов в серезе не ровно 3 вылетает ошибка
		fmt.Println("Строка не является математической операцией либо арифметическая операция не удовлетворяет заданию")
		return
	}

	//множественное присвание значения элементов среза переменным, в соответствии с их порядком, переменной а присваивается
	//значение элемента среза parts[0] и так длаее
	a, operator, b := parts[0], parts[1], parts[2]
	//проверяется валидность переменных (a, operator, b) с помощью функций isValidInput, isValidOperator. если при проверке
	// результат хоть одной переменной будет true то результат общего выражения будет true и будет выполняться команды
	// в фигурных скобках
	if !isValidInput(a) || !isValidOperator(operator) || !isValidInput(b) {
		fmt.Println("Неподдерживаемый ввод")
		return
	}
	//вызывается функция calculate, в параметах указываюся переменные (a, operator, b), результат вычисления
	//сохраняется в переменную result
	result, err := calculate(a, operator, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	// результа вычисления сохраненный в переменной result выводиться в консоль
	fmt.Println("Результат: ", result)

}

func isValidInput(input string) bool {
	if isRomanNumber(input) { //функция используется для проверки является текст римским числом
		return true
	}

	number, err := strconv.Atoi(input) //преобразование введеного текста в число с помощью strconv.Atoi
	// проверка, если err равно nil (отсутствие ошибок) и переменная number указана в диапозоне от 1 до 10
	// то все выражение возвращает true и означает что введенное значение является допустимым вводом
	if err == nil && number >= 1 && number <= 10 {
		return true
	}
	return false
}

func isRomanNumber(input string) bool {
	//регулярное выражение romanPattern - шаблон для распознавания римских чисел от I до X, регулярное выражение
	// используется для сопоставления и поиска текстовых шаблонов
	romanPattern := `^(I|II|III|IV|V|VI|VII|VIII|IX|X|XL|L|XC|C|CD|D|CM|M)$`
	//regexp.MatchString принимает регулярное выражение romanPattern и входную строку input, возвращает true если
	//срока соответствует регулярному выражению, результат записывается в переменную match
	match, _ := regexp.MatchString(romanPattern, input)
	return match

}
func isValidOperator(operator string) bool {
	//срез содержащий допустимые операторы
	operators := []string{"+", "-", "*", "/"}
	// перебор допустимых операторов с помощью опеартора for и оператора range
	for _, validOperator := range operators {
		//проверка внутри цикла, если введенный operator совпадает с допустимым оператором возвращается true
		// что означает что оператор допустим в противном случае false
		if operator == validOperator {
			return true
		}
	}
	return false
}

func calculate(a, operator, b string) (string, error) {

	var numA, numB int
	var isRoman bool
	//проверка являются ли a и b римскими или арабскими, если оба числа римских они преобразуются в арабские
	if isRomanNumber(a) && isRomanNumber(b) {
		numA = romanToArabic(a)
		numB = romanToArabic(b)
		isRoman = true
	//преобразование а и b в целое число с помощью функции strconv.Atoi
	} else {
		var err error
		numA, err = strconv.Atoi(a)
		if err != nil {
			return "", errors.New("Используются одновременно разные системы счисления")
		}
		numB, err = strconv.Atoi(b)
		if err != nil {
			return "", errors.New("Используются одновременно разные системы счисления")
		}
	}
	var result int
	switch operator {
	case "+":
		result = numA + numB
	case "-":
		result = numA - numB
	case "*":
		result = numA * numB
	case "/":
		if numB == 0 {
			return "", errors.New("Деление на ноль")
		}
		result = numA / numB
	default:
		return "", errors.New("Недопустимый оператор")
	}
	if isRoman && result < 1 {
		return "", errors.New("В римской системе нет отрицательных чисел")

	}
	if isRoman {
		return arabicToRoman(result), nil
	}
	return strconv.Itoa(result), nil

	//return arabicToRoman(result), nil
}



var romanMap = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

// Функция для преобразования арабских чисел в римские
func arabicToRoman(num int) string {
	// совершается проход по всем элементам среза keys, при соотвествии условию
	//num >= key запускается внутрений цикл, где в переменную roman записывается
	//элемент из карты romanMap в соответствии с его ключом 
	roman := ""
	keys := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	for _, key := range keys {
		for num >= key {
			roman += romanMap[key]
			num -= key
		}
	}
	return roman
}

// функция преобразования римской цифры в арабские
func romanToArabic(roman string) int {

	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	previousValue := 0


	// цикл проходит в обратном порядке по элементам сроки roman
	for i := len(roman) - 1; i >= 0; i-- {
		// переменной currentValue присваивается значение из карты romanMap.
		// currentValue представляет числовое значние риамкой цифры.
		currentValue := romanMap[rune(roman[i])]
		// если currentValue меньше previousValue, currentValue вычитается из result
		if currentValue < previousValue {
			result -= currentValue
		//	currentValue прибавляется result
		} else {
			result += currentValue
		}
		// текущее значение сохраняется в предыдущее для последующего сравнения с новой римской
		// цифрой при последующей итерации
		previousValue = currentValue
	}

	return result
}
