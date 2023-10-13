package main

import (
	"fmt"
	"github.com/zaoshi00/examples-go/EmbedVersionInfo/cmd"
	"github.com/zaoshi00/examples-go/EmbedVersionInfo/version"
)

func main() {
	cmd.Execute()
	fmt.Println(version.GoVersion)
}
