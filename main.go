package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := 0
	flag.IntVar(&port, "p", 8080, "port number(default:8080)")
	flag.Parse()
	if terminal.IsTerminal(0) {
		fmt.Println("pls send data over pipe.")
	} else {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("see here: localhost:" + strconv.Itoa(port))
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, string(b))
		})
		err = http.ListenAndServe(":"+strconv.Itoa(port), nil)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
