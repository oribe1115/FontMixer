/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/oribe1115/fontmixer/auth"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
	"google.golang.org/api/slides/v1"
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

	srv, err := slides.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return fmt.Errorf("failed to get service: %w", err)
	}

	// テスト用のコード

	// Prints the number of slides and elements in a sample presentation:
	// https://docs.google.com/presentation/d/1EAYk18WDjIG-zp_0vLm3CsfQh_i8eXc67Jo2O9C6Vuc/edit
	presentationId := "1EAYk18WDjIG-zp_0vLm3CsfQh_i8eXc67Jo2O9C6Vuc"
	presentation, err := srv.Presentations.Get(presentationId).Do()
	if err != nil {
		return fmt.Errorf("unable to retrieve data from presentation: %v", err)
	}

	fmt.Printf("The presentation contains %d slides:\n", len(presentation.Slides))
	for i, slide := range presentation.Slides {
		fmt.Printf("- Slide #%d contains %d elements.\n", (i + 1),
			len(slide.PageElements))
	}

	return nil
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
