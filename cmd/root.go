/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/oribe1115/fontmixer/auth"
	"github.com/oribe1115/fontmixer/googleslides"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fontmixer",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	RunE: run,
}

func run(cmd *cobra.Command, args []string) error {
	fmt.Println("called")

	ctx := context.Background()
	client, err := auth.GetClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to get client: %w", err)
	}

	presentationID, err := getPresentationID()
	if err != nil {
		return fmt.Errorf("failed to get presentation ID: %w", err)
	}

	googlseSlides, err := googleslides.SetupGoogleSlides(ctx, client, presentationID)
	if err != nil {
		return fmt.Errorf("failed to setup Google Slides: %w", err)
	}

	return nil
}

func getPresentationID() (string, error) {
	validate := func(input string) error {
		if len(input) == 0 {
			return fmt.Errorf("presentation ID is required")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Presentation ID",
		Validate: validate,
	}

	presentationID, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return presentationID, nil
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fontmixer.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
