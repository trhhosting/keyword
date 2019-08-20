package gogit

import (
	"testing"
	"fmt"
)
var defaultURL = "https://github.com/git-fixtures/basic.git"

func TestClone(t *testing.T) {
	err := Clone(defaultURL, "tmp")
	if err != nil {
		fmt.Println(err)
	   t.Fail()
	}
}