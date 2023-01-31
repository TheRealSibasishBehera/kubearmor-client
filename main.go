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
	u := checks.NewUpdateChecker()
	err := u.Init()
	if err != nil {
		fmt.Printf("Error checking for updates: %v\n", err)
		return
	}
	cmd.Execute()

}
