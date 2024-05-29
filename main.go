package main

import (
	"fmt"
	"strconv"
)

func main() {
	var expr1 string
	var expr2 string
	var expr3 string
	var err string

	fmt.Scanf("%s %s %s %s", &expr1, &expr2, &expr3, &err)

	if expr2 == "" || expr3 == "" {
		panic("Cтрока не является математической операцией")
	}
	if err != "" {
		panic("Формат математической операции не удовлетворяет заданию")
	}
	leng := search(expr1, expr3)
	calculation(leng, expr1, expr2, expr3)

}

func search(s1 string, s2 string) string {
	rim := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	arab := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	var val1 int
	var val2 int
	var val3 int
	var val4 int

	for _, value := range rim {
		if s1 == value {
			val1 = 1
		}
		if s2 == value {
			val2 = 1
		}
	}

	if val1 == 1 && val2 == 1 {
		return "rim"
	}

	for _, value := range arab {
		if s1 == value {
			val3 = 1
		}
		if s2 == value {
			val4 = 1
		}
	}

	if val3 == 1 && val4 == 1 {
		return "arab"
	}

	panic("Используются одновременно разные системы счисления.")
}

func calculation(leng string, expr1 string, expr2 string, expr3 string) {

	rim_arab := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
		"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10}

	arab_rim := map[int]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X"}

	var result int
	if leng == "rim" {
		switch expr2 {
		case "+":
			result = rim_arab[expr1] + rim_arab[expr3]
		case "-":
			result = rim_arab[expr1] - rim_arab[expr3]
		case "/":
			result = rim_arab[expr1] / rim_arab[expr3]
		case "*":
			result = rim_arab[expr1] * rim_arab[expr3]

		}
		if result > 10 {
			panic("В словаре нет римских чисел больше 10...нужно добавлять")
		}
		if result < 0 {
			panic("В римской системе нет отрицательных чисел.")
		} else {
			fmt.Println(arab_rim[result])
		}

	} else if leng == "arab" {
		e1, _ := strconv.Atoi(expr1)
		e3, _ := strconv.Atoi(expr3)
		switch expr2 {
		case "+":
			result = e1 + e3
		case "-":
			result = e1 - e3
		case "/":
			result = e1 / e3
		case "*":
			result = e1 * e3

		}
		fmt.Println(result)
	}

}
