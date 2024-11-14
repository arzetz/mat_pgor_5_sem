package main

import (
	"fmt"
	"strings"
)

func cypher(plaintext string, key string) string {
	key = strings.ToLower(key)
	plaintext = strings.ToLower(plaintext)
	var encrypted strings.Builder
	keyindex := 0
	for _, i := range plaintext {
		p := byte(i) - 'a'
		k := key[keyindex] - 'a'
		c := (p + k) % 26
		encrypted.WriteByte(c + 'a')
		keyindex = (keyindex + 1) % len(key)
	}
	return encrypted.String()
}

func decypher(encrypted string, key string) string {
	key = strings.ToLower(key)
	encrypted = strings.ToLower(encrypted)
	var plaintext strings.Builder
	keyindex := 0
	for _, i := range encrypted {
		c := byte(i) - 'a'
		k := key[keyindex] - 'a'
		p := (c - k + 26) % 26
		plaintext.WriteByte(p + 'a')
		keyindex = (keyindex + 1) % len(key)
	}

	return plaintext.String()
}

func main() {
	fmt.Printf("Зашифрованное сообщение %s \n", cypher("lol", "kek"))
	fmt.Printf("Дешифрованное сообщение %s", decypher("vsv", "kek"))
}
