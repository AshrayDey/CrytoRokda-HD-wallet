package utils

import (
	"fmt"

	"github.com/bgadrian/go-mnemonic/bip39"
)

func GenerateMneunonic(passphrase string) string {
	newMneumonic, err := bip39.NewMnemonicRandom(128, passphrase)
	if err != nil {
		panic(fmt.Sprintf("Error in generating Mnemonic : %v", newMneumonic))
	}

	words, err := newMneumonic.GetSentence()
	if err != nil {
		panic(fmt.Sprintf("Words not generated : %v", err))
	}
	return words
}
