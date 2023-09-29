package task1

import (
	"awesomeProject/module3/task1/somegetters"
	"awesomeProject/utility"
	"fmt"
	"github.com/huandu/xstrings"
)

func Task() {
	utility.AnnounceNewTask(3, 1)
	fmt.Printf("Random digit is %s\n", somegetters.Digit())
	fmt.Printf("Random city is %s\n", somegetters.City())
	fmt.Printf("Random city is %s\n", somegetters.City())
	fmt.Printf("Random digit is %s\n", somegetters.Digit())

	fmt.Printf("Random shuffled city is %s\n", xstrings.Shuffle(somegetters.City()))

}
