package Functor

import (
	"fmt"
	"testing"
)

type testType struct {
	value int
}

func (test *testType) Print(value interface{}) {
	fmt.Println(value, test.value)
}

func Print(args []interface{}) {
	fmt.Printf("hello,world")
}

func TestFunctor(t *testing.T) {
	test := new(testType)
	test.value = 10
	t.Parallel()
	funcor := GetFunctor(test.Print, 1)
	funcor.RunFunc()
}

func TestA(t *testing.T) {
	t.Parallel()
	t.Skipf("111111111111111111111111111111111111111")
}

func TestB(t *testing.T) {
	t.Skipf("222222222222222222")
}

func Benchmark(b *testing.B) {
	arrA := make(map[int]int, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arrA[i] = i
		b.SetBytes(16)
	}
}
