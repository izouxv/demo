package handler

type Handler interface {
	IntegrationHandler
	ReceivedChanEmail()    chan  PayloadEmail
}

type IntegrationHandler interface {
	Close() error
}

