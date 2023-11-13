package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}
	result, err := evaluateExpression(input)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		resultWithoutQuotes := strings.ReplaceAll(result, `"`, "")
		fmt.Println("Результат:", resultWithoutQuotes)

	}
}

func evaluateExpression(input string) (string, error) {

	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")

	var operand1 string
	var operator string
	var operand2 string

	if len(parts) == 4 {
		operand1 = parts[0] + " " + parts[1]
		operator = parts[2]
		operand2 = parts[3]
	} else if len(parts) == 3 {
		operand1 = parts[0]
		operator = parts[1]
		operand2 = parts[2]
	}

	switch operator {
	case "+":
		return addStrings(operand1, operand2), nil
	case "-":
		return subtractStrings(operand1, operand2), nil
	case "*":
		num, err := strconv.Atoi(operand2)
		if err != nil || num < 1 || num > 10 {
			return "", fmt.Errorf("Неподдерживаемое значение числа")
		}
		return multiplyStrings(operand1, num), nil
	case "/":
		num, err := strconv.Atoi(operand2)
		if err != nil || num < 1 || num > 10 {
			return "", fmt.Errorf("Неподдерживаемое значение числа")
		}
		return divideStrings(operand1, num), nil

	default:
		return "", fmt.Errorf("Указан неподдерживаемый оператор - ", operator)
	}
}

func addStrings(operand1, operand2 string) string {

	return operand1 + operand2
}
func subtractStrings(operand1, operand2 string) string {

	return strings.Replace(operand1, operand2, "", -1)
}

func multiplyStrings(operand1 string, num int) string {

	result := ""

	for i := 0; i < num; i++ {
		result += operand1
	}
	return result
}

func divideStrings(operand1 string, num int) string {
	length := len(operand1) / num
	if length > 40 {
		return operand1[:40] + "..."
	}
	return operand1[:length]

}
