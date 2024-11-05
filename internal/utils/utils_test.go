package utils

import (
	"fmt"
	"testing"
)

func TestCheckHashPassword(t *testing.T) {
	password := "foo"
	hash := HashPassword("foo")
	result := CheckHashPassword(hash, password)
	if !result {
		t.Fatalf("CheckHashPassword should returns true but return %t", result)
	}

}

func ExampleCheckHashPassword() {
	hash := HashPassword("bar")
	// Output: true
	fmt.Println(CheckHashPassword(hash, "bar"))
}
