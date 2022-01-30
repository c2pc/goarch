package cmd

import (
	"context"
	"fmt"
	"github.com/chincharovpc/goarch/config"
	"github.com/chincharovpc/goarch/repository"
	"github.com/chincharovpc/goarch/service"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"log"
	"strings"
	"time"
)

var name string
var displayName string
var description string

var roleCmd = &cobra.Command{
	Use:   "role",
	Short: "Create Role command",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := config.GetConnection()
		if err != nil {
			log.Fatal(err)
			return
		}
		defer db.Close()

		logger, err := zap.NewDevelopment()
		if err != nil {
			log.Fatal(err)
			return
		}

		roleRepository := repository.NewRoleRepository(db, logger)
		roleService := service.NewRoleService(roleRepository)

		const timeout = 5 * time.Second

		ctx, shutdown := context.WithTimeout(context.Background(), timeout)
		defer shutdown()

		if displayName == "" {
			displayName = strings.Title(strings.ToLower(name))
		}
		role, err := roleService.Create(ctx, service.CreateRoleInput{
			Name:        name,
			DisplayName: displayName,
			Description: description,
		})
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Printf("%s successfully created", role.DisplayName)
	},
}

func init() {
	appCmd.AddCommand(roleCmd)

	pf := roleCmd.PersistentFlags()

	pf.StringVarP(&name, "name", "n", "", "name")
	pf.StringVar(&displayName, "display_name", "", "display name")
	pf.StringVar(&description, "description", "", "description")

	_ = cobra.MarkFlagRequired(pf, "name")

}
