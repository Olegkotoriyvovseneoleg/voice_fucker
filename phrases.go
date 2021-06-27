
package main

import (
	"math/rand"
)
var badWords = []string{
	"в очко себе это надиктуй",
	"что за ебанный писк?",
	"у меня кошка, когда блюет, ито звучит лучше",
	"судя по интонации у тебя явные признаки ДЦП",
}

// GetBadWord ...
func GetBadWord() string {
	return badWords[rand.Intn(len(badWords))]
}
