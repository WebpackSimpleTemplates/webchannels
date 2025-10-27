package manager

func NewCore() *Core {
	addChan := make(chan *client, 100)
	removeChan := make(chan *client, 100)
	sendChan := make(chan *message, 100)

	clients := make([]*client, 0)

	go func() {
		for {
			select {
			case client := <-addChan:
				clients = append(clients, client)
			case client := <-removeChan:

				var id int

				for i, item := range clients {
					if item == client {
						id = i
						break
					}
				}

				clients = append(clients[:id], clients[id+1:]...)
			case message := <-sendChan:
				for _, client := range clients {
					if message.channel == client.channel {
						client.Send(message.data, removeChan)
					}
				}
			}
		}
	}()

	return &Core{
		addChan,
		removeChan,
		sendChan,
	}
}
