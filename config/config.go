package config

import (
	"fmt"
	"os"
	"strings"
)

// GetSecret читает секрет из файла (Docker Secret) или переменной окружения
func GetSecret(secretName string) (string, error) {
	secretPath := fmt.Sprintf("/run/secrets/%s", secretName)

	// Если файл существует (Docker Secrets)
	if data, err := os.ReadFile(secretPath); err == nil {
		return strings.TrimSpace(string(data)), nil
	}

	// Если файла нет, читаем из ENV (локальная отладка)
	secretValue := os.Getenv(secretName)
	if secretValue == "" {
		return "", fmt.Errorf("secret %s not found", secretName)
	}
	return secretValue, nil
}
