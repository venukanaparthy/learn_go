//test.go
package tests

import (
	"fmt"
	"os"
	"testing"
)

func TestAssertion(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(homeDir)

}
