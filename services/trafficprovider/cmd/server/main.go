package main

import (
	json "encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"sudoaptgetmach.me/trafficprovider/internal/domain"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env")
		return
	}

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

	ticker := time.NewTicker(5 * time.Second)

	for range ticker.C {
		mockFlight := &domain.Flight{
			ScenarioId:      uuid.New(),
			Source:          domain.VatsimLive,
			Aircraft:        domain.Aircraft{},
			FlightPlan:      domain.FlightPlan{},
			EnvironmentMock: domain.EnvironmentMock{},
			ExpectedState:   domain.Clearance,
		}

		mockFlightJson, jsonErr := json.Marshal(mockFlight)

		if jsonErr != nil {
			log.Printf("Erro ao criar JSON: %v", err)
			continue
		}

		err = ch.Publish(
			"flight_events",
			"flight.vatsim.update",
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        mockFlightJson,
			},
		)

		if err != nil {
			log.Printf("Erro ao publicar: %s", err)
		} else {
			log.Printf(" [x] Sent flight update")
		}
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
