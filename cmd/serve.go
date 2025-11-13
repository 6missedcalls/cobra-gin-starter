package cmd

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/6missedcalls/cobra-gin-starter/api"
	"github.com/6missedcalls/cobra-gin-starter/web/views"
	"github.com/a-h/templ"
	"github.com/spf13/cobra"
)

var port string

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the application locally",
	Long:  `Serves a web application built in HTMX for an easy-to-use local interface.`,

	Run: func(cmd *cobra.Command, args []string) {
		mux := http.NewServeMux()

		// API
		apiEngine := api.SetupRouter()
		mux.Handle("/api/", apiEngine)

		// Static
		staticDir := http.FileServer(http.Dir(filepath.Join("web", "static")))
		mux.Handle("/static/", http.StripPrefix("/static/", staticDir))

		// Pages
		mux.Handle("/", templ.Handler(views.Hello()))

		err := http.ListenAndServe(":"+port, mux)
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	},
}

func init() {
	serveCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to run the server on")

	rootCmd.AddCommand(serveCmd)
}
