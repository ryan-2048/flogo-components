# MQTT
This trigger allows you to listen to messages on MQTT.

## Installation

### Flogo CLI
```bash
flogo install github.com/ryan-2048/flogo-components/trigger/mqtt
```

## Configuration

### Settings:
| Name          | Type   | Description
| :---          | :---   | :---
| broker        | string | The broker URL - ***REQUIRED***
| id            | string | The id of client - ***REQUIRED***
| username      | string | The user's name
| password      | string | The user's password
| store         | string | The store for message persistence
| cleanSession  | bool   | Clean session flag
| keepAlive     | int    | Keep Alive time in seconds
| autoReconnect | bool   | Enable Auto-Reconnect
| sslConfig     | object | SSL configuration

 #### *sslConfig* Object:
 | Property      | Type   | Description
 |:---           | :---   | :---     
 | skipVerify    | bool   | Skip SSL validation, defaults to true
 | useSystemCert | bool   | Use the systems root certificate file, defaults to true
 | caFile        | string | The path to PEM encoded root certificates file
 | certFile      | string | The path to PEM encoded client certificate
 | keyFile       | string | The path to PEM encoded client key

 *Note: used if broker URI is ssl*

### Handler Settings
| Name       | Type   | Description
| :---       | :---   | :---
| topic      | string | The topic to listen on - ***REQUIRED***
| qos        | int    | The Quality of Service
| retain     | bool   | Retain Messages
| connectReplyTopic | string | The topic to connect reply
| connectReplyMessage | string | The data to connect reply  

#### Topics
MQTT wildcard syntax is supported. For example if the topic is '/x/+/y/#' then the `topicParams` `output` will be populated with the wildcard values. The first wildcard will be in `topicParams` with key '0' and the second with key '1'. Topic wildcards can also be given a name: '/x/+param1/y/#param2'. Then the names 'param1' and 'param2' can be used to access the wildcards in the `topicParams` `output`.

### Output:

| Name        | Type   | Description
| :---        | :---   | :---
| message     | string | The message received
| topic       | string | The MQTT topic
| topicParams | params | The topic parameters

### Reply:

| Name  | Type   | Description
| :---  | :---   | :---
| replyTopic  | string | The data recieved
| replyQos  | int | The data recieved
| replyRetain  | bool | The data recieved
| replyData  | object | The data recieved

## Example

```json
{
  "id": "mqtt-trigger",
  "name": "Mqtt Trigger",
  "ref": "github.com/ryan-2048/flogo-components/trigger/mqtt",
  "settings": {
      "broker" : "tcp://localhost:1883",
     	"id":"client-1"
  },
  "handlers": {
    "settings": {
    	"topic": "flogo",
    	"qos": "0"
    }
  }
}
```
