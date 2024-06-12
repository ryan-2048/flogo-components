package mqtt

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	Broker        string                 `md:"broker,required"` // The broker URL
	Id            string                 `md:"id,required"`     // The id of client
	Username      string                 `md:"username"`        // The user's name
	Password      string                 `md:"password"`        // The user's password
	Store         string                 `md:"store"`           // The store for message persistence
	CleanSession  bool                   `md:"cleanSession"`    // Clean session flag
	KeepAlive     int                    `md:"keepAlive"`       // Keep Alive time in seconds
	AutoReconnect bool                   `md:"autoReconnect"`   // Enable Auto-Reconnect
	SSLConfig     map[string]interface{} `md:"sslConfig"`       // SSL Configuration
}

type HandlerSettings struct {
	Topic               string `md:"topic,required"`      // The topic to listen on
	Qos                 int    `md:"qos"`                 // The Quality of Service
	Retain              bool   `md:"retain"`              // retain
	ConnectReplyTopic   string `md:"connectReplyTopic"`   // Connect Reply Topic
	ConnectReplyMessage string `md:"connectReplyMessage"` // Connect Reply Message
}

type Output struct {
	Message     string            `md:"message"`     // The message recieved
	Topic       string            `md:"topic"`       // The MQTT topic
	TopicParams map[string]string `md:"topicParams"` // The topic parameters
}

type Reply struct {
	replyTopic  string      `md:"reply topic,required"` // The topic to reply with
	replyQos    int         `md:"reply qos"`            // The qos to reply with
	replyRetain bool        `md:"reply retain"`         // The retain to reply with
	replyData   interface{} `md:"reply data"`           // The data to reply with
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"message":     o.Message,
		"topic":       o.Topic,
		"topicParams": o.TopicParams,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.Message, err = coerce.ToString(values["message"])
	if err != nil {
		return err
	}
	o.Topic, err = coerce.ToString(values["topic"])
	if err != nil {
		return err
	}
	o.TopicParams, err = coerce.ToParams(values["topicParams"])
	if err != nil {
		return err
	}

	return nil
}

func (r *Reply) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"replyTopic":  r.replyTopic,
		"replyQos":    r.replyQos,
		"replyRetain": r.replyRetain,
		"replyData":   r.replyData,
	}
}

func (r *Reply) FromMap(values map[string]interface{}) error {

	var err error
	r.replyTopic, err = coerce.ToString(values["replyTopic"])
	if err != nil {
		return err
	}
	r.replyQos, err = coerce.ToInt(values["replyQos"])
	if err != nil {
		return err
	}
	r.replyRetain, err = coerce.ToBool(values["replyRetain"])
	if err != nil {
		return err
	}
	r.replyData = values["replyData"]
	return nil
}
