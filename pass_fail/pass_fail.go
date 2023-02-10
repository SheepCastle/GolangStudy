// pass_fail 프로그램은 성적의 합격 여부를 알려 줍니다.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n') // 에러 반환 값을 다시 변수에 저장합니다.
	log.Fatal(err) // 에러를 보고 하고 프로그램을 종료합니다.
	fmt.Println(input)
}