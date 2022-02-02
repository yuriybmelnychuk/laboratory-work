package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	var (
		txt, key string
		mode     int32
	)

	fmt.Println("Введите текст (лат.) для зашифровки или расшифровки")
	fmt.Scanf("%s\n", &txt) // Ожидаем ввода текста (string)

	fmt.Println("Введите ключ")
	fmt.Scanf("%s\n", &key) // Ожидаем ввода ключа (string)

	fmt.Println("Введите 'e' чтобы зашифровать или 'd' чтобы расшифровать")
	fmt.Scanf("%c", &mode) // Ожидаем ввода режима (character)

	if mode == int32('e') { // Если нажата "e"
		res, err := Encrypt(txt, key)
		if err != nil {
			fmt.Println(err.Error()) // Вывод сообщения об ошибке
			return
		}
		fmt.Println("Зашифрованный текст: ", res) // Передаем полученные значения на зашифровку
	} else if mode == int32('d') { // Если нажата "d"
		res, err := Decrypt(txt, key)
		if err != nil {
			fmt.Println(err.Error()) // Вывод сообщения об ошибке
			return
		}
		fmt.Println("Расшифрованный текст: ", res) // Передаем полученные значения на расшифровку
	} else {
		fmt.Println("Ошибка ввода") // Или вывод сообщения об ошибке
	}
}

// Зашифровка текста
func Encrypt(text string, key string) (string, error) {
	if err := checkKey(key); err != nil { // Проверка ключа
		return "", errors.New("не удалось зашифровать(")
	}

	keyLower := strings.ToLower(key)
	keyUpper := strings.ToUpper(key)

	j := 0
	n := len(text)
	res := ""
	for i := 0; i < n; i++ {
		r := text[i]
		if r >= 'a' && r <= 'z' {
			res += string((r+keyLower[j]-2*'a')%26 + 'a') // вычисление для строчных букв
		} else if r >= 'A' && r <= 'Z' {
			res += string((r+keyUpper[j]-2*'A')%26 + 'A') // вычисление для заглавных букв
		} else {
			res += string(r)
		}

		j = (j + 1) % len(key)
	}

	return res, nil
}

// Расшифровка текста
func Decrypt(text string, key string) (string, error) {
	if err := checkKey(key); err != nil { // Проверка ключа
		return "", errors.New("не удалось расшифровать(")
	}

	keyLower := strings.ToLower(key)
	keyUpper := strings.ToUpper(key)

	res := ""
	j := 0
	n := len(text)
	for i := 0; i < n; i++ {
		r := text[i]
		if r >= 'a' && r <= 'z' {
			res += string((r-keyLower[j]+26)%26 + 'a') // вычисление для строчных букв
		} else if r >= 'A' && r <= 'Z' {
			res += string((r-keyUpper[j]+26)%26 + 'A') // вычисление для заглавных букв
		} else {
			res += string(r)
		}

		j = (j + 1) % len(key)
	}

	return res, nil
}

// Проверка ключа
func checkKey(key string) error {
	if len(key) == 0 {
		return errors.New("длина ключа равняется 0")
	}

	for _, r := range key {
		if !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z') { // Проверка буквенного диапазона
			return errors.New("неверный ключ")
		}
	}

	return nil
}
