package app

import (
	"fmt"
	"log"
	"os"

	"github.com/aunz/api-mobile-dashboard-golang/internal/database"
	"github.com/aunz/api-mobile-dashboard-golang/internal/handlers"
	"github.com/aunz/api-mobile-dashboard-golang/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type App struct {
	Engine *gin.Engine
	DB     *database.PostgresDB
}

func NewApp() *App {
	return &App{}
}

func (a *App) Init() {
	godotenv.Load()

	// Initialize PostgreSQL database
	db, err := database.NewPostgresDB()
	if err != nil {
		log.Printf("PostgreSQL initialization failed: %v", err)
		log.Println("\nTroubleshooting steps:")
		log.Println("1. Ensure PostgreSQL is running")
		log.Println("2. Check database connection parameters in environment variables:")
		log.Println("   - DB_HOST (default: localhost)")
		log.Println("   - DB_PORT (default: 5432)")
		log.Println("   - DB_USER (default: postgres)")
		log.Println("   - DB_PASSWORD (default: password)")
		log.Println("   - DB_NAME (default: mobile_dashboard)")
		log.Println("   - DB_SSLMODE (default: disable)")
		log.Println("3. Verify database exists and user has proper permissions")
		log.Fatalln("Database initialization failed:", err)
	}

	a.DB = db
	a.Engine = gin.Default()

	// Register handlers with dependencies
	h := handlers.NewBuildInfoHandler(a.DB)
	routes.InitializeRoutes(a.Engine, h)

	fmt.Println("Server will start at port " + a.GetPort())
}

func (a *App) GetPort() string {
	port := os.Getenv("MyPort")
	if port == "" {
		port = "3000"
	}
	return ":" + port
}

func (a *App) Run() {
	a.Engine.Run(a.GetPort())
}
