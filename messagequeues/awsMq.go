package messagequeues

import (
	"fmt"

	"github.com/streadway/amqp"

	"github.com/salemzii/go-watchdog/service"
)

func MakeAwsMqQuery(mq *MsgQueue) service.ServiceCheck {

	uri := mq.AwsMqDsn()

	if mq.Uri_Only() {
		uri = mq.UriOnly
	}

	conn, err := amqp.Dial(uri)

	if err != nil {
		return *service.HandleError("awsmq", err)
	}

	fmt.Println(conn.Properties)

	return *service.HandleSuccess("awsmq", nil)

}
