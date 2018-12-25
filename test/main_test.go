package test

import (
	"testing"
	"fmt"
)

func testPrint(t *testing.T) {
	ret := PrintMain()
	if ret == 0 {
		t.Errorf("error")
	}

	fmt.Println(ret)
}
func testPrint2(t *testing.T) {
	ret := PrintMain()
	if ret == 0 {
		t.Errorf("error")
	}
	ret++
	fmt.Println(ret)
}

func TestAll(t *testing.T) {
	t.Run("testPrint", testPrint)
	t.Run("testPrint2", testPrint2)
}

func TestMain(m *testing.M) {
	fmt.Println("Tests begins")
	m.Run()
}

func testb(n int) int {
	for n > 0 {
		n--
	}
	return n
}
//也受TestMain 控制
func BenchmarkAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testb(n)
	}
}
