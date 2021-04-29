// dbTest project main.go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"log"
)

//数据库配置
const (
	userName = "root"
	password = "root"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "demo"
)

var DB *sql.DB
var (
	ErrNoRows = errors.New("错误类型ErrNoRows")
	ErrConnD  = errors.New("数据库链接失败")
)

func InitDB() error {
	path := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", userName, password, ip, port, dbName)
	DB, _ = sql.Open("mysql", path)
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		return errors.Wrapf(ErrConnD, "连接数据库失败~")
	}
	return nil
}

func main() {
	err := InitDB()
	if err != nil {
		fmt.Printf("链接数据库日志:%+v\n", err)
		panic("数据库链接失败~")
	}
	user, err := SelectUserById("WQ")
	if err != nil {
		fmt.Printf("查询数据日志:%+v\n", err)
	}
	fmt.Println(user)
}

type User struct {
	Id   int
	Name string
	Age  int
	Sex  int
}

func SelectUserById(name string) (users []User, err error) {
	var (
		user User
		s    sql.NullInt32
	)

	err = DB.QueryRow("SELECT age FROM user where name=?", "Wq").Scan(&s) //即将得到的name值转换成s.String类型并存储到&s中
	log.Println("测试零值：", err)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {//数据可能为空
			return nil, errors.Wrapf(ErrNoRows, "查询字符串为空~")
		}
		return nil, errors.Wrapf(ErrNoRows, "其他错误类型1~")

	}

	err = DB.QueryRow("select * from user where name = ?", "WQ").Scan(&user)
	fmt.Println(errors.Is(err, sql.ErrNoRows))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			//return users, errors.WithMessage(ErrNoRows, "改行数据不存在~")
			return users, errors.Wrapf(ErrNoRows, "改行数据不存在~")
		}
		return nil, errors.Wrapf(ErrNoRows, "其他错误类型~")
	}
	users = append(users, user)
	return users, err
}
