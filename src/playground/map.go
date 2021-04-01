package main

import "fmt"

func main() {
	var ok bool
	ranks := map[string]int{}
	var rank int
	ranks["brozen"] = 3
	rank, ok = ranks["brozen"]
	fmt.Printf("rank: %d, ok:%v\n", rank, ok)
	delete(ranks, "brozen")
	rank, ok = ranks["brozen"]
	fmt.Printf("rank: %d, ok:%v\n", rank, ok)

	isPrime := map[int]bool{}
	var prime bool
	isPrime[5] = true
	prime, ok = isPrime[5]
	fmt.Printf("prime:%v,ok:%v\n", prime, ok)
	delete(isPrime, 5)
	prime, ok = isPrime[5]
	fmt.Printf("prime:%v,ok:%v\n", prime, ok)

}
