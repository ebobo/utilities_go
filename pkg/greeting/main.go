package greeting

import (
	"fmt"
	"time"
)

func HelloQi() {
	fmt.Printf("Hello, Qi!  Today is %s \n", time.Now().Format("2006-January-02"))
}
