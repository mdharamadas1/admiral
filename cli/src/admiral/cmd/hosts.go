/*
 * Copyright (c) 2016 VMware, Inc. All Rights Reserved.
 *
 * This product is licensed to you under the Apache License, Version 2.0 (the "License").
 * You may not use this product except in compliance with the License.
 *
 * This product may include a number of subcomponents with separate copyright notices
 * and license terms. Your use of these subcomponents is subject to the terms and
 * conditions of the subcomponent's license, as noted in the LICENSE file.
 */

package cmd

import (
	"admiral/hosts"
	"fmt"

	"admiral/help"

	"admiral/functions"

	"github.com/spf13/cobra"
)

func init() {
	initHostAdd()
	initHostDisable()
	initHostEnable()
	initHostRemove()
	initHostUpdate()
	initHostList()
}

var hostAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add host",
	Long:  "Add host",

	Run: func(cmd *cobra.Command, args []string) {
		var (
			newID string
			err   error
		)
		newID, err = hosts.AddHost(ipF, resPoolF, deplPolicyF, credName, publicCert, privateCert, userName, passWord,
			autoAccept,
			custProps)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Host added: " + newID)
		}
	},
}

func initHostAdd() {
	hostAddCmd.Flags().StringVar(&publicCert, "public", "", "(Required if adding new credentials)"+publicCertDesc)
	hostAddCmd.Flags().StringVar(&privateCert, "private", "", "(Required if adding new credentials)"+privateCertDesc)
	hostAddCmd.Flags().StringVar(&userName, "username", "", "(Required if adding new credentials)"+"Username.")
	hostAddCmd.Flags().StringVar(&passWord, "password", "", "(Required if adding new credentials)"+"Password.")
	hostAddCmd.Flags().StringVar(&ipF, "ip", "", "(Required) Address of host.")
	hostAddCmd.Flags().StringVar(&resPoolF, "resource-pool", "", "(Required) Resource pool ID.")
	hostAddCmd.Flags().StringVar(&credName, "credentials", "", "(Required if using existing one.) Credentials ID.")
	hostAddCmd.Flags().StringVar(&deplPolicyF, "deployment-policy", "", "Deployment policy ID.")
	hostAddCmd.Flags().BoolVar(&autoAccept, "accept", false, "Auto accept if certificate is not trusted.")
	hostAddCmd.Flags().StringSliceVar(&custProps, "cp", []string{}, custPropsDesc)
	HostsRootCmd.AddCommand(hostAddCmd)
}

var hostDisableCmd = &cobra.Command{
	Use:   "disable [HOST-ADDRESS]",
	Short: "Disable host with address provided.",
	Long:  "Disable host with address provided.",

	Run: func(cmd *cobra.Command, args []string) {
		var (
			hostAddress string
			ok          bool
		)
		if hostAddress, ok = ValidateArgsCount(args); !ok {
			fmt.Println("Enter host address.")
			return
		}
		newID, err := hosts.DisableHost(hostAddress)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Host disabled " + newID)
		}
	},
}

func initHostDisable() {
	HostsRootCmd.AddCommand(hostDisableCmd)
}

var hostEnableCmd = &cobra.Command{
	Use:   "enable [HOST-ADDRESS]",
	Short: "Enable host with address provided.",
	Long:  "Enable host with address provided.",

	Run: func(cmd *cobra.Command, args []string) {
		var (
			hostAddress string
			ok          bool
		)
		if hostAddress, ok = ValidateArgsCount(args); !ok {
			fmt.Println("Enter host address.")
			return
		}
		newID, err := hosts.EnableHost(hostAddress)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Host enabled: " + newID)
		}
	},
}

func initHostEnable() {
	HostsRootCmd.AddCommand(hostEnableCmd)
}

var hostListCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists existing hosts.",
	Long:  "Lists existing hosts.",

	Run: func(cmd *cobra.Command, args []string) {
		hl := &hosts.HostsList{}
		count := hl.FetchHosts(queryF)
		if count < 1 {
			fmt.Println("n/a")
			return
		}
		hl.Print()
	},
}

func initHostList() {
	hostListCmd.Flags().StringVarP(&queryF, "query", "q", "", "Add query.")
	hostListCmd.SetUsageTemplate(help.DefaultUsageListTemplate)
	HostsRootCmd.AddCommand(hostListCmd)
}

var hostRemoveCmd = &cobra.Command{
	Use:   "rm [HOST-ADDRESS]",
	Short: "Remove existing host.",
	Long:  "Remove existing host.",

	Run: func(cmd *cobra.Command, args []string) {
		var (
			address string
			ok      bool
		)
		if address, ok = ValidateArgsCount(args); !ok {
			fmt.Println("Enter host address.")
			return
		}
		fmt.Printf("Are you sure you want to remove %s? (y/n)\n", address)
		answer := functions.PromptAgreement()
		if answer == "n" || answer == "no" {
			fmt.Println("Remove command aborted!")
			return
		}

		newID, err := hosts.RemoveHost(address, asyncTask)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Host removed: " + newID)
		}
	},
}

func initHostRemove() {
	hostRemoveCmd.Flags().BoolVar(&asyncTask, "async", false, asyncDesc)
	HostsRootCmd.AddCommand(hostRemoveCmd)
}

var hostUpdateCmd = &cobra.Command{
	Use:   "update [ADDRESS]",
	Short: "Edit existing hosts.",

	Run: func(cmd *cobra.Command, args []string) {
		var (
			address string
			ok      bool
		)
		if address, ok = ValidateArgsCount(args); !ok {
			fmt.Println("Enter host address.")
			return
		}
		newID, err := hosts.EditHost(address, hostName, resPoolF, deplPolicyF, credName, autoAccept)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Host updated: " + newID)
		}
	},
}

func initHostUpdate() {
	hostUpdateCmd.Flags().StringVar(&hostName, "name", "", "New host name.")
	hostUpdateCmd.Flags().StringVar(&credName, "credentials", "", "New credentials ID.")
	hostUpdateCmd.Flags().StringVar(&resPoolF, "resource-pool", "", "New resource pool ID.")
	hostUpdateCmd.Flags().StringVar(&deplPolicyF, "deployment-policy", "", "New deployment policy ID.")
	hostUpdateCmd.Flags().BoolVar(&autoAccept, "accept", false, "Auto accept if certificate is not trusted.")
	HostsRootCmd.AddCommand(hostUpdateCmd)
}
