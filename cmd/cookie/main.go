package main

import (
	"fmt"
	"math/rand"
	"time"

	ft "github.com/fteem/fortune_teller"
)

func main() {
	rand.Seed(time.Now().Unix())
	quotes := ft.Quotes()
	fmt.Println(quotes[rand.Intn(len(quotes))])
}
