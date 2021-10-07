package main

import (
	"fmt"
	"testing"
)

func TestPrint1(t *testing.T) {
	date := Printdate()
	if date != 2021 {
		t.Errorf("value error")
	}
}

func TestPrint2(t *testing.T) {
	date := Printdate()
	if date != 2020 {
		t.Errorf("value error")
	}
}

func TestPrint(t *testing.T) {
	t.Run("print1", TestPrint1)
	t.Run("print2", TestPrint2)
}

func TestMain(m *testing.M) {
	fmt.Println("test begining...")
	m.Run()
}

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Printdate()
	}
}
