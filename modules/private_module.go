package main

import (
	"fmt"
	
	schemaregcommon "gitlab.com/swissblock/schema-registry-golang"
)

func main() {
	versions := schemaregcommon.NewVersions()
	fmt.Printf("Versions: %+v \n", versions)
}