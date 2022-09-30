package roommodel

type Rooms struct {
	Id            int    `json:"id"`
	RoomName      string `json:"room_name"`
	MaximumPerson int    `json:"maximum_person"`
}

func (Rooms) TableName() string {
	return "rooms"
}
