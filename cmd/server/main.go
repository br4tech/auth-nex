package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open(viper.GetString("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados")
	}

	router := chi.NewRouter()

	authHandler := InitializeAuthHandler(db)

	router.Route("/auth", func(r chi.Router) {
		r.Get("/authenticate", authHandler.AuthenticateHandler)
		r.Get("/register", authHandler.RegisterHandler)
	})

	router.Get("/check", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("AuthNex is running"))
	})

	fmt.Println("AuthNext listening on port" + viper.GetString("WEB_SERVER_PORT") + "!")

	http.ListenAndServe(":"+viper.GetString("WEB_SERVER_PORT"), router)
}
