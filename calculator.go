package main

import (
	"fmt"
	"strconv"
)

// Каулятор для счета арабских цифр
func calc_arab(a int, b string, c int) string {
	if a > 10 || c > 10 {
		return "Значение чисел больше 10"
	}

	if (a == 0 && b == "/") || (c == 0 && b == "/") {
		return "Делить на 0 нельзя"
	}

	var res int
	if b == "+" {
		res = a + c
	} else if b == "-" {
		res = a - c
	} else if b == "/" {
		res = a / c
	} else if b == "*" {
		res = a * c
	} else {
		return "Вывод ошибки, так как строка не является математической операцией."
	}
	return strconv.Itoa(res)
}

// Калькулятор для счета римских цифр
func calc_rim(a, b, c string) string {

	// карта с ключами - для перевода с римских в арабские цифры
	rim := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	// карта с ключами - для перевода с арабских в римские цифры
	rim_finish := map[int]string{
		1:   "I",
		2:   "II",
		3:   "III",
		4:   "IV",
		5:   "V",
		6:   "VI",
		7:   "VII",
		8:   "VIII",
		9:   "IX",
		10:  "X",
		11:  "XI",
		12:  "XII",
		13:  "XIII",
		14:  "XIV",
		15:  "XV",
		16:  "XVI",
		17:  "XVII",
		18:  "XVIII",
		19:  "XIX",
		20:  "XX",
		21:  "XXI",
		22:  "XXII",
		23:  "XXIII",
		24:  "XXIV",
		25:  "XXV",
		26:  "XXVI",
		27:  "XXVII",
		28:  "XXVIII",
		29:  "XXIX",
		30:  "XXX",
		31:  "XXXI",
		32:  "XXXII",
		33:  "XXXIII",
		34:  "XXXIV",
		35:  "XXXV",
		36:  "XXXVI",
		37:  "XXXVII",
		38:  "XXXVIII",
		39:  "XXXIX",
		40:  "XL",
		41:  "XLI",
		42:  "XLII",
		43:  "XLIII",
		44:  "XLIV",
		45:  "XLV",
		46:  "XLVI",
		47:  "XLVII",
		48:  "XLVIII",
		49:  "XLIX",
		50:  "L",
		51:  "LI",
		52:  "LII",
		53:  "LIII",
		54:  "LIV",
		55:  "LV",
		56:  "LVI",
		57:  "LVII",
		58:  "LVIII",
		59:  "LIX",
		60:  "LX",
		61:  "LXI",
		62:  "LXII",
		63:  "LXIII",
		64:  "LXIV",
		65:  "LXV",
		66:  "LXVI",
		67:  "LXVII",
		68:  "LXVIII",
		69:  "LXIX",
		70:  "LXX",
		71:  "LXXI",
		72:  "LXXII",
		73:  "LXXIII",
		74:  "LXXIV",
		75:  "LXXV",
		76:  "LXXVI",
		77:  "LXXVII",
		78:  "LXXVIII",
		79:  "LXXIX",
		80:  "LXXX",
		81:  "LXXXI",
		82:  "LXXXII",
		83:  "LXXXIII",
		84:  "LXXXIV",
		85:  "LXXXV",
		86:  "LXXXVI",
		87:  "LXXXVII",
		88:  "LXXXVIII",
		89:  "LXXXIX",
		90:  "XC",
		91:  "XCI",
		92:  "XCII",
		93:  "XCIII",
		94:  "XCIV",
		95:  "XCV",
		96:  "XCVI",
		97:  "XCVII",
		98:  "XCVIII",
		99:  "XCIX",
		100: "C",
	}

	fmt.Println(rim[a])

	x1 := rim[a]
	x2 := rim[c]

	if x1 == 0 || x2 == 0 {
		return "Вывод ошибки, один из параметров не является римской цифрой."
	}

	if b == "-" && x1-x2 < 0 {
		return "Вывод ошибки, так как в римской системе нет отрицательных чисел."
	}

	switch {
	case b == "+":
		return rim_finish[x1+x2]
	case b == "-":
		return rim_finish[x1-x2]
	case b == "/":
		return rim_finish[x1/x2]
	case b == "*":
		return rim_finish[x1*x2]
	default:
		return "Вывод ошибки, так как строка не является математической операцией."
	}
}

func main() {
	fmt.Println("Введите данные:")

	var a, b, c string

	fmt.Scanln(&a, &b, &c)

	// strconv.Atoi - функция проверяет является ли строка числовым значением.
	// возвращает 2 значения: 1 - int, 2 - err (ошибка если она есть)
	valA, errA := strconv.Atoi(a)
	valC, errC := strconv.Atoi(c)

	// проверяем a и с числами. если являются вызываем функцию calc_arab
	// если a или с не являются числом, выводим ошибку
	// иначе вызываем функцию calc_rim
	if errA == nil && errC == nil {
		fmt.Println(calc_arab(valA, b, valC))
	} else if errA == nil && errC != nil || errA != nil && errC == nil {
		fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
	} else {
		fmt.Println(calc_rim(a, b, c))
	}
}
