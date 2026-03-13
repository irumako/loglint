package loglint

// Ключевые слова для определения чувствительных данных.
var sensitiveDataKeywords = []string{
	"password",
	"passwd",
	"pwd",
	"secret",
	"api key",
	"api_key",
	"apikey",
	"authorization",
	"bearer",
	"credential",
	"private key",
	"session",
	"cookie",
	"ssn",
	"card",
	"cvv",
	"iban",
}
