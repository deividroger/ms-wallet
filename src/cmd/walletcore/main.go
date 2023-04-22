package main

import (
	"context"
	"database/sql"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/deividroger/ms-wallet/src/internal/database"
	event "github.com/deividroger/ms-wallet/src/internal/events"
	"github.com/deividroger/ms-wallet/src/internal/events/handler"
	createaccount "github.com/deividroger/ms-wallet/src/internal/usecase/create_account"
	createclient "github.com/deividroger/ms-wallet/src/internal/usecase/create_client"
	createtransaction "github.com/deividroger/ms-wallet/src/internal/usecase/create_transaction"
	"github.com/deividroger/ms-wallet/src/internal/web"
	"github.com/deividroger/ms-wallet/src/internal/web/webserver"
	"github.com/deividroger/ms-wallet/src/pkg/events"
	"github.com/deividroger/ms-wallet/src/pkg/kafka"
	"github.com/deividroger/ms-wallet/src/pkg/uow"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("root:root@tcp(mysql:3306)/wallet?parseTime=true"))

	if err != nil {
		panic(err)
	}
	defer db.Close()

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}

	kafkaProducer := kafka.NewKafkaProducer(&configMap)

	eventDispatcher := events.NewEventDispatcher()

	eventDispatcher.Register("TransactionCreated", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))
	eventDispatcher.Register("BalanceUpdated", handler.NewBalanceUpdatedKafkaHandler(kafkaProducer))

	transactionCreatedEventHandler := event.NewTransactionCreated()
	balanceUpdatedEventHandler := event.NewBalanceUpdated()

	clientDb := database.NewClientDb(db)
	accountDb := database.NewAccountDb(db)

	ctx := context.Background()

	uow := uow.NewUow(ctx, db)

	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDb(db)
	})

	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})

	createClientUseCase := createclient.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := createtransaction.NewCreateTransactionUseCase(uow, eventDispatcher, transactionCreatedEventHandler, balanceUpdatedEventHandler)

	webserver := webserver.NewWebServer(":8080")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	fmt.Println("Server is running")
	webserver.Start()

}
