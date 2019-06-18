package randString

import (
	"fmt"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	fmt.Println(RandStringBytes(8))
}
