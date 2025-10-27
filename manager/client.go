package manager

type client struct {
	channel  string
	dataChan chan interface{}
}

func (c *client) Send(data interface{}, closeChan chan *client) {
	defer func() {
		if x := recover(); x != nil {
			closeChan <- c
		}
	}()

	c.dataChan <- data
}
