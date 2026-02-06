/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename the APK file with a specified name",
	Long: `
		Rename the APK file with name speficied by quatitation marks.
		If you cannot add a flag command won't do anything.

		You can also add dynamic version values to new name. To specified version naming
		use "version" command to more about it run the "version" command.

		Example :
		flugo rename "<specified_name>"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rename called")
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)

	currentPath, err := os.Getwd()

	if err != nil {
		fmt.Println("❗️Error occurred when getting current path")
		return
	}

	fmt.Println(currentPath)

	lsCmdStruct := exec.Command("ls")
	lsCmdOutput, lsErr := lsCmdStruct.Output()

	if lsErr != nil {
		fmt.Println("❗️Error occurred when running ls command.")
		return
	}

	currentPathItems := string(lsCmdOutput)

	isLibFolderExists := strings.Contains(currentPathItems, "lib")
	isAndroidFolderExists := strings.Contains(currentPathItems, "android")
	isBuildFolderExists := strings.Contains(currentPathItems, "build")

	if isAndroidFolderExists && isLibFolderExists && isBuildFolderExists {
		fmt.Println("This is a flutter project")
		exec.Command("cd build/app/outputs/flutter-apk/app-release.apk")
		lsBuildCmdOutput, err := exec.Command("ls").Output()

		if err != nil {
			fmt.Println("❗️Error occurred when running ls command in build folder.")
			return
		}

		isApkExists := strings.Contains(string(lsBuildCmdOutput), "app-release.apk")

		if !isApkExists {
			return
		}

		renamingOperationOutput, _ := exec.Command("mv", "app-release.apk","<new_name_should_be_here>").Output()
		isRenamingSucceed := strings.Contains(string(renamingOperationOutput), "No such file or directory")

		if !isRenamingSucceed {
			fmt.Println("❗️Error occurred during renaming APK file.")
			return
		}

		fmt.Println("✅ Renaming completed successfully")

	} else {
		fmt.Println("Ooops this isn't a flutter project")
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// renameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// renameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
