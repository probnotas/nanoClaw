// NanoClaw - Ultra-lightweight personal AI agent
// Inspired by and based on nanobot: https://github.com/HKUDS/nanobot
// License: MIT
//
// Copyright (c) 2026 NanoClaw contributors

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/probnotas/nanoClaw/cmd/nanoclaw/internal"
	"github.com/probnotas/nanoClaw/cmd/nanoclaw/internal/agent"
	"github.com/probnotas/nanoClaw/cmd/nanoclaw/internal/auth"
	"github.com/probnotas/nanoClaw/cmd/nanoclaw/internal/cron"
	"github.com/probnotas/nanoClaw/cmd/nanoclaw/internal/gateway"
	"github.com/probnotas/nanoClaw/cmd/nanoclaw/internal/migrate"
	"github.com/probnotas/nanoClaw/cmd/nanoclaw/internal/onboard"
	"github.com/probnotas/nanoClaw/cmd/nanoclaw/internal/skills"
	"github.com/probnotas/nanoClaw/cmd/nanoclaw/internal/status"
	"github.com/probnotas/nanoClaw/cmd/nanoclaw/internal/version"
)

func NewNanoclawCommand() *cobra.Command {
	short := fmt.Sprintf("%s nanoclaw - Personal AI Assistant v%s\n\n", internal.Logo, internal.GetVersion())

	cmd := &cobra.Command{
		Use:     "nanoclaw",
		Short:   short,
		Example: "nanoclaw list",
	}

	cmd.AddCommand(
		onboard.NewOnboardCommand(),
		agent.NewAgentCommand(),
		auth.NewAuthCommand(),
		gateway.NewGatewayCommand(),
		status.NewStatusCommand(),
		cron.NewCronCommand(),
		migrate.NewMigrateCommand(),
		skills.NewSkillsCommand(),
		version.NewVersionCommand(),
	)

	return cmd
}

func main() {
	cmd := NewNanoclawCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
