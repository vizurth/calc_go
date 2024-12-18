package calc

import (
	"strings"
	"fmt"
	"strconv"
	"slices"
)

const alphas = "abcdefghijklmnopqrstuvwxyz"
const signs = "+-*/"
const numbers = "0123456789"
func checkOnAlpha(s string) bool{
	for _, r := range s{
		if strings.Contains(alphas, string(r)){
			return false
		}
	}
	return true
}

func checkOnSign(s string) bool{
	for _, r := range s{
		if strings.Contains(signs, string(r)){
			return true
		}
	}
	return false
}

func findSolution(stack []string) ([]string, error){
	counter := 0
	for len(stack)!=1{
		if strings.Contains("+-*/", stack[counter]){
			for index, elem := range stack{
				if elem == "*" || elem == "/" || elem == "+" || elem == "-"{
					counter = index
				}
			}
			a, erra := strconv.ParseFloat(stack[counter-1],64)
			b, errb := strconv.ParseFloat(stack[counter+1],64)
			if erra != nil || errb != nil{
				return []string{}, fmt.Errorf("err") 
			}
			switch stack[counter]{
			case "*":
				result := a * b

				temp := strconv.FormatFloat(result,'f', 10, 64)
				newStack := make([]string, 0)

				newStack = append(newStack, slices.Clone(stack[:counter-1])...)
				newStack = append(newStack, temp)
				newStack = append(newStack, slices.Clone(stack[counter+2:])...)
				stack = newStack
				counter = 0
			case "/":
				result := float64(0)
				if b != 0{
					result = a / b
				}
				temp := strconv.FormatFloat(result,'f', 10, 64)
				newStack := make([]string, 0)

				newStack = append(newStack, slices.Clone(stack[:counter-1])...)
				newStack = append(newStack, temp)
				newStack = append(newStack, slices.Clone(stack[counter+2:])...)
				stack = newStack
				counter = 0
			case "+":
				result := a + b

				temp := strconv.FormatFloat(result,'f', 10, 64)
				newStack := []string{temp}

				newStack = append(newStack, slices.Clone(stack[counter+2:])...)
				stack = newStack
				counter = 0
			case "-":
				result := a - b

				temp := strconv.FormatFloat(result,'f', 10, 64)
				newStack := []string{temp}

				newStack = append(newStack, slices.Clone(stack[counter+2:])...)
				stack = newStack
				counter = 0
			}
		} 
		counter++
	}
	return stack, nil
}

func Calc(expression string) (float64, error){
	if checkOnAlpha(expression) && checkOnSign(expression){
		expressionByte := []byte(expression)
		openingBrakets := 0
		closingBrakets := 0
		numStr := make([]byte, 0)
		stack := make([]string, 0)
		for i:=0; i<len(expressionByte); i++{
			//дополнительная проверка на четность скобок
			if string(expressionByte[i]) == "("{openingBrakets++}
			if string(expressionByte[i]) == ")"{closingBrakets++}

			if strings.Contains(numbers, string(expressionByte[i])){
				numStr = append(numStr, expressionByte[i])
			}

			if strings.Contains("+-/*", string(expressionByte[i])){
				if string(numStr) != ""{
					stack = append(stack, string(numStr))
				}
				stack = append(stack, string(expressionByte[i]))
				numStr = numStr[:0]
			}


			if string(expressionByte[i]) == "("{
				inBraket := make([]byte, 0)
				lenStr := 0
LOOP:		for _, elem := range expressionByte[i+1:]{
					inBraket = append(inBraket, elem)
					lenStr++
					if string(elem) == ")"{
						break LOOP
					}
				}
				i = i + lenStr
				resBraket, _ := Calc(string(inBraket))
				temp := strconv.FormatFloat(resBraket,'f', 10, 64)
				stack = append(stack, temp)
			}
		}
		// 10 + ))10
		if (openingBrakets > 0 && closingBrakets == 0) || (openingBrakets == 0 && closingBrakets > 0) || ((openingBrakets+closingBrakets)%2 != 0){
			return 0, fmt.Errorf("some error")
		}
		if string(numStr) != ""{
			stack = append(stack, string(numStr))
		}
		if strings.Contains("+-/*", stack[0]) || strings.Contains("+-/*", stack[len(stack)-1]){
			return 0, fmt.Errorf("some error")
		}
		if len(stack) < 3{
			return 0, fmt.Errorf("some error")
		}
		resSlice, err := findSolution(stack)
		if err != nil{
			return 0, fmt.Errorf("some error") 
		}
		result, _ := strconv.ParseFloat(resSlice[0],64)
		return result, nil
	} else {
		return 0, fmt.Errorf("some error")
	}
}