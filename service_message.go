package transport

import (
	"encoding/json"
	"github.com/ronappleton/gk-message-transport/storage"
)

// ServiceMessage message passing orders to a service
// Used by websocket server to pass incoming client messages to a service
// and services sending orders to other services.
type ServiceMessage struct {
	SessionId     string                      `json:"session_id"` // Client Identifier
	UserId        string                      `json:"user_id"`    // User Id used for storage
	Topic         string                      `json:"topic"`      // The topic channel to send the command on
	Command       string                      `json:"command"`    // The command the client wants executed
	Arguments     map[string]interface{}      `json:"arguments"`  // The arguments the service may need for executing the command
	ArgumentStore storage.KeyValueStoreAccess `json:"-"`          // KeyValueStoreAccess version of Results for typed usage
	Results       map[string]interface{}      `json:"results"`    // The results of any computation
	ResultStore   storage.KeyValueStoreAccess `json:"-"`          // KeyValueStoreAccess version of Results for typed usage
	ByteAble
}

// ClientMessage Is a message that can be returned to the client, it is noticeable
// missing the Arguments, UserID and SessionId, Arguments as not needed, and UserId
// and SessionId as the client should never know either.
type ClientMessage struct {
	Topic   string                 `json:"topic"`   // The topic channel to send the command on
	Command string                 `json:"command"` // The command the client wants executed
	Results map[string]interface{} `json:"results"` // The results of any computation
	ByteAble
}

// ByteAble Is to ensure that messages have a uniform method of conversion for sending.
type ByteAble interface {
	ToBytes() []byte
}

// NewServiceMessage Is for getting a new instance of a ServiceMessage
func NewServiceMessage() ServiceMessage {
	return ServiceMessage{}
}

// NewClientMessage Is for getting a new instance of a ClientMessage
func NewClientMessage() ClientMessage {
	return ClientMessage{}
}

// ToBytes Is used to convert a ServiceMessage for sending to a service.
func (message *ServiceMessage) ToBytes() []byte {
	out, _ := json.Marshal(message)

	return out
}

// ToBytes Is used to convert a ClientMessage for sending to the client.
func (message *ClientMessage) ToBytes() []byte {
	out, _ := json.Marshal(message)

	return out
}

// ToClientMessage Is used to convert a service result to a reply.
// We remove UserId and SessionId from the reply as the client
// should never know about them.
func (message *ServiceMessage) ToClientMessage() ClientMessage {
	clientMessage := NewClientMessage()
	clientMessage.Topic = message.Topic
	clientMessage.Command = message.Command
	clientMessage.Results = message.Results

	return clientMessage
}

// BytesToServiceMessage Is used to convert an incoming client message in bytes
// to a ServiceMessage or from a service message in bytes.
func BytesToServiceMessage(bytes []byte) ServiceMessage {
	serviceMessage := NewServiceMessage()

	_ = json.Unmarshal(bytes, &serviceMessage)

	serviceMessage.ArgumentStore = storage.New()
	serviceMessage.ArgumentStore.Populate(serviceMessage.Arguments)
	serviceMessage.ResultStore = storage.New()
	serviceMessage.ResultStore.Populate(serviceMessage.Results)

	return serviceMessage
}
