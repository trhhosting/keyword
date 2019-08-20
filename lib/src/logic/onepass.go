package logic

import (
	"os"
)
func CheckOnePass(onepass string) (bool) {
	if onepass == os.Getenv("ONEPASS"){
		return true
	} else {
		return false
	}
}
