package main

import (
	"flag"
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

func main() {
	var (
		i = flag.Int("int", 1, "int flag")
		s = flag.String("str", "default", "string flag")
		b = flag.Bool("bool", true, "bool flag")
	)
	flag.Parse()
	fmt.Println(*i, *s, *b)
	fmt.Println(randomdata.SillyName())
}
