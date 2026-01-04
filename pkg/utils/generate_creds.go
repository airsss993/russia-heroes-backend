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

// GenerateRandString - генерирует криптографически случайные строки
//
// Принимает:
// - int - кол-во символом для выходной строки
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

// GenerateAdminCredentials - генерирует новый набор учетных данных администратора с криптографически стойкими случайными именем пользователя и паролем
//
// Всегда возвращает nil в качестве ошибки для обратной совместимости
func GenerateAdminCredentials() Credentials {
	username := GenerateRandString(lenghtUsername)
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
