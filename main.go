package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func digitMath (a string, b uint8, c int) string {
	result := ""
	if b == 42 {
		for i := 1; i <= c; i++ {
			result += a
		}
	}
	if b == 47 {
		result = a[:(len(a)/c)]
	}
	return result
}

func stringMath (a string, b uint8, c string) string {
	result := ""
	if b == 43 {
		result = a + c
	}
	if b == 45 {
		if len(c) > len(a) {
			panic("Некорректная операция")
		} else {
			result = strings.ReplaceAll(a, c, "")
		}
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите значение в формате \"Строка\" + \"Строка\" или \"Строка\" * число")
		text, _ := reader.ReadString('\n')
		input := text[:len(text)-1] + strings.Repeat(" ", 30)
		result := ""
		// Будем работать с числами в юникоде
		// Проверка на первую кавычку
		if input[0] != 34 {
			panic("Первый аргумент должен быть строкой и выделяться \"")
		} else {
			firstFlag := false
			idx := 0
			// Проверка, что слово заканчивается на кавычку и после кавычки есть пробел, а так же длина строки до 10 символов
			for i := 1; i < 12; i++ {
				if (input[i] == 34) && (input[i+1] == 32) {
					firstFlag = true
					idx = i
					break
				}
			}
			if !firstFlag {
				panic("Первое слово длиннее 10 символов или не закрыто кавычками")
			}
			firstWord := input[1:idx]
			operator := input[idx+2]
			// Проверяем, что введем корректный мат. оператор и отделен пробелами
			if (input[idx+1] == 32) && (input[idx+3] == 32) {
				// Если ввели + или -
				idxNew := idx+4
				if (operator == 43) || (operator == 45) {
					if input[idxNew] != 34 {
						panic("При использовании + или - последним аргументом должна быть строка и выделяться \"")
					} else {
						lastFlag := false
						idxLast := idxNew
						for i := 1; i < 12; i++ {
							if (input[idxNew+i] == 34) && (input[idxNew+i+1] == 32) {
								lastFlag = true
								idxLast += i
								break
							}
						}
						if !lastFlag {
							panic("Последнее слово длиннее 10 символов или не закрыто кавычками")
						}
						lastWord := input[idxNew+1:idxLast]
						result = stringMath (firstWord, operator, lastWord)
					}
				} else if (operator == 42) || (operator == 47) { //Если ввели * или /
					accept := [10]uint8{49,50,51,52,53,54,55,56,57}
					convertedNumber := 11
					digit := input[idxNew:idxNew+3]
					digitFlag := false
					for _,v := range accept {
						if v == digit[0] {
							digitFlag = true
						}
					}
					if digitFlag {
						if (digit[0] != 49) && (digit[1] == 32) {
							numberStr := string(digit[0]) // преобразуем uint8 в строку
							convertedNumber, _ = strconv.Atoi(numberStr)
						} else if (digit[0] == 49) && (digit[1] == 32) {
							convertedNumber = 1
						} else if (digit[0] == 49) && (digit[1] == 48) && (digit[2] == 32) {
							convertedNumber = 10
						} else { 
							panic("Второй аргумент должен быть целым числом от 1 до 10")
						}
						result = digitMath (firstWord, operator, convertedNumber)
					} else {
						panic("Второй аргумент должен быть целым числом от 1 до 10")
					}
				} else {
				 	panic("Некорректный мат. оператор")
				}
			} else {
				panic("аргументы должны быть разделены пробелами")
			}
		}
		if len(result) < 41 {
			fmt.Printf("\"%s\"\n", result)
		} else {
			fmt.Printf("\"%s...\"\n", result[:40])
		}
	}	
}	
