package server

import (
	"time"

	"github.com/jdvjdv82/wabbit"
)

type (
	// Delivery is an interface to delivered messages
	Delivery struct {
		data          []byte
		headers       wabbit.Option
		tag           uint64
		consumerTag   string
		originalRoute string
		messageId     string
		channel       *Channel
		contentType   string
	}
)

func NewDelivery(ch *Channel, data []byte, tag uint64, messageId string, hdrs wabbit.Option, contentType string) *Delivery {
	return &Delivery{
		data:        data,
		headers:     hdrs,
		channel:     ch,
		tag:         tag,
		messageId:   messageId,
		contentType: contentType,
	}
}

func (d *Delivery) Ack(multiple bool) error {
	return d.channel.Ack(d.tag, multiple)
}

func (d *Delivery) Nack(multiple, requeue bool) error {
	return d.channel.Nack(d.tag, multiple, requeue)
}

func (d *Delivery) Reject(requeue bool) error {
	return d.channel.Nack(d.tag, false, requeue)
}

func (d *Delivery) Body() []byte {
	return d.data
}

func (d *Delivery) Headers() wabbit.Option {
	return d.headers
}

func (d *Delivery) DeliveryTag() uint64 {
	return d.tag
}

func (d *Delivery) ConsumerTag() string {
	return d.consumerTag
}

func (d *Delivery) MessageId() string {
	return d.messageId
}

func (d *Delivery) Timestamp() time.Time {
	return time.Now()
}

func (d *Delivery) ContentType() string {
	return d.contentType
}
