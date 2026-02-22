/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/erolegerkan/flugo/common"
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
	var isVerboseMode bool = true
	
	var currentDirectoryItems []string
	var isLibFolderExist bool
	var isAndroidFolderExist bool
	var isBuildFolderExist bool
	
	var apkPathDirectoryItems []string
	
	rootCmd.AddCommand(renameCmd)

	currentPath, err := os.Getwd()

	if err != nil {
		fmt.Println("❗️Error occurred when getting current path")
		return
	}

	common.VerbosePrint("Current Path : " + currentPath)

	currentDirectoryRead, currentDirectoryReadError := os.ReadDir(currentPath)

	if currentDirectoryReadError != nil {
		fmt.Println("❗️Error occurred when getting current directory.")
		return
	}
	
	for _, value  := range currentDirectoryRead {
		currentDirectoryItems = append(currentDirectoryItems, value.Name())
		
		if strings.Contains(value.Name(), "lib") && value.IsDir() {
			isLibFolderExist = true
		}
		
		if strings.Contains(value.Name(), "build") && value.IsDir() {
			isBuildFolderExist = true
		}
	
		if strings.Contains(value.Name(), "android") && value.IsDir() {
			isAndroidFolderExist = true
		}
	}
	
	if isVerboseMode {
		common.VerbosePrint("Current Directory Items : \n" + strings.Join(currentDirectoryItems," "))
		
		if isLibFolderExist {
			common.VerbosePrint("lib folder found in the project directory.")
		} else {
			common.VerbosePrint("lib folder did NOT found in the project directory.")
		}
		
		if isAndroidFolderExist {
			common.VerbosePrint("android folder found in the project directory.")
		} else {
			common.VerbosePrint("android folder did NOT found in the project directory.")
		}
		
		if isBuildFolderExist {
			common.VerbosePrint("build folder found in the project directory.")
		} else {
			common.VerbosePrint("build folder did NOT found in the project directory.")
		}
	}
	
	if !(isLibFolderExist && isBuildFolderExist && isAndroidFolderExist) {
		common.ErrorPrint("Ooops, this is not a Flutter Mobile project for Android platform.")
		return
	}
	
	apkPath := filepath.Join("build", "app", "outputs", "flutter-apk")
	
	apkPathDirectoryRead, apkPathDirectoryReadError := os.ReadDir(apkPath)
	
	if apkPathDirectoryReadError != nil {
		common.ErrorPrint("Error occurred when getting APK directory.")
		return
	}
	
	common.VerbosePrint("APK Path read : " + apkPath)
	
	var releaseApkCount int = 0
	var releaseApkName string
	for _, value  := range apkPathDirectoryRead {
		apkPathDirectoryItems = append(apkPathDirectoryItems, value.Name())
		if value.Name() == "app-release.apk" {
			releaseApkName = value.Name()
			releaseApkCount++
		}
	}
	
	common.VerbosePrint("APK Path Directory Items : \n" + strings.Join(apkPathDirectoryItems, " "))
	
	if releaseApkCount != 1 {
		common.ErrorPrint("Release APK file did not found in the directory with " + releaseApkName + "name.")
		return
	}
	
	oldApkPath := filepath.Join(apkPath,releaseApkName)
	newApkPath := filepath.Join(apkPath,"new_apk_name.apk")
	apkRenamingError := os.Rename(oldApkPath, newApkPath)
	
	if apkRenamingError != nil {
		common.ErrorPrint("Error occurred when APK file renaming.")
		return 
	}
	
	common.SuccessPrint("APK file renamed successfully.\nNew name for the APK : <new_apk_name.apk>")
	
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

