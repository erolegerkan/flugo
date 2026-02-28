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
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		newApkName := args[0]

		var currentDirectoryItems []string
		var isLibFolderExist bool
		var isAndroidFolderExist bool
		var isBuildFolderExist bool

		var apkPathDirectoryItems []string

		isVerboseModeActive, _ := cmd.Flags().GetBool("verbose")

		currentPath, err := os.Getwd()

		if err != nil {
			fmt.Println("❗️Error occurred when getting current path")
			return
		}

		if isVerboseModeActive {
			common.VerbosePrint("Currently executing : " + cmd.CommandPath())
			common.VerbosePrint("Current Path : " + currentPath)
		}

		currentDirectoryRead, currentDirectoryReadError := os.ReadDir(currentPath)

		if currentDirectoryReadError != nil {
			fmt.Println("❗️Error occurred when getting current directory.")
			return
		}

		for _, value := range currentDirectoryRead {
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

		if isVerboseModeActive {
			common.VerbosePrint("Current Directory Items : \n" + strings.Join(currentDirectoryItems, " "))

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
		for _, value := range apkPathDirectoryRead {
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

		oldApkPath := filepath.Join(apkPath, releaseApkName)
		newApkPath := filepath.Join(apkPath, newApkName+".apk")
		apkRenamingError := os.Rename(oldApkPath, newApkPath)

		if apkRenamingError != nil {
			common.ErrorPrint("Error occurred when APK file renaming.")
			return
		}

		common.SuccessPrint("APK file renamed successfully.\nNew name for the APK : " + newApkName + ".apk")

	},
}

func init() {
	rootCmd.AddCommand(renameCmd)
}
