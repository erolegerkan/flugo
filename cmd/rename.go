/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

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
		
	},
}

func init() {
	var currentDirectoryItems []string
	
	rootCmd.AddCommand(renameCmd)

	currentPath, err := os.Getwd()

	if err != nil {
		fmt.Println("❗️Error occurred when getting current path")
		return
	}

	fmt.Println("Verbose Mode : " + currentPath)

	currentDirectoryRead, currentDirectoryReadError := os.ReadDir(currentPath)

	if currentDirectoryReadError != nil {
		fmt.Println("❗️Error occurred when getting current directory.")
		return
	}
	
	for _, value  := range currentDirectoryRead {
		currentDirectoryItems = append(currentDirectoryItems, value.Name())
	}
	
}
	
	// isLibFolderExists := strings.Contains(currentPathItems, "lib")
	// isAndroidFolderExists := strings.Contains(currentPathItems, "android")
	// isBuildFolderExists := strings.Contains(currentPathItems, "build")
	
	// // Checking for Flutter project 
	// if isAndroidFolderExists && isLibFolderExists && isBuildFolderExists {
	// 	fmt.Println("Verbose Mode : This is a flutter project")
		
	// 	// Run "ls" command inside of the Flutter app
	// 	lsCmd := exec.Command("ls")
	// 	lsCmd.Dir = "build/app/outputs/flutter-apk/"
	// 	lsCmdOutput, cdCmdError  := lsCmd.CombinedOutput()
		
	// 	if cdCmdError != nil {
	// 		fmt.Println("❗️Error occurred when running ls command in build folder insid Flutter app.")
	// 		return
	// 	}

	// 	isApkExists := strings.Contains(string(lsCmdOutput), "app-release.apk")

	// 	if !isApkExists {
	// 		fmt.Println("⚠️\tRelease APK file does not exist!\nFirst build the release apk with \"flutter build apk\"")
	// 		return
	// 	}

	// 	renamingOperation := exec.Command("mv", "app-release.apk","new_apk_name.apk")
	// 	renamingOperation.Dir = "build/app/outputs/flutter-apk/"
	// 	renamingOperationOutput, renamingError := renamingOperation.CombinedOutput()
		
	// 	if renamingError != nil {
	// 		fmt.Println("Verbose Mode : " + string(renamingOperationOutput))
	// 		fmt.Println("❗️Error occurred during renaming APK file.")
	// 		return
	// 	}
		

	// 	fmt.Println("✅ Renaming completed successfully")

	// } else {
	// 	fmt.Println("Ooops this isn't a flutter project")
	// }

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// renameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// renameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

