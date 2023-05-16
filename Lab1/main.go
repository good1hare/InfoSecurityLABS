package main

import (
	"fmt"
)

func horizontalEncryption(text string, key int) string {
	ctext := ""
	for i := 0; i < key; i++ {
		for j := i; j < len(text); j += key {
			ctext += string(text[j])
		}
	}
	return ctext
}

func horizontalDecryption(text string, key int) string {
	ctext := make([]byte, len(text))
	k := 0
	for i := 0; i < key; i++ {
		for j := i; j < len(text); j += key {
			ctext[j] = text[k]
			k++
		}
	}
	return string(ctext)
}

// Задания к лабораторной работе
// 1.	Разработать алгоритм для шифрования сообщений горизонтальным перестановочным шифром;
// 2.	Разработать алгоритм для дешифрования сообщений зашифрованных горизонтальным перестановочным шифром;
// 3.	Составить приложение для шифрования/дешифрования с использованием горизонтальным перестановочного шифра;
func main() {
	plaintext1 := "Lorem Ipsum is simply dummy text of the printing and typesetting industry."
	key := 5
	ciphertext := horizontalEncryption(plaintext1, key)
	fmt.Println(ciphertext)

	plaintext := horizontalDecryption(ciphertext, key)
	fmt.Println(plaintext)
}
