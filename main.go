package main

import (
	"fmt"
	"os"

	"yiting.com/my_k8s_client/k8s_lib"
)

const COMMAND_LIST_NAMESPACES = "list_namespaces"
const SHORTED_COMMAND_LIST_NAMESPACES = "ln"
const COMMAND_CREATE_NAMESPACE = "create_namespace"
const SHORTED_COMMAND_CREATE_NAMESPACE = "cn"

func main() {
	commandLineArgs := os.Args
	command := commandLineArgs[1]

	// list namespaces
	if command == COMMAND_LIST_NAMESPACES || command == SHORTED_COMMAND_LIST_NAMESPACES {
		namespaces := k8s_lib.ListNamespaces()
		for idx, namespace := range namespaces.Items {
			fmt.Printf("%d: %s\n", idx, namespace.Name)
		}
	}

	// create namespace
	if command == COMMAND_CREATE_NAMESPACE || command == SHORTED_COMMAND_CREATE_NAMESPACE {
		new_namespace := commandLineArgs[2]
		k8s_lib.CreateNewNamespace(new_namespace)
	}
}
