package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var a int8 = 127         // -128 to 127
	var b int16 = 32767      // -32768 to 32767
	var c int32 = 2147483647 // -2147483648 to 2147483647
	d := 9223372036854775807 // -9223372036854775808 to 9223372036854775807

	fmt.Println("a: ", a, unsafe.Sizeof(a))
	fmt.Println("b: ", b, unsafe.Sizeof(b))
	fmt.Println("c: ", c, unsafe.Sizeof(c))
	fmt.Println("d: ", d, unsafe.Sizeof(d))

	var (
		a1 uint8  = 255                  // 0 to 255
		b1 uint16 = 65535                // 0 to 65535
		c1 uint32 = 4294967295           // 0 to 4294967295
		d1 uint64 = 18446744073709551615 // 0 to 18446744073709551615
	)

	fmt.Println("a1: ", a1, unsafe.Sizeof(a1))
	fmt.Println("b1: ", b1, unsafe.Sizeof(b1))
	fmt.Println("c1: ", c1, unsafe.Sizeof(c1))
	fmt.Println("d1: ", d1, unsafe.Sizeof(d1))
	fmt.Printf("Data type of variable a1 is %T", a1)
	fmt.Printf("Data type of variable b1 is %T", b1)
	fmt.Printf("Data type of variable c1 is %T", c1)
	fmt.Printf("Data type of variable d1 is %T", d1)
}
