package main

import "fmt"

func main() {

	plaintext := "LABORATORYWORK"
	keyword := "QWERTY"

	j := 0
	n := len(plaintext)
	res := ""
	for i := 0; i < n; i++ {
		r := plaintext[i]
		res += string((r+keyword[j]-2*'A')%26 + 'A')

		j = (j + 1) % len(keyword)
	}
	fmt.Println("Зашифрованный текст: ", res)

}
