package utils

import "testing"

// func Test_SelectOne(t *testing.T) {

// 	SelectOne("test", []string{"a", "b", "c"})
// }

func Test_TableWriter(t *testing.T) {

	terminal := NewTerminal("")

	terminal.TableWriter([]string{"a", "b", "c"}, [][]string{{"1", "2", "3"}})
}
