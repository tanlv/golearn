package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Khai bao global
var koreanGreeting string = "안녕, 친구"

// khai báo biến và gán giá trị
var enGreeting = "Hello"

// Cách khai báo biến trong func thường dùng
var cnGreeting = "你好朋友"

func getGreeting(language string) {
	// If/else condition
	if language == "en" {
		fmt.Println(enGreeting)
	} else if language == "kr" {
		fmt.Println(koreanGreeting)
	} else if language == "cn" {
		fmt.Println(cnGreeting)
	} else {
		fmt.Println("Language is not support")
	}
}

func chooseSecretPerson() {
	var answer int
	var perZero, perOne, perTwo string
	perZero = "Max"
	perOne = "Alex"
	perTwo = "Tom"

	// Show number for person
	fmt.Println("Number of secret person 0: " + perZero + ", 1: " + perOne + ", 2: " + perTwo + ".")

	// Random in Go
	source := rand.NewSource(time.Now().UnixNano()) // Can tra google ve viec random trong Go
	r := rand.New(source)
	//fmt.Println(r.Intn(3))

	// Gán secret person
	secretPer := r.Intn(3)
	//fmt.Println(secretPer)

	// Đọc giá trị từ người dùng nhập vào
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Please enter your answer with numer 0 - 2: ")
	//input, _ := reader.ReadString('\n')
	//input = strings.TrimSuffix(input, "\n")
	// Convert string to int
	//answer, err := strconv.Atoi(input)

	// một cách viết khác để lấy input do KH nhập vào
	scanner := bufio.NewScanner(os.Stdin)
	// Check input number
	for {
		fmt.Printf("Please enter your answer with numer 0 - 2: ")
		scanner.Scan()
		input := scanner.Text()
		answer, err := strconv.Atoi(input)

		if answer >= 0 && answer <= 2 {
			break
		}

		// Check answer neu có lỗi
		if err != nil {
			//fmt.Printf("Error is: %v \n", err)
			log.Panic(err)
		}

	}

	// switch case
	var secretName string
	switch secretPer {
	case 0:
		secretName = perZero
	case 1:
		secretName = perOne
	case 2:
		secretName = perTwo
	default:
		secretName = "No Name"
	}

	// Run Game
	if secretPer == answer {

		fmt.Printf("Congratulations!! You answer is correct, %s is the secret person.\n", secretName)
	} else {

		fmt.Printf("Sorry!! You answer is incorrect,%d: %s is the secret person.\n", secretPer, secretName)
	}
}

func showMultiTable() {
	var number1, number2 int
	var err1, err2 error
	// Người dùng nhập vào bảng cửu chương muốn xem
	// Người dùng nhập vào khoảng ranger muốn in
	// Chương trình in ra từng dòng kết quả

	// Input nhập bản cửu chương cần xem
	scanner1 := bufio.NewScanner(os.Stdin)
	// check giá trị đầu vào
	for {
		// lấy giá trị
		fmt.Printf("Nhập vào bảng cửu chương muốn xem: ")
		scanner1.Scan()
		input1 := scanner1.Text()

		if input1 == "" {
			continue
		}

		// Chuyển string to int
		number1, err1 = strconv.Atoi(input1)

		if number1 >= 1 && number1 <= 100 {
			break
		}

		if err1 != nil {
			log.Panic(err1)
		}

	}

	// Input nhập ranger cần xem
	scanner2 := bufio.NewScanner(os.Stdin)

	// check giá trị đầu vào
	for {
		fmt.Printf("Nhập vào khoảng ranger của bảng cửu chương: ")
		scanner2.Scan()
		input2 := scanner2.Text()
		if input2 == "" {
			continue
		}
		number2, err2 = strconv.Atoi(input2)

		if number2 >= 1 && number2 <= 20 {
			break
		}

		if err2 != nil {
			log.Panic(err2)
		}

	}

	// In ra màn hình bảng cửu chương
	fmt.Printf("Bảng Cửu Chương: %d\n", number1)
	inc := 0
	for {
		res := number1 * inc
		// vòng lặp print bảng cửu chương
		fmt.Printf("%[1]d x %[2]d = %[3]d\n", number1, inc, res)
		inc++
		if inc > number2 {
			break
		}
	}

}

func main() {

	// Cần tìm hiểu về cách sử dụng flag
	lang := flag.String("lang", "en", "greeting language")
	flag.Parse()

	// function get greeting
	getGreeting(*lang)

	// Func game Mini game 1: Chọn người bí mật
	//fmt.Println("Game \"Chọn Người bí mật\"")
	//chooseSecretPerson()
	// Bảng cửu chương
	showMultiTable()

}
