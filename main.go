package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . something standard")
		return
	}

	file := "banners/standard.txt"

	if len(os.Args) == 3 {
		file = "banners/" + os.Args[2] + ".txt"
	}

	_, err := ValidateInput(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	banner, err := LoadBanner(file)
	if err != nil {
		fmt.Println("Error Reading File", err)
		return
	}

	fmt.Print(GenerateArt(os.Args[1], banner))
}