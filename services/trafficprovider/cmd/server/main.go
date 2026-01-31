package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"sudoaptgetmach.me/trafficprovider/internal/adapter/vatsim"
	"sudoaptgetmach.me/trafficprovider/internal/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env")
		return
	}

	service.InitCache()

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:5672/", os.Getenv("RABBITMQ_USER"), os.Getenv("RABBITMQ_PASS"), os.Getenv("RABBITMQ_HOST")))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			failOnError(err, "Something went wrong")
		}
	}(conn)

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			failOnError(err, "Something went wrong")
		}
	}(ch)

	err = ch.ExchangeDeclare(
		"flight_events",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		failOnError(err, "Failed to declare an exchange")
	}

	ticker := time.NewTicker(15 * time.Second)

	for range ticker.C {
		flights := vatsim.FetchData()

		if flights == nil {
			log.Println("ERR: Nenhum dado recebido (ou erro na API).")
			continue
		}

		log.Printf("INFO: Processando %d voos...", len(flights))

		for _, flight := range flights {
			body, err := json.Marshal(flight)
			if err != nil {
				log.Printf("Erro ao serializar voo %s: %v", flight.Aircraft.Callsign, err)
				continue
			}

			err = ch.Publish(
				"flight_events",
				"flight.vatsim.update",
				false,
				false,
				amqp.Publishing{
					ContentType: "application/json",
					Body:        body,
				},
			)

			if err != nil {
				log.Printf("Erro ao publicar %s: %v", flight.Aircraft.Callsign, err)
			}
		}
		log.Printf("INFO: Ciclo finalizado.")
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
