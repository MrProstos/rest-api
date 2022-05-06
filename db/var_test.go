package db

import (
	"fmt"
	"testing"
)

func TestNewOperator(t *testing.T) {
	s, err := NewOperator("test", "token")
	fmt.Println(s, err)
}
