package main

import "fmt"

func encryptVigenere(plaintext string, key string) string {
	ciphertext := ""
	keyIndex := 0
	for _, char := range plaintext {
		if char >= 'a' && char <= 'z' {
			var shift int32 = int32(key[keyIndex] - 'a')
			ciphertext += string('a' + ((char-'a')+shift)%26)
			keyIndex = (keyIndex + 1) % len(key)
		} else if char >= 'A' && char <= 'Z' {
			var shift int32 = int32(key[keyIndex] - 'a')
			ciphertext += string('A' + ((char-'A')+shift)%26)
			keyIndex = (keyIndex + 1) % len(key)
		} else {
			ciphertext += string(char)
		}
	}
	return ciphertext
}

func decryptVigenere(ciphertext string, key string) string {
	plaintext := ""
	keyIndex := 0
	for _, char := range ciphertext {
		if char >= 'a' && char <= 'z' {
			var shift int32 = int32(key[keyIndex] - 'a')
			plaintext += string('a' + ((char-'a')-shift+26)%26)
			keyIndex = (keyIndex + 1) % len(key)
		} else if char >= 'A' && char <= 'Z' {
			var shift int32 = int32(key[keyIndex] - 'A')
			plaintext += string('A' + ((char-'A')-shift+26)%26)
			keyIndex = (keyIndex + 1) % len(key)
		} else {
			plaintext += string(char)
		}
	}
	return plaintext
}

func main() {
	plaintext := "thequickbrownfoxjumpsoverthelazydog"
	key := "secret"
	fmt.Println("Plaintext: ", plaintext)
	fmt.Println("Key: ", key)
	ciphertext := encryptVigenere(plaintext, key)
	fmt.Println("Ciphertext: ", ciphertext)
	decryptedText := decryptVigenere(ciphertext, key)
	fmt.Println("Decrypted text: ", decryptedText)
}
