package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
)

var (
	host   string
	passwd string
	user   string
	port   string
	WG     sync.WaitGroup
)

type UserHistory struct {
	Id         int64  `json:"f_id" db:"f_id"`
	UserId     int64  `json:"f_user_id" db:"f_user_id"`
	CorpId     int64  `json:"f_corp_id" db:"f_corp_id"`
	DeptId     int64  `json:"f_dept_id" db:"f_dept_id"`
	UserName   string `json:"f_user_name" db:"f_user_name"`
	CorpName   string `json:"f_corp_name" db:"f_corp_name"`
	DeptName   string `json:"f_dept_name" db:"f_dept_name"`
	Stat       int64  `json:"f_stat" db:"f_stat"`
	Account    string `json:"f_account" db:"f_account"`
	Qq         string `json:"f_qq" db:"f_qq"`
	Email      string `json:"f_email" db:"f_email"`
	Title      string `json:"f_title" db:"f_title"`
	OptUser    int64  `json:"f_opt_user" db:"f_opt_user"`
	CreateTime string `json:"f_create_time" db:"f_create_time"`
	ModifyTime string `json:"f_modify_time" db:"f_modify_time"`
	Face       string `json:"f_face" db:"f_face"`
}


func main() {
	flag.StringVar(&host, "h", "192.168.1.150", "数据库连接")
	flag.StringVar(&passwd, "p", "ecuser", "密码")
	flag.StringVar(&user, "u", "ecuser", "用户")
	flag.StringVar(&port, "port", "3306", "用户")
	flag.Parse()
	mysql := fmt.Sprintf("%s:%s@(%s:%s)/d_ec_user?charset=utf8mb4", user, passwd, host, port)
	fmt.Println("连接数据库:", mysql)
	db, err := sql.Open("mysql", mysql)
	db.SetMaxOpenConns(10)
	if err != nil {
		log.Fatalf("连接数据库失败%v", err)
		return
	}
	detailMap := make(map[int64]DeptInfo)
	detailMap, err = getAllUserInfo(db)
	if err != nil {
		log.Fatalf("查询所有用户失败%v", err)
	}

	for _, v := range detailMap {
		temp := UserHistory{}
		if _, ok := detailMap[v.Uid]; ok {
			temp.UserId = detailMap[v.Uid].Uid
			temp.DeptId = detailMap[v.Uid].DeptId
		}
		fmt.Println(temp)
		insertAllUserInDb(temp, db)
	}
	fmt.Println("刷总数", len(detailMap))
}

type DeptInfo struct {
	Uid    int64 `json:"f_uid"`
	DeptId int64 `json:"f_dept_id"`
}

func getAllUserInfo(db *sql.DB) (map[int64]DeptInfo, error) {
	sqlQuery := "select a.f_uid ,a.f_dept_id from t_user a ,t_user_history b  where a.f_uid = b.f_user_id and a.f_dept_id != b.f_dept_id"
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]DeptInfo)
	for rows.Next() {
		var id int64
		var deptId int64

		if err := rows.Scan(&id, &deptId); err != nil {
			log.Fatal(err)
		}
		t := DeptInfo{
			Uid:    id,
			DeptId: deptId,
		}

		result[id] = t
	}
	return result, nil
}

func insertAllUserInDb(v UserHistory, db *sql.DB) {
	sqlQuery := "update t_user_history set f_dept_id = ? where f_user_id= ?"
	_, err := db.Exec(sqlQuery, v.DeptId, v.UserId)
	if err != nil {
		fmt.Println("插入错误", v.UserId)
		log.Fatalf("insert error%v\n", err)
	}
}
