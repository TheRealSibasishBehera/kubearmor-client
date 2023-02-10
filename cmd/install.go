// SPDX-License-Identifier: Apache-2.0
// Copyright 2021 Authors of KubeArmor

package cmd

import (
	"fmt"
	"github.com/kubearmor/kubearmor-client/install"
	"github.com/spf13/cobra"
)

var installOptions install.Options

// installCmd represents the get command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install KubeArmor in a Kubernetes Cluster",
	Long:  `Install KubeArmor in a Kubernetes Clusters`,
	RunE: func(cmd *cobra.Command, args []string) error {
		installOptions.Animation = true
		if err := installOptions.Env.CheckAndSetValidEnvironmentOption(cmd.Flag("env").Value.String()); err != nil {
			return fmt.Errorf("error in checking environment option: %v", err)
		}
		if err := install.K8sInstaller(client, installOptions); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.Flags().StringVarP(&installOptions.Namespace, "namespace", "n", "kube-system", "Namespace for resources")
	installCmd.Flags().StringVarP(&installOptions.KubearmorImage, "image", "i", "kubearmor/kubearmor:stable", "Kubearmor daemonset image to use")
	installCmd.Flags().StringVarP(&installOptions.Audit, "audit", "a", "", "Kubearmor Audit Posture Context [all,file,network,capabilities]")
	installCmd.Flags().BoolVar(&installOptions.Save, "save", false, "Save KubeArmor Manifest ")
	installCmd.Flags().StringVarP(&installOptions.Env.Environment, "env", "e", "", "Supported KubeArmor Environment [k3s,microK8s,minikube,gke,bottlerocket,eks,docker,oke,generic]")

}
