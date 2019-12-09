package fruit

import (
	"fmt"
	"github.com/pingcap/failpoint"
)

func Banana() {
	fmt.Println("banana...")

	failpoint.Inject("BananaPanic", func() {
		panic("banana failpoint triggerd")
	})
}
