package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	lenghtUsername = 8                                                                // Определяет длину генерируемого имени пользователя
	lenghtPassword = 16                                                               // Определяет длину генерируемого пароля
	symbols        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // Содержит набор символов, используемых для генерации учетных данных
)

// Credentials представляет набор учетных данных пользователя
type Credentials struct {
	Username string // Сгенерированное имя пользователя
	Password string // Сгенерированный пароль
}

// GenerateRandString - генерирует криптографически случайные строки
//
// Принимает:
// - cost - кол-во символов для выходной строки
//
// Возвращает:
// - string - сгенерированную строку или
// пустую строку в случае ошибки генерации случайного числа.
func GenerateRandString(cost int) []byte {
	b := make([]byte, cost)
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(symbols))))
		if err != nil {
			return []byte{}
		}
		b[i] = symbols[n.Int64()]
	}

	return b
}

// GenerateUsername генерирует уникальное имя админа в формате "admin_{8 случайных символов}"
//
// Возвращает:
// - []byte - срез байт типа "admin_a7b3c9d2"
func GenerateUsername() []byte {
	prefix := []byte("admin_")
	b := make([]byte, len(prefix)+lenghtUsername)
	copy(b, prefix)

	for i := len(prefix); i < len(b); i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(symbols))))
		if err != nil {
			return nil
		}
		b[i] = symbols[n.Int64()]
	}

	return b
}

// GenerateAdminCredentials - генерирует новый набор учетных данных администратора с криптографически стойкими случайными именем пользователя и паролем
//
// Возвращает:
// - Credentials - структуру с заполненными данными для логина
func GenerateAdminCredentials() Credentials {
	// Создаем случайные имя и пароль пользователя
	username := GenerateUsername()
	password := GenerateRandString(lenghtPassword)
	return Credentials{
		Username: string(username),
		Password: string(password),
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
