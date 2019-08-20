package logic

import (
	"fmt"
	"os"
	"testing"
)

func TestCheckOnePass(t *testing.T) {
	os.Setenv("ONEPASS", "Nunquam examinare secula.")
	x := CheckOnePass(os.Getenv("ONEPASS"))
	if x != true  {
		t.Error("test failed")
		return
	}
	fmt.Printf("CheckOnePass Passed\n")
}

