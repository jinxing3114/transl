package transl

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	fmt.Println(T("i don't no"))
	fmt.Println(T("Delete"))
	fmt.Println(T("Yes", "fi"))
	fmt.Println(T("No", "id"))
}

func Benchmark_t(b *testing.B){
	for i := 0; i < b.N; i++ {
		T("i don't no")
		T("Delete")
		T("Yes", "fi")
		T("No", "id")
	}
}
