// pass_fail reports whether a grade is passing or failing
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade:")
	grade, err := getFloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade > 60.0 {
		status = "passing"
		fmt.Println("passing")
	} else {
		status = "failing"
		fmt.Println("failing")
	}
	fmt.Println("status =", status)

	bool, err := strconv.ParseBool("true")
	fmt.Println(bool, ":", err)

	file, err := os.Open("D:/Go/src/os/file.go")
	fmt.Println(file.Name(), ":", err)

	fileInfo, err := os.Stat("D:/Go/src/os/file.go")
	fmt.Println("size =", fileInfo.Size())
	if err != nil {
		log.Fatal(err)
	}

	response, err := http.Get("http://baidu.com/something_not_exists.html")
	fmt.Println(response.Body, ":", err)
}

func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, error := reader.ReadString('\n')
	if error != nil {
		return 0, error
	}
	fmt.Println(input)
	grade, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Should be a number")
		return 0, err
	}
	return grade, nil
}
