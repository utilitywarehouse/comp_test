package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"

	"github.com/golang/snappy"
)

func processPdf(fileName string) {
	fi, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s size %d", strings.Trim(fileName, "./"), fi.Size())
	src, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	encoded := snappy.Encode(nil, src)
	log.Printf("Compressed size %d", int64(len(encoded)))

	fs := float64(fi.Size())
	el := float64(len(encoded))
	diff := (el - fs) / fs * 100
	log.Printf("Reduction %F%%", math.Abs(math.Round(diff*100)/100))
}

func main() {
	processPdf("./small.pdf")
	fmt.Println("")
	processPdf("./large.pdf")
}
