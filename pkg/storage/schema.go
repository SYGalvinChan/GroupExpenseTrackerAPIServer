package storage

type User struct {
	UserID         int64 `gorm:"column:user_id;primary_key"`
	TelegramUserID int64 `gorm:"column:telegram_user_id"`
}

type Group struct {
	GroupID        int64 `gorm:"column:group_id;primary_key"`
	TelegramChatID int64 `gorm:"column:telegram_chat_id"`
}

type UsersInGroups struct {
	GroupID int64 `gorm:"column:group_id;primary_key"`
	UserID  int64 `gorm:"column:user_id"`
}

type Expense struct {
	ExpenseID   int64  `gorm:"column:expense_id;primary_key"`
	GroupID     int64  `gorm:"column:group_id"`
	ExpenseName string `gorm:"column:expense_name"`
	Amount      string `gorm:"column:amount"`
}

type PaidBy struct {
	ExpenseID int64  `gorm:"column:expense_id;primary_key"`
	UserID    int64  `gorm:"column:user_id"`
	Amount    string `gorm:"column:amount"`
}

func (m *PaidBy) TableName() string {
	return "paid_by"
}

type SplitBetween struct {
	ExpenseID int64  `gorm:"column:expense_id;primary_key"`
	UserID    int64  `gorm:"column:user_id"`
	Amount    string `gorm:"column:amount"`
}

func (m *SplitBetween) TableName() string {
	return "split_between"
}
