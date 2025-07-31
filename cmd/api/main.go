package main

import "github.com/aunz/api-mobile-dashboard-golang/internal/app"

func main() {
    route := app.NewApp()
    route.Init()
    route.Run()
}