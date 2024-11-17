package ws

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

// architecture --> Hub (consist rooms (go routines)) --> Room(consist Id , room name ,different clients ) --> Client
type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

// initialize hub constructor
func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}

func (h *Hub) Run() {
	for {
		select {
		//check whether the room id which we got from client is present or not
		case cl := <-h.Register:
			if _,ok := h.Rooms[cl.RoomID]; ok{
				r := h.Rooms[cl.RoomID]

			// add this client with this room id
			if _,ok := r.Clients[cl.ID]; !ok{
				r.Clients[cl.ID] = cl
			}

			}
		case cl := <-h.Unregister:
			// check roomid and clientid for rooms and delete it
			if _, ok := h.Rooms[cl.RoomID]; ok {
				if _, ok := h.Rooms[cl.RoomID].Clients[cl.ID]; ok {

					//boardcase message that user left the chat
					if len(h.Rooms[cl.RoomID].Clients) != 0 {
						h.Broadcast <- &Message{
							Content:  "user left the chat",
							RoomID:   cl.RoomID,
							Username: cl.Username,
						}
					}

					delete(h.Rooms[cl.RoomID].Clients, cl.ID)
					close(cl.Message)
				}

			}
		case m := <-h.Broadcast:
			// check the room existence and send all the clients message
			if _, ok := h.Rooms[m.RoomID]; ok {
				for _, cl := range h.Rooms[m.RoomID].Clients {
					cl.Message <- m
				}
			}
		}

	}
}
