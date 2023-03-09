package main

import (
	"database/sql"
	"fmt"

	"github.com/deividroger/ms-wallet/src/internal/database"
	event "github.com/deividroger/ms-wallet/src/internal/events"
	createaccount "github.com/deividroger/ms-wallet/src/internal/usecase/create_account"
	createclient "github.com/deividroger/ms-wallet/src/internal/usecase/create_client"
	createtransaction "github.com/deividroger/ms-wallet/src/internal/usecase/create_transaction"
	"github.com/deividroger/ms-wallet/src/internal/web"
	"github.com/deividroger/ms-wallet/src/internal/web/webserver"
	"github.com/deividroger/ms-wallet/src/pkg/events"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("root:root@tcp(localhost:3306)/wallet?parseTime=true"))

	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEventHandler := event.NewTransactionCreated()
	//eventDispatcher.Register("TransactionCreated", handler)

	clientDb := database.NewClientDb(db)
	accountDb := database.NewAccountDb(db)
	transactionDb := database.NewTransactionDB(db)

	createClientUseCase := createclient.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := createtransaction.NewCreateTransactionUseCase(transactionDb, accountDb, eventDispatcher, transactionCreatedEventHandler)

	webserver := webserver.NewWebServer(":3000")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	webserver.Start()

}
