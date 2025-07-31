package app

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aunz/api-mobile-dashboard-golang/internal/firestore"
	"github.com/aunz/api-mobile-dashboard-golang/internal/handlers"
	"github.com/aunz/api-mobile-dashboard-golang/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type App struct {
	Engine *gin.Engine
	Client *firestore.ClientWrapper
	ctx    context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) Init() {
	godotenv.Load()

	a.ctx = context.Background()

	sa := option.WithCredentialsFile("serviceAccountKey.json")
	fbApp, err := firebase.NewApp(a.ctx, nil, sa)
	if err != nil {
		log.Fatalln("Firebase init error:", err)
	}

	client, err := fbApp.Firestore(a.ctx)
	if err != nil {
		log.Fatalln("Firestore client error:", err)
	}

	a.Client = firestore.NewClientWrapper(client, a.ctx)

	a.Engine = gin.Default()

	// Register handlers with dependencies
	h := handlers.NewBuildInfoHandler(a.Client)
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
