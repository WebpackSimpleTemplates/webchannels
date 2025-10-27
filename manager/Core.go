package manager

type message struct {
	channel string
	data    interface{}
}

type Core struct {
	addChan    chan *client
	removeChan chan *client
	sendChan   chan *message
}

func (core *Core) Send(channel string, data interface{}) {
	core.sendChan <- &message{channel, data}
}

func (core *Core) Add(channel string, chanSize int) chan interface{} {
	dataChan := make(chan interface{}, chanSize)

	core.addChan <- &client{channel, dataChan}

	return dataChan
}
