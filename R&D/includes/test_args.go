package includes

import (
	"fmt"
	"os"
)

func checkEnvironment() {
	var ENVConfig string
	CLIConfig := os.Args
	if len(CLIConfig) > 1 {
		switch CLIConfig[1] {
		case "DEBUG":
			ENVConfig = "DEBUG"
		case "PROD":
			ENVConfig = "PROD"
		}
	} else {
		ENVConfig = "PROD"
	}
	fmt.Println(ENVConfig)
}
