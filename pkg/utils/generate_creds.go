package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	lenghtUsername = 12                                                               // Определяет длину генерируемого имени пользователя
	lenghtPassword = 16                                                               // Определяет длину генерируемого пароля
	symbols        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // Содержит набор символов, используемых для генерации учетных данных
)

// Credentials представляет набор учетных данных пользователя
type Credentials struct {
	Username string // Сгенерированное имя пользователя
	Password string // Сгенерированный пароль
}

// Генерирует криптографически стойкий случайный пароль
//
// Возвращает пустую строку в случае ошибки генерации случайного числа
func generatePassword() string {
	b := make([]byte, lenghtPassword)
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(symbols))))
		if err != nil {
			return ""
		}
		b[i] = symbols[n.Int64()]
	}

	return string(b)
}

// Генерирует криптографически стойкое случайное имя пользователя
//
// Возвращает пустую строку в случае ошибки генерации случайного числа
func generateUsername() string {
	b := make([]byte, lenghtUsername)
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(symbols))))
		if err != nil {
			return ""
		}
		b[i] = symbols[n.Int64()]
	}
	return string(b)
}

// GenerateAdminCredentials Генерирует новый набор учетных данных администратора с криптографически стойкими случайными именем пользователя и паролем
//
// Всегда возвращает nil в качестве ошибки для обратной совместимости
func GenerateAdminCredentials() Credentials {
	username := generateUsername()
	password := generatePassword()
	return Credentials{
		Username: username,
		Password: password,
	}
}

// PrintCredentials выводит учетные данные администратора в консоль
//
// Отображает сгенерированные username и password с предупреждением о необходимости их сохранения
func PrintCredentials(creds Credentials) {
	fmt.Println("\n=== Учетные данные администратора ===")
	fmt.Printf("Логин: %s\n", creds.Username)
	fmt.Printf("Пароль: %s\n", creds.Password)
	fmt.Println("\n⚠️  ВАЖНО: Сохраните эти данные в безопасном месте!")
	fmt.Println("Повторный вывод будет невозможен.\n")
}
