package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"github.com/armandwipangestu/golang-simple-restful-api/internal/db"
	"github.com/armandwipangestu/golang-simple-restful-api/internal/handlers"
	"github.com/armandwipangestu/golang-simple-restful-api/internal/logger"
	"github.com/armandwipangestu/golang-simple-restful-api/internal/models"
	"github.com/armandwipangestu/golang-simple-restful-api/internal/repository"
	"github.com/armandwipangestu/golang-simple-restful-api/internal/seed"
)

func main() {
	// Logger
	logz, _ := logger.New()
	defer logz.Sync()

	if err := godotenv.Load(); err != nil {
		logz.Warn("No .env file found, reading env from system")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	ssl := os.Getenv("DB_SSLMODE")

	dsn := db.DSNFromEnv(host, port, user, pass, name, ssl)
	gdb, err := db.Connect(dsn)
	if err != nil {
		logz.Sugar().Fatalf("Cannot connect to db: %v", err)
	}

	logz.Info("DB connected")

	// Migration + seeder
	logz.Info("Running migration")
	if err := gdb.AutoMigrate(&models.Address{}, &models.Student{}); err != nil {
		logz.Fatal("Migration failed", zap.Error(err))
	}

	logz.Info("Running seeder")
	if err := seed.MigrateAndSeed(gdb); err != nil {
		logz.Fatal("Seeder failed", zap.Error(err))
	}

	// Repo + handler
	studentRepo := repository.NewStudentRepo(gdb)
	studentHandler := handlers.NewStudentHandler(studentRepo)

	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/students", studentHandler.List)
		api.POST("/students", studentHandler.Create)
		api.GET("/students/:id", studentHandler.GetById)
	}

	portApp := os.Getenv("APP_PORT")
	if portApp == "" {
		portApp = "3090"
	}
	logz.Sugar().Infof("Application running on http://localhost:%v", portApp)
	r.Run(":" + portApp)
}