package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Repinoid/diploma56/internal/handlers"
	"github.com/Repinoid/diploma56/internal/models"
	"github.com/Repinoid/diploma56/internal/securitate"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var host = "localhost:8081"

// пока без горутин select for update и проч

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic("cannot initialize zap")
	}
	defer logger.Sync()
	models.Sugar = *logger.Sugar()

	if err := initEnvs(); err != nil {
		panic(err)
	}

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	var err error
	ctx := context.Background()

	securitate.DataBase, err = securitate.ConnectToDB(ctx)

	if err != nil {
		fmt.Printf("database connection error  %v", err)
		return err
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/user/register", handlers.RegisterUser).Methods("POST")
	router.HandleFunc("/api/user/login", handlers.LoginUser).Methods("POST")
	router.HandleFunc("/api/user/balance/withdraw", handlers.Withdraw).Methods("POST")

	router.HandleFunc("/api/user/orders", handlers.PutOrder).Methods("POST")
	router.HandleFunc("/api/user/orders", handlers.GetOrders).Methods("GET")
	router.HandleFunc("/api/user/withdrawals", handlers.GetWithDrawals).Methods("GET")
	router.HandleFunc("/api/user/balance", handlers.GetBalance).Methods("GET")

	return http.ListenAndServe(host, router)
}
