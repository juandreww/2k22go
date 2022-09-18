package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str := "Love is but a song to sing Fear's the way we die You can make the mountains ring Or make the angels cry Though the bird is on the wing And you may not know why Come on people now Smile on your brother Everybody get together Try to love one another Right now"
	encodeStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(str))
	fmt.Println(len(str))
	fmt.Println(len(s64))
	fmt.Println(str)
	fmt.Println(s64)
}
