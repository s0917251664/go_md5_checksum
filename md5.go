package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if count := len(os.Args); count > 1 {
		osIDX := 0
		for _, arg := range os.Args {
			osIDX++
			if osIDX <= 1 {
				continue
			}
			fmt.Printf("in [%s] \n", arg)
			f_byte, err := ioutil.ReadFile(arg)
			if err == nil {
				fmt.Printf("out [%s] : %x \n", arg, md5.Sum(f_byte))
			} else {
				fmt.Printf("[%s] read fail. err: %s \n", arg, err)
			}

		}
	} else {
		fmt.Println("please input file path .")
	}
}
