package password

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)


var Letters = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

var Numbers = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

var Symbols = []string{
	"(", ")", "{", "}", "[", "]", "<", ">", ",", ".", ";", ":", "'", "\"", "=", "+", "-", "*", "/", "\\", "!", "?", "%", "&", "|", "^", "_", "#", "@", "$", "`", "~",
}

func RandomSelector(num int, slice []string) string {
	passString := []string{}
	var password string

	rand.Seed(time.Now().Unix())
	for i := 0; i < num; i++ {
		randomElement := slice[rand.Intn(len(slice))]
		passString = append(passString, randomElement)
	}

	for _, element := range passString {
		password += fmt.Sprint(element)
	}

	return password
}

func ShufflePass(text string) string {
	textChars := []rune(text)
	rand.Seed(time.Now().UnixNano())

	n := len(textChars)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		textChars[i], textChars[j] = textChars[j], textChars[i]
	}

	shuffledText := string(textChars)
	return shuffledText
}

func DivideEqually(size int) (int, int, int) {
	var nOfLetters int
	var nOfNumbers int
	var nOfSymbols int

	if size%2 == 0 {
		nOfLetters = size / 2
		nOfNumbers = (size - nOfLetters) * 2 / 3
		nOfSymbols = size - nOfLetters - nOfNumbers
	} else {
		floatValue := size / 2
		ceilResult := math.Ceil(float64(floatValue))
		nOfLetters = int(ceilResult)
		nOfNumbers = (size - nOfLetters) * 2 / 3
		nOfSymbols = size - nOfNumbers - nOfLetters
	}
	return nOfLetters, nOfNumbers, nOfSymbols
}

func GeneratePassword(size int) string {
	nOfLetters, nOfNumbers, nOfSymbols := DivideEqually(size)
	passLetters := RandomSelector(nOfLetters, Letters)
	passSymbols := RandomSelector(nOfSymbols, Symbols)
	passNum := RandomSelector(nOfNumbers, Numbers)

	conPass := fmt.Sprintf("%s%s%s", passLetters, passSymbols, passNum)
	password := ShufflePass(conPass)
	return password
}
