package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/benbjohnson/tiny-ego"
	"github.com/benbjohnson/tiny-ego/http"
	"github.com/benbjohnson/tiny-ego/inmem"
)

func main() {
	if err := run(); err == flag.ErrHelp {
		os.Exit(1)
	} else if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	// Parse command line arguments.
	fs := flag.NewFlagSet("tiny-ego", flag.ContinueOnError)
	addr := fs.String("addr", ":8000", "bind address")
	if err := fs.Parse(os.Args[1:]); err != nil {
		return err
	}

	// Instantiate database & services.
	db := inmem.NewDB()
	widgetService := inmem.NewWidgetService(db)

	// Instantiate HTTP handler.
	handler := http.NewHandler()
	handler.WidgetService = widgetService

	// Start server.
	server := http.NewServer()
	server.Addr = *addr
	if err := server.Open(); err != nil {
		return err
	}
	defer server.Close()

	// Shutdown on SIGINT (CTRL-C).
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Fprintln(m.Stderr, "received interrupt, shutting down...")

	return nil
}
