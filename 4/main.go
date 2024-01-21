package main

import "fmt"

func main() {

	var num int = 10

	fmt.Println("numのアドレス:", &num) //numのアドレス: 0x1400009c018
	fmt.Println("numの値:", num)     //numの値: 10

	var ptr *int = &num //numのアドレスをptrに代入 (変数numのアドレスを格納する変数ptrを宣言)

	fmt.Println("ptrが指し示すアドレス:", ptr) //ptrが指し示すアドレス: 0x1400009c018 (numのアドレスと同じ) ← ここがポイント
	fmt.Println("ptrが指し示す値:", *ptr)   //ptrが指し示す値: 10
	fmt.Println("ptr自体のアドレス:", &ptr)  //ptr自体のアドレス: 0x1400009c020

	*ptr = 20 //numの値を20に変更

	fmt.Println("numの新しい値:", num) //numの新しい値: 20

}
