package pointer

import "fmt"

func PointerPractice() {

	var num int = 10
	fmt.Println("num:", num)       // 10
	fmt.Println("numのアドレス:", &num) // 0xc0000b6010

	var ptr *int = &num
	fmt.Println("ptr（ptrに格納しているnumのアドレス）:", ptr)            // 0xc0000b6010
	fmt.Println("ptr（ptrに格納しているnumのアドレスに格納されている）の値:", *ptr) // 10
	fmt.Println("ptrのアドレス（変数ptr自身のアドレス）:", &ptr)            // 0xc0000ae020
}
