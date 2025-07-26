package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp,err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch : %v\n",err)
		}

		b ,err := io.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err!= nil {
			fmt.Fprintf(os.Stderr, "fetch: eading %s: %vv\n", url,err)
			os.Exit(1)
		}

		fmt.Printf("%s",b)
	}
	
	b := [5]int{2,1,4,5}
}