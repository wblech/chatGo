package consumer

import "chatGo/src/infrastructure/queue"

func RunAllConsumers(qBroker *queue.Broker) {
	go qBroker.ConsumeMessage(GetStockData, "bot-send")
	go qBroker.ConsumeMessage(SendShareInfo, "bot-receive")
}
