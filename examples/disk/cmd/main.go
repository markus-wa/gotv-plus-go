package main

import (
	"flag"
	"log"
	"net"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/FlowingSPDG/gotv-plus-go/examples/disk"
	"github.com/FlowingSPDG/gotv-plus-go/gotv"
)

func main() {
	addr := flag.String("addr", "localhost:8080", "address to listen on")
	auth := flag.String("auth", "SuperSecureStringDoNotShare", "authentication token")
	dataDir := flag.String("data", "gotv_plus_binary", "data directory")

	flag.Parse()

	if *auth == "" {
		log.Fatal("auth argument is required")
	}

	m := disk.NewDiskGOTV(*auth, *dataDir)
	app := fiber.New()
	g := app.Group("/gotv") // /gotv
	g.Use(logger.New())
	gotv.SetupStoreHandlersFiber(m, g)
	gotv.SetupBroadcasterHandlersFiber(m, g)

	p := net.JoinHostPort("0.0.0.0", "8080")

	// Start server
	log.Println("Start listening on:", p)
	if err := app.Listen(p); err != nil {
		panic(err)
	}
}
