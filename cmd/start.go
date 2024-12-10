/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/franciscolkdo/family-feud/config"
	"github.com/franciscolkdo/family-feud/internal/game"
	"github.com/spf13/cobra"
)

var configPath string

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the game!",
	Long:  `Start the family-feud game.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.GetConfig(configPath)
		if err != nil {
			return fmt.Errorf("error starting game: %w", err)
		}
		_, err = tea.NewProgram(
			game.New(cfg), tea.WithMouseCellMotion()).Run()
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	startCmd.Flags().StringVarP(&configPath, "config", "c", "", "config file to use")
	rootCmd.AddCommand(startCmd)
}
