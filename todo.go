package todoapp

type TodoList struct {
	Id          int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoItem struct {
	Id          int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        string `json:"done"`
}

// таблицы связей списки -> пользователи
type UsersList struct {
	Id     int
	UserId int
	ListId int
}

// таблицы связей списки -> записи
type ItemsList struct {
	Id     int
	ItemId int
	ListId int
}
