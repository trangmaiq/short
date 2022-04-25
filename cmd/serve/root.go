package serve

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/trangmaiq/short/registry"
)

func NewServeCmd() *cobra.Command {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Run the Short server",
		Run: func(cmd *cobra.Command, args []string) {
			println("serving...")

			routes := gin.Default()

			_, err := registry.New(routes)
			if err != nil {
				log.Println("unable to create new registry")
				os.Exit(1)
			}

			err = routes.Run(":9090")
			if err != nil {
				log.Println("unable to run `short` service")
				os.Exit(1)
			}
		},
	}

	return serveCmd
}
