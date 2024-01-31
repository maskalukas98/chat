package value_object

type UserId struct {
	value int64
}

func NewUserId(id int64) *UserId {
	return &UserId{value: id}
}

func (r *UserId) GetId() int64 {
	return r.value
}
