package third

import "github.com/jimmyhollywood/first"
import "fmt"

func Third() {
	fmt.Printf("Third. Imported first\n")
	pack.DoNothing()
	fmt.Printf("Here should be first\n")
}
