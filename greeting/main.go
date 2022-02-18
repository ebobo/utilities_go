package greeting

import (
	"fmt"
	"time"
)

//HelloQi greeting to Qi
func HelloQi() {
	fmt.Printf("Hello, Qi!  Today is %s \n", time.Now().Format("2006-January-02"))
}

//Hello greeting to (string)
func Hello(name string) {
	fmt.Printf("Hello, %s !  Today is %s \n", name, time.Now().Format("2006-January-02"))
}

//Foreword statement (string)
func Foreword(words string) {
	fmt.Printf("%s ! \n  Today is %s \n", words, time.Now().Format("2006-January-02"))
}
