package store

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"time"
)

type MySQL struct {
	DB *sql.DB
}
type users struct {
	UserID       int
	UserName     string
	Subscription string
	TimeSub      int
	IDchat       int
}
type tests struct {
	UserID  int
	TimeSub int
	TestEnd []byte
}
type userList struct {
	List []int
}

func (u *userList) AddId(user int) {
	u.List = append(u.List, user)
}
func NewUlist() userList {
	u := userList{}
	u.List = make([]int, 0)
	return u
}
func NewStore(db *sql.DB) MySQL {
	return MySQL{DB: db}
}
func (s *MySQL) AddUser(userName string, IDChat, UserID int) error {
	_, err := s.DB.Exec("INSERT INTO tgbot.users (user_id, user_name, subscription, id_chat) VALUES (?, ?, ?, ?)", UserID, userName, mysql.NullTime{}, IDChat)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (s *MySQL) GetUserByName(userName string) (*users, error) {
	row := s.DB.QueryRow("SELECT * from tgbot.users where tgbot.users.user_name = ?", userName)
	u := users{}

	err := row.Scan(&u.UserName, &u.Subscription, &u.IDchat)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &u, nil

}
func (s *MySQL) GetUserByID(UserID int) (*users, error) {
	row := s.DB.QueryRow("SELECT * from tgbot.users where tgbot.users.user_id = ?", UserID)
	u := users{}

	err := row.Scan(&u.UserID, &u.UserName, &u.IDchat, &u.Subscription, &u.TimeSub)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &u, nil

}
func (s *MySQL) AddSubscription(UserID int, timeM int) error {
	t := time.Now().String()
	t = t[:10]
	res, err := s.DB.Exec("UPDATE tgbot.users set subscription = ?,subscription_longtime = ?  where user_id = ?", t, timeM, UserID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	r, err := res.RowsAffected()
	if r == 0 {
		return fmt.Errorf("Error update row")
	}
	return nil
	//_, err =s.DB.Exec("UPDATE tgbot.users set subscription_longtime = ? where user_id = ?",  time, UserID)
	//if err != nil {
	//	fmt.Println(err)
	//}
}
func (s *MySQL) DeleteUser(UserID int) error {
	_, err := s.DB.Exec("DELETE FROM tgbot.users where user_id  = ?", UserID)
	return err
}
func (s *MySQL) GetUserList() *userList {
	u := NewUlist()
	rows, err := s.DB.Query("select user_id FROM tgbot.users")
	if err != nil {
		fmt.Println(err)
	}
	var id int
	for rows.Next() {
		rows.Scan(&id)
		u.AddId(id)

	}
	return &u
}
func (s *MySQL) GetTestUsers() []tests {
	var usrs []tests
	var t tests
	rows, err := s.DB.Query("select * FROM tgbot.tests where testEnd = 00")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		rows.Scan(&t.UserID, &t.TimeSub, &t.TestEnd)
		usrs = append(usrs, t)
	}
	return usrs
}
func (s *MySQL) GetTestUser(UserID int) tests {

	var t tests
	row := s.DB.QueryRow("select * FROM tgbot.tests where user_id = ?", UserID)

	row.Scan(&t.UserID, &t.TimeSub, &t.TestEnd)

	return t
}
func (s *MySQL) EndSub(userID int) {
	_, err := s.DB.Exec("UPDATE tgbot.tests set testEnd = ? where user_id = ?", true, userID)
	fmt.Println(err)
}
func (s *MySQL) AddUserTest(userID int) {
	t := time.Now().Hour()
	_, err := s.DB.Exec("INSERT INTO tgbot.tests (user_id, testNow, testEnd) VALUES (?,?,?)", userID, t, false)
	fmt.Println(err)
}
