package myquote

import "fmt"
import "rsc.io/quote"

func PrintQuotes() {
	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
}
