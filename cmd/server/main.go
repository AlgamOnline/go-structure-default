package main

import (
	"database/sql"
	"flag"
	"fmt"
	"golang-default/config"
	"golang-default/controllers"
	"golang-default/middlewares"
	"golang-default/services"
	"golang-default/ws"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/gorilla/mux"
)

func main() {

	cfg := config.LoadConfig()

	// Buat DSN MySQL
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	// Cek koneksi
	if err := db.Ping(); err != nil {
		log.Fatalf("DB not reachable: %v", err)
	}

	// Ambil port dari command line, default 5050
	port := flag.Int("port", 5050, "Port untuk menjalankan server HTTP")
	flag.Parse()

	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, 0755)
	}

	//jwtSecret := "your-secret-key"

	// Init services
	//sessionService := services.NewSessionService(jwtSecret)
	sessionService := services.NewSessionService("SECRET_KEY")
	authService := services.NewAuthService(db, sessionService)
	uploadService := services.NewUploadService(uploadDir)
	gpsService := services.NewGPSService(db)
	userService := services.NewUserService(db)
	unitService := services.NewUnitService(db)

	r := mux.NewRouter()
	r.HandleFunc("/api/login", controllers.LoginHandler(authService)).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()

	api.Use(middlewares.SessionMiddleware(sessionService))

	api.HandleFunc("/upload", controllers.UploadImageHandler(uploadService)).Methods("POST")
	api.HandleFunc("/gps", controllers.GPSHandler(gpsService)).Methods("POST")

	api.HandleFunc("/users/{id}", controllers.GetUserHandler(userService)).Methods("GET")
	api.HandleFunc("/users", controllers.CreateUserHandler(userService)).Methods("POST")
	api.HandleFunc("/users/{id}", controllers.UpdateUserHandler(userService)).Methods("PUT")
	api.HandleFunc("/users/{id}", controllers.DeleteUserHandler(userService)).Methods("DELETE")

	api.HandleFunc("/unit/{id}", controllers.GetUnitHandler(unitService)).Methods("GET")
	api.HandleFunc("/unit", controllers.CreateUnitHandler(unitService)).Methods("POST")
	api.HandleFunc("/unit/{id}", controllers.UpdateUnitHandler(unitService)).Methods("PUT")
	api.HandleFunc("/unit/{id}", controllers.DeleteUnitHandler(unitService)).Methods("DELETE")

	http.HandleFunc("/ws", ws.HandleConnections)

	addr := ":" + strconv.Itoa(*port)
	log.Printf("Server running on %s\n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
