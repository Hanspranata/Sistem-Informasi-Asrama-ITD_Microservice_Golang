package models

type Student struct {
    ID       uint64 `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
    RoomID   uint64 `json:"room_id"`
}
