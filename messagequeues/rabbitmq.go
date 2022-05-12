package messagequeues

import (
	"fmt"

	"github.com/streadway/amqp"

	"github.com/salemzii/go-watchdog/service"
)

func MakeRabbitmqQuery(mq *MsgQueue) service.ServiceCheck {

	uri := mq.RabbitmqDsn()

	if mq.Uri_Only() {
		uri = mq.UriOnly
	}

	conn, err := amqp.Dial(uri)

	if err != nil {
		return *service.HandleError("rabbitmq", err)
	}

	fmt.Println(conn.Properties)

	return *service.HandleSuccess("rabbitmq", nil)
}
