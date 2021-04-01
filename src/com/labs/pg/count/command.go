package main

import (
	"com/labs/pg/datafile"
	"fmt"
	"log"
	"sort"
)

func main() {
	lines, err := datafile.GetString("src/com/labs/pg/datafile/vote.txt")
	if err != nil {
		log.Fatal(err)
	}

	counts := map[string]int{}

	for _, line := range lines {
		counts[line]++
	}

	fmt.Println(counts)

	for key, val := range counts {
		fmt.Printf("key=%s,val=%d\n", key, val)
	}

	grades := map[string]float64{"Alma": 74.2, "Robin": 86.5, "Carl": 59.7}

	names := []string{}

	for name := range grades {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s has a grade of %0.1f\n", name, grades[name])
	}

}
