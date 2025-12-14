/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

const AppVersion = "0.1.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "mspawn",
	Short:   "Utility to manage long running processes",
	Long:    `mspawn is a utility program to orchestrate long running processes. It's main usecase is running and exiting long running processes.`,
	Version: AppVersion,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	runCmd.Flags().StringArrayVar(&commands, "cmd", []string{}, "Commands to run in paralell")

	err := runCmd.MarkFlagRequired("cmd")
	if err != nil {
		panic(err)
	}
	rootCmd.AddCommand(runCmd)
}
