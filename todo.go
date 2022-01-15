package todoapp

// списки
type TodoList struct {
	Id          int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// записи
type TodoItem struct {
	Id          int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        string `json:"done"`
}

// таблица связей пользователь -> списки
type UsersList struct {
	Id     int
	UserId int
	ListId int
}

// таблица связей списки -> записи
type ItemsList struct {
	Id     int
	ItemId int
	ListId int
}
