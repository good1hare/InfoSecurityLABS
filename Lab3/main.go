package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "secretmessage" // текст, который нужно зашифровать
	keyword := "golang"     // кодовое слово

	// приведем весь текст и кодовое слово к нижнему регистру
	text = strings.ToLower(text)
	keyword = strings.ToLower(keyword)

	// создаем слайсы для хранения зашифрованного и дешифрованного текста
	encrypted := make([]byte, len(text))
	decrypted := make([]byte, len(text))

	// проходимся по всем символам в тексте
	for i := 0; i < len(text); i++ {
		// вычисляем номер символа в алфавите текста
		textChar := text[i] - 'a'
		// вычисляем номер символа в алфавите кодового слова
		keywordChar := keyword[i%len(keyword)] - 'a'
		// вычисляем номер зашифрованного символа в алфавите
		encryptedChar := (textChar + keywordChar) % 26
		// вычисляем номер дешифрованного символа в алфавите
		decryptedChar := (encryptedChar - keywordChar + 26) % 26
		// приводим номера символов обратно к символам в алфавите
		encrypted[i] = byte(encryptedChar) + 'a'
		decrypted[i] = byte(decryptedChar) + 'a'
	}

	// выводим зашифрованный и дешифрованный текст
	fmt.Println("Зашифрованный текст: ", string(encrypted))
	fmt.Println("Дешифрованный текст: ", string(decrypted))
}
