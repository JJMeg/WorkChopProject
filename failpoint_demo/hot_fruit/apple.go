package hot_fruit

import (
	"fmt"
	"github.com/pingcap/failpoint"
)

func Apple() {
	fmt.Println("Apple...")

	failpoint.Inject("ApplePanic", func() {
		panic("Apple failpoint triggerd")
	})
}
