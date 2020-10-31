package store

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestStore(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3308)/tgbot")
	if err != nil {
		t.Error(err)
	}
	s := NewStore(db)
	s.AddUser("@Heman01", 119596917, 777777777)
	t.Log("TEST ADD ok")
	u, err := s.GetUserByID(777777777)
	if u == nil {
		fmt.Println("User not active")
	}
	fmt.Println(u)
}
func TestStore_AddSubscription(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3308)/tgbot")
	if err != nil {
		t.Error(err)
	}
	s := NewStore(db)
	//s.AddSubscription(777777777, 2)
	errr := s.AddSubscription(119596916, 1)
	fmt.Println(errr)
}
func TestStore_DeleteUser(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3308)/tgbot")
	if err != nil {
		t.Error(err)
	}
	s := NewStore(db)
	err = s.DeleteUser(119596916)
	if err != nil {
		t.Error(err)
	}

}
func TestMySQL_GetUserList(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3308)/tgbot")
	if err != nil {
		t.Error(err)
	}
	s := NewStore(db)
	u := s.GetUserList()
	fmt.Println(u)
}
func TestMySQL_AddUserTest(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3308)/tgbot")
	if err != nil {
		t.Error(err)
	}
	s := NewStore(db)
	s.AddUserTest(119596916)
}
func TestMySQL_EndSub(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3308)/tgbot")
	if err != nil {
		t.Error(err)
	}
	s := NewStore(db)
	s.EndSub(119596916)
}
func TestMySQL_GetTestUsers(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3308)/tgbot")
	if err != nil {
		t.Error(err)
	}
	s := NewStore(db)
	sl := s.GetTestUsers()
	fmt.Println(sl)
}
func TestMySQL_GetTestUser(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3308)/tgbot")
	if err != nil {
		t.Error(err)
	}
	s := NewStore(db)
	sl := s.GetTestUser(119596916)
	fmt.Println(sl)
	fmt.Println(sl.TestEnd[0] == []byte{0}[0])
}
