package loglint

import "unicode"

func isEmoji(r rune) bool {
	switch {
	case r >= 0x1F300 && r <= 0x1FAFF:
		return true
	case r >= 0x2600 && r <= 0x26FF:
		return true
	case r >= 0x2700 && r <= 0x27BF:
		return true
	default:
		return false
	}
}

func isSpecialSymbol(r rune) bool {
	return unicode.IsSymbol(r) || unicode.IsPunct(r)
}
