package main

import "fmt"

func main() {
	names := map[string][]string{
		"bond_James":      {"shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}

	for _, value := range names {
		for i, v := range value {
			fmt.Println("index: ", i, " value: ", v)
		}
	}
}
