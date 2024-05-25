package utility

import "math/rand"

func Reverse(s string) (reversed string) {
    for _, character := range s {
        reversed = string(character) + reversed
    }
    return
}

func RandomQuotation() string {
    var quotations = []string{"Hello World!", "Hello Go!", "Hello Backend!"}
    randomIndex := rand.Intn(len(quotations))
    return quotations[randomIndex]
}

func Sum(firstNumber int, secondNumber int) int {
    return firstNumber + secondNumber
}