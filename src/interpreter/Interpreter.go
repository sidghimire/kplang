package interpreter

import "fmt"

func PRINT(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Print(v)
	case int:
		fmt.Print(v)
	case float64:
		fmt.Println(v)
	default:
		fmt.Println("Unsupported type")
	}
}
func PRINTLN(value interface{}) {

	switch v := value.(type) {
	case string:
		fmt.Println(v)
	case int:
		fmt.Println(v)
	case float64:
		fmt.Println(v)
	default:
		fmt.Println("Unsupported type")
	}
}
