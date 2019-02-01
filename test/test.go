package main

import (
	"fmt"
	"github.com/hlfstr/flagger"
	"os"
)

func main() {
	f := flagger.New()

	//Bool flag
	b := f.Bool("Test Bool", "-b", "--bool", "--boolean")

	//int flag
	i := f.Int(8, "Test Int", "-i", "--integer")

	//string flag
	s := f.String("h", "Test String", "-s", "--string")

	//help flag
	var help bool
	f.BoolVar(&help, "Show the Help", "-h", "--help")

	fmt.Printf("Before\n")
	fmt.Printf("  Bool:   %t\n", *b)
	fmt.Printf("  Int:    %d\n", *i)
	fmt.Printf("  String: %s\n", *s)
	d, err := f.Parse(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("After")
	fmt.Printf("  Bool:   %t\n", *b)
	fmt.Printf("  Int:    %d\n", *i)
	fmt.Printf("  String: %s\n", *s)
	fmt.Print("Data: ")
	fmt.Println(d)
	if help {
		f.Print("Available Flags: ")
	}
}
