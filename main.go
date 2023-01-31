// SPDX-License-Identifier: Apache-2.0
// Copyright 2021 Authors of KubeArmor

// Package main is responsible for the execution of CLI
package main

import (
	"fmt"
	"github.com/kubearmor/kubearmor-client/checks"
	"github.com/kubearmor/kubearmor-client/cmd"
)

func main() {
	if err := checks.CheckForUpdates(); err != nil {
		fmt.Println("Error checking for updates:", err)
		return
	}
	cmd.Execute()

}
