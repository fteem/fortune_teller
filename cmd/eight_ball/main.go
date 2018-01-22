package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	ft "github.com/fteem/fortune_teller"
)

func main() {
	rand.Seed(time.Now().Unix())
	reader := bufio.NewReader(os.Stdin)
	answers := ft.Answers()
	fmt.Print("What is your question? ")
	reader.ReadString('\n')

	fmt.Println(answers[rand.Intn(len(answers))])
}
