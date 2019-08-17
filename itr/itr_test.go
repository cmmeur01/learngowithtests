package itr

import (
	"fmt"
	"testing"
)

func TestRepeater(t *testing.T) {
	repeated := Repeater("a", 3)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expected %q got %q", repeated, expected)
	}
}

// func BenchmarkRepeater(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Repeater("a", 7)
// 	}
// }

func ExampleRepeater() {
	repeated := Repeater("b", 7)
	fmt.Println(repeated)
	//Output: bbbbbbb
}
