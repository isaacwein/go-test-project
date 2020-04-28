package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func main() {
	uuid := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%d#@&*()%s", time.Now().UnixNano(), "in.Receiver"))))
	fmt.Println(uuid)
}
