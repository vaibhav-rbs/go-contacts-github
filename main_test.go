package main

import (
	"fmt"
	"strconv"
	"testing"
)

func addUsers(count int) {
	if count < 1 {
		count = 1
	}
	for i := 0; i < count; i++ {
		statement := fmt.Sprintf("INSERT INTO users(name, age) VALUES('%s', %d)", ("User " + strconv.Itoa(i+1)), ((i + 1) * 10))
		a.DB.Exec(statement)
	}
}

func TestRest(t *testing.T) {
	tt := []struct {
		input  string
		output string
	}{}
}
