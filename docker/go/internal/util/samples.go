package util

import "fmt"

func Sample(a int, b int) string {
	res := a + b
	return fmt.Sprintf("from util.Sample! with %d", res)
}
