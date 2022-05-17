package messagequeues

import (
	"fmt"
	"strings"
	"time"

	"github.com/salemzii/go-watchdog/service"
)

type MsgQueue struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Addrs    string `json:"addrs"`
	Timeout  uint   `json:"timeout"`
	UriOnly  string `json:"uri_only"` // when db credentials are already preconfigured to a single address/URL
}

func (mq *MsgQueue) Uri_Only() bool {
	return mq.UriOnly != ""
}

func (mq *MsgQueue) RabbitmqDsn() string {

	return fmt.Sprintf("amqp://%s:%s@%s/%s", mq.Username, mq.Password, mq.Addrs, mq.Password)
}

func (mq *MsgQueue) AwsMqDsn() string {
	return fmt.Sprintf("amqp+ssl://%s", mq.Addrs)
}

//amqp+ssl://b-4de5a854-a3c3-43e3-9f3d-d8fe82060e84-1.mq.us-east-1.amazonaws.com:5671

func (mq *MsgQueue) GetMqDriver() service.ServiceCheck {

	switch strings.ToLower(mq.Type) {
	case "rabbitmq":
		return MakeRabbitmqQuery(mq)
	case "activemq":
		return MakeActivemqQuery(mq)
	case "awsmq":
		return MakeAwsMqQuery(mq)
	}

	return service.ServiceCheck{}
}

func (mq *MsgQueue) GetOrSetConnTimeout() time.Duration {
	if mq.Timeout != 0 {
		return time.Duration(mq.Timeout * uint(time.Second))
	}
	mq.Timeout = uint(10)
	return time.Duration(mq.Timeout * uint(time.Second))
}
