package validation

import "github.com/badoux/checkmail"

//IsEmail verifique se uma string é um email válido
func IsEmail(v string) bool {
	if err := checkmail.ValidateFormat(v); err != nil {
		return false
	}
	return true
}
