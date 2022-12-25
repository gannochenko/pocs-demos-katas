package main

import (
	"os"
	"strconv"

	"copydata/internal/copier"
	"copydata/internal/util"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "copy-data",
		Usage: "Just another Golang kata",
		Commands: []*cli.Command{
			{
				Name:    "copy",
				Aliases: []string{"c"},
				Usage:   "Copy data from one db instance to another",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "file",
						Usage: "File containing a list of IDs",
					},
					&cli.StringFlag{
						Name:  "src-port",
						Value: "5432",
						Usage: "Source instance port",
					},
					&cli.StringFlag{
						Name:  "src-user",
						Value: "test",
						Usage: "Source database user",
					},
					&cli.StringFlag{
						Name:  "dst-port",
						Value: "5432",
						Usage: "Destination instance port",
					},
					&cli.StringFlag{
						Name:  "dst-user",
						Value: "test",
						Usage: "Destination database user",
					},
				},
				Action: func(cCtx *cli.Context) error {
					filePath := cCtx.String("file")

					srcPort := cCtx.String("src-port")
					srcUser := cCtx.String("src-user")
					srcPassword := util.PromptUser("Source database password")

					dstPort := cCtx.String("dst-port")
					dstUser := cCtx.String("dst-user")
					dstPassword := util.PromptUser("Source database password")

					srcPortNumber, err := strconv.Atoi(srcPort)
					if err != nil {
						panic(err)
					}
					dstPortNumber, err := strconv.Atoi(dstPort)
					if err != nil {
						panic(err)
					}

					idList, err := util.ReadCSV(filePath)

					copierInstance := copier.New(&copier.Options{
						SrcDBPort:     int32(srcPortNumber),
						SrcDBUser:     srcUser,
						SrcDBPassword: srcPassword,
						DstDBPort:     int32(dstPortNumber),
						DstDBUser:     dstUser,
						DstDBPassword: dstPassword,
					})

					err = copierInstance.CopyElements(idList[0])
					if err != nil {
						panic(err)
					}

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
