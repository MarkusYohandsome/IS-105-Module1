package myquote

import "fmt"
import "rsc.io/quote"

func PrintQuotes() {
	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
}

// Her er koden du ville kjørt for å lage en go.mod fil som andre //
//kan hente ting fra
//go mod init github.com/MarkusYohandsome/IS-105-Module1
