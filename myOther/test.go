package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
	"time"
)

const URLHOST = "39.97.171.148" // TODO "localhost"

const (
	USERNAME = "root"
	PASSWORD = "201955"
	NETWORK  = "tcp"
	SERVER   = URLHOST
	PORT     = 3306
	DATABASE = "water_flow"
)

func main() {
	http.HandleFunc("/", host)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Println("ListenAndServe Error", err)
		return
	}
}

func connectDB() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Open MySQL Error", err)
		return nil
	}
	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(100)                  //设置最大连接数
	DB.SetMaxIdleConns(16)                   //设置闲置连接数
	return DB
}

func insertDB(DB *sql.DB, name string, score int) {
	result, err := DB.Exec("insert into rank(score,name) value (?,?)", score, name)
	if err != nil {
		log.Println("Insert Error", err)
		return
	}
	lastInsertID, err := result.LastInsertId() //插入数据的主键id
	if err != nil {
		log.Println("Get lastInsertID failed,err:%v", err)
		return
	}
	log.Println("LastInsertID:", lastInsertID)
}

func host(writer http.ResponseWriter, request *http.Request) {
	db := connectDB()
	request.ParseForm()
	log.Println(request.Form)
	score, err := strconv.Atoi(request.Form.Get("score"))
	if err != nil {
		log.Println("Score Can't To Int Error", err)
		return
	}
	insertDB(db, request.Form.Get("name"), score)
	//writer.Write([]byte("hello world"))
}
