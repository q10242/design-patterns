package main
import ("fmt"
"github.com/q10242/design-patterns/singleton"
)

func main() {
	instance1 := singleton.GetInstance()
	instance2 := singleton.GetInstance()
	if instance1 == instance2 {
		fmt.Println("instance1 and instance2 are same")
	} else {
		fmt.Println("instance1 and instance2 are not same")
	}
}