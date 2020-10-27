package store

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestStore(t *testing.T) {
	db, err:=sql.Open("mysql", "root:root@tcp(localhost:3308)/tgbot")
	if err != nil {
		t.Error(err)
	}
	s:= NewStore(db)
	s.AddUser("@Heman01", 119596917, 777777777)
	t.Log("TEST ADD ok")
	u, err := s.GetUserByID(777777777)
	if u == nil {
		fmt.Println("User not active")
	}
	fmt.Println(u)
}
func TestStore_AddSubscription(t *testing.T) {
	db, err:=sql.Open("mysql", "root:root@tcp(localhost:3308)/tgbot")
	if err != nil {
		t.Error(err)
	}
	s:= NewStore(db)
	//s.AddSubscription(777777777, 2)
	errr := s.AddSubscription(119596916, 1)
	fmt.Println(errr)
}
func TestStore_DeleteUser(t *testing.T) {
	db, err:=sql.Open("mysql", "root:root@tcp(localhost:3308)/tgbot")
	if err != nil {
		t.Error(err)
	}
	s:= NewStore(db)
	 err =s.DeleteUser(119596916)
	 if err != nil {
		 t.Error(err)
	 }

}