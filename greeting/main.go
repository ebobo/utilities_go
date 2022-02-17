package greeting

import (
	"fmt"
	"time"
)

func HelloQi() {
	fmt.Printf("Hello, Qi!  Today is %s \n", time.Now().Format("2006-January-02"))
}

func Hello(name string) {
	fmt.Printf("Hello, %s !  Today is %s \n", name, time.Now().Format("2006-January-02"))
}
