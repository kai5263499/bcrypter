package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	rs, _ := reader.ReadString('\n')
	pw := strings.TrimSpace(rs)
	hashed, _ := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)

	fmt.Printf("%s\n", hashed)
}
