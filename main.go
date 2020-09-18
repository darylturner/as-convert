package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/darylturner/as-convert/asn"
)

func main() {
	var input string
	if len(os.Args) != 2 {
		inBytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimRight(string(inBytes), "\n")
	} else {
		input = os.Args[1]
	}
	if strings.Contains(input, ".") {
		out, err := asn.ConvertASdot(input)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(out)
	} else {
		out, err := asn.ConvertASplain(input)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(out)
	}
}
