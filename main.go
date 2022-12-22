package main

import (
	"context"
	"os"

	"github.com/Mayurhole95/Brandscope-go/app"
	"github.com/Mayurhole95/Brandscope-go/config"
	csv_upload "github.com/Mayurhole95/Brandscope-go/csv_upload"
	csv_validate "github.com/Mayurhole95/Brandscope-go/csv_validate"
	"github.com/Mayurhole95/Brandscope-go/db"
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
		//If again want to convert to api , we need the below part to start server

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
				// Run Validate
				csv := csv_validate.NewService(dbstorer, logger)

				csv.Validate(context.TODO(), c.Args().Get(0))
				return nil
			},
		},
		{
			Name:  "upload",
			Usage: "run upload code",
			Action: func(c *cli.Context) error {
				// Run Validate as well as upload
				csv := csv_validate.NewService(dbstorer, logger)

				csv.Validate(context.TODO(), c.Args().Get(0))
				csv1 := csv_upload.NewService(dbstorer, logger)

				csv1.Upload(context.TODO(), c.Args().Get(0))
				return nil
			},
		},
	}
	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}
}
