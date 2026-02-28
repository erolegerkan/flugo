/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/erolegerkan/flugo/common"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "To configure flugo CLI",
	Long: `
		Configure flugo CLI tool with config command. 
		To get existing configurations run this command with "--list", "-l" flag.
		To insert new configuration run this command with "--insert", "-i" flag.
		To get configuration schema, run this command with "--schema", "-s" flag.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		
		isVerboseModeActive, _  := cmd.Flags().GetBool("verbose")
		isListModeActive,_ := cmd.Flags().GetBool("list")
		isInsertModeActive,_ := cmd.Flags().GetBool("insert")
		isSchemaModeActive,_ := cmd.Flags().GetBool("schema")
		
		isDefaultModeActive := !isListModeActive && !isInsertModeActive && !isSchemaModeActive
		
		if isVerboseModeActive {
			common.VerbosePrint("Currently executing : "+ cmd.CommandPath())
		}
		
		if isDefaultModeActive {
			common.VerbosePrint("Default mode active. By default \"config\" command returns configurations.")
			common.VerbosePrint("To different options for \"config\" command run the command with \"help\" flag")
		}
		
		
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	
	configCmd.Flags().BoolP("list","l", false,"Returns the configuration lists.")
	configCmd.Flags().BoolP("insert","i", false,"Inserts new configuration file.")
	configCmd.Flags().BoolP("schema","s", false,"Returns the configuration schema.")
	
}
