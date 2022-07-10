package main

import (
	"fmt"
	"os"

	"yiting.com/my_k8s_client/k8s_lib"
)

const COMMAND_LIST_NAMESPACES = "list_namespaces"
const SHORTED_COMMAND_LIST_NAMESPACES = "ln"

func main() {
	commandLineArgs := os.Args
	command := commandLineArgs[1]

	if command == COMMAND_LIST_NAMESPACES || command == SHORTED_COMMAND_LIST_NAMESPACES {
		namespaces := k8s_lib.ListNameSpaces()
		for idx, namespace := range namespaces.Items {
			fmt.Printf("%d: %s\n", idx, namespace.Name)
		}
	}
}
