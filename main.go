package main

import (
	"context"
	"os"

	"github.com/Mayurhole95/Brandscope-go/app"
	"github.com/Mayurhole95/Brandscope-go/config"
<<<<<<< HEAD
	csv_validate "github.com/Mayurhole95/Brandscope-go/csv_validate"
	"github.com/Mayurhole95/Brandscope-go/db"
=======
	"github.com/Mayurhole95/Brandscope-go/db"
	"github.com/Mayurhole95/Brandscope-go/server"
>>>>>>> 8b4ccc9f32baaca90bbc165e1936b396c96f17da
	"github.com/urfave/cli"
)

func main() {
	config.Load()
	app.Init()
	defer app.Close()

	dbstorer := db.NewStorer(app.GetDB())
	logger := app.GetLogger()

	cliApp := cli.NewApp()
	cliApp.Name = "Golang App"
	cliApp.Version = "1.0.0"
	cliApp.Commands = []cli.Command{
		// {
		// 	Name:  "start",
		// 	Usage: "start server",
		// 	Action: func(c *cli.Context) error {
		// 		server.StartAPIServer()
		// 		return nil
		// 	},
		// },
		{
			Name:  "validate",
			Usage: "run validations code",
			Action: func(c *cli.Context) error {
				// fmt.Println("Hii")
				csv := csv_validate.NewService(dbstorer, logger)

				csv.Validate(context.TODO(), c.Args().Get(0))
				return nil
			},
		},

		{
			Name:  "create_migration",
			Usage: "create migration file",
			Action: func(c *cli.Context) error {
				return db.CreateMigrationFile(c.Args().Get(0))
			},
		},
		{
			Name:  "migrate",
			Usage: "run db migrations",
			Action: func(c *cli.Context) error {
				err := db.RunMigrations()
				return err
			},
		},
		{
			Name:  "rollback",
			Usage: "rollback migrations",
			Action: func(c *cli.Context) error {
				return db.RollbackMigrations(c.Args().Get(0))
			},
		},
	}
	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}
}
