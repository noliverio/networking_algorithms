package crc32

import "fmt"

var Polynomial int = 0x04C11DB7
var message string = "This is my message, I sure hope it is not corrupted in transit. So I will use CRC-32 based error checking to validate the message's integrity. :)"

func Process() {
	fmt.Println(Polynomial)
	fmt.Println(message)
}
