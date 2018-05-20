package main

import (
	"log"
	"net/http"
	"time"

	"./history"
	"./mobile"
	"./overview"
	"./send"
	"./store"
	"./users"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	//USERS API ENDPOINT
	r.HandleFunc("/api/login", users.Login).Methods("POST")
	r.HandleFunc("/api/isLoged", users.Islogged).Methods("POST")
	r.HandleFunc("/api/register", users.Register).Methods("POST")
	r.HandleFunc("/api/verifyAccount", users.VerifyAccount).Methods("POST")
	r.HandleFunc("/api/updateProfile", users.UpdateProfile).Methods("POST")
	r.HandleFunc("/api/getAccountProfile", users.GetAccountProfile).Methods("POST")
	r.HandleFunc("/api/recoverPassword", users.RecoverPassword).Methods("POST")
	r.HandleFunc("/api/newPassword", users.NewPassword).Methods("POST")
	//OVERVIEW API ENDPOINTS
	r.HandleFunc("/api/getUserData", overview.GetUserGenericData).Methods("POST")
	r.HandleFunc("/api/getCurrencyPrice", overview.GetCoinPrice).Methods("POST")
	r.HandleFunc("/api/getUserTrans", overview.GetTransactions).Methods("POST")
	//STORE API ENDPOINTS
	r.HandleFunc("/api/getOrders", store.GetOrdersEndpoint).Methods("POST")
	r.HandleFunc("/api/getUserOrders", store.GetOrdersUserEndpoint).Methods("POST")
	r.HandleFunc("/api/saveOrder", store.CreateOrder).Methods("POST")
	r.HandleFunc("/api/getOrderBalance", store.GetBalanceOnOrder).Methods("POST")
	r.HandleFunc("/api/payOrder", store.BuyOrder).Methods("POST")
	//SEND API ENDPOINTS
	r.HandleFunc("/api/getAddresses", send.GetUserAddresses).Methods("POST")
	r.HandleFunc("/api/getNewAddress", send.GetNewAddress).Methods("POST")
	r.HandleFunc("/api/sendLocal", send.SendLocal).Methods("POST")
	r.HandleFunc("/api/sendExternal", send.SendExternal).Methods("POST")
	r.HandleFunc("/api/getFee", send.GetNetworkFee).Methods("POST")
	//HISTORY API ENDPOINTS
	r.HandleFunc("/api/loginHistory", history.GetLoginHistory).Methods("POST")
	r.HandleFunc("/api/actionHistory", history.GetActionsHistory).Methods("POST")
	r.HandleFunc("/api/orderHistory", history.GetOrderHistory).Methods("POST")
	//MOBILE API ENDPOINT
	r.HandleFunc("/apiv2/getLogin", mobile.LoginByMobileID).Methods("POST")
	r.HandleFunc("/apiv2/getMobileBalances", mobile.MobileUserBalances).Methods("POST")
	server := &http.Server{
		Addr:           ":8080",
		Handler:        cors.Default().Handler(r),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Listening http://localhost:8080 ...")
	log.Fatal(server.ListenAndServe())
}
