package main

//Если начать вводить римские символы на "Ру" языке, а потом исправить на "Англ",
//то выводится ошибка "Неизвестный символ" Хотя по итогу символ правильный.
//Так же если ввести "Ру" символ, а потом вписать арабскую цифру, выйдет такая же ошибка(Не знаю как исправить)
//Арабскими цифрами можно вводить отрицательные значения и работать с ними (Могу исправить если надо)
//Вводить цифры и математический символ можно через пробел ИЛИ с новой строки через "Энтер"
//Можно ввести 1 + и хоть обспамиться энтером, паники не будет,
//вам всё ещё нужно будет ввести второе число для продолжения (Пока хз как сделать по другому, но если надо - сделаю)
import (
	"fmt"
	"strconv"
)

var a, y, b string

func main() {
	fmt.Println("Введите задачу")
	fmt.Scan(&a, &y, &b)
	aNum, aRom := strconv.Atoi(a)
	bNum, bRom := strconv.Atoi(b)
	if aNum > 10 || bNum > 10 {
		panic("AlARM!!! Вводимые цифры не могут быть больше 10!")
	} else if aNum < -10 || bNum < -10 {
		panic("Вводимое число меньше -10")
	} else if aRom != nil && bRom != nil {
		numRome(a, b, y)
	} else if aRom != nil || bRom != nil {
		panic("AlARM!!! Вы пытаетесь сложить арабские и римские цифры одновременно!")
	} else if aRom == nil && bRom == nil {
		cNum := calculArab(aNum, bNum, y)
		fmt.Println(cNum)
	}
}

func calculArab(aNum, bNum int, y string) int {
	var cNum int
	switch y {
	case "+":
		cNum = aNum + bNum
		return cNum
	case "-":
		cNum = aNum - bNum
		return cNum
	case "*":
		cNum = aNum * bNum
		return cNum
	case "/":
		cNum = aNum / bNum
		return cNum
	default:
		panic("Неизвестный символ - 'y' ")
	}
}

func numRome(a, b, y string) {
	romnum := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}
	romnum2 := map[int]string{
		1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
		20: "XX", 30: "XXX", 40: "XL", 50: "L", 60: "LX", 70: "LXX", 80: "LXXX", 90: "XC", 100: "C",
	}
	aNum, okA := romnum[a]
	bNum, okB := romnum[b]
	if okA == false || okB == false {
		panic("Неизвестное Римское число! / Вводимое число не может быть больше X(10) и меньше I(1)")
	} else {
		cNum := calculRome(aNum, bNum, y)
		if cNum <= 0 {
			panic("Полученное число меньше I(1)")
		} else if cNum <= 10 {
			fmt.Println(romnum2[cNum])
		} else if cNum == 100 {
			fmt.Printf(romnum2[cNum])
		} else {
			zNum := cNum / 10
			oNum := zNum * 10
			vNum := cNum - oNum
			fmt.Printf("%s%s", romnum2[oNum], romnum2[vNum])
		}
	}
}

func calculRome(aNum, bNum int, y string) int {
	var cNum int
	switch y {
	case "+":
		cNum = aNum + bNum
		return cNum
	case "-":
		cNum = aNum - bNum
		return cNum
	case "*":
		cNum = aNum * bNum
		return cNum
	case "/":
		cNum = aNum / bNum
		return cNum
	default:
		panic("Неизвестный символ - 'y' ")
	}
}
