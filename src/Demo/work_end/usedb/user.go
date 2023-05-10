package usedb

import (
	"Demo/work_end/model"
	"Demo/work_end/utils"
)

//根据用户名和密码在数据库中插询一条记录      ,验证用户名和密码
func CheckUserNameAndPassWord(username string, password string) (*model.User, error) {
	//写sql语句
	sqlStr := "select id,username,password from users where username = ? and password =?"
	row := utils.Db.QueryRow(sqlStr, username, password)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password)
	return user, nil
}

//验证用户名
func CheckUserName(username string) (*model.User, error) {
	//写sql语句
	sqlStr := "select id,username,password from users where username = ?"
	row := utils.Db.QueryRow(sqlStr, username)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password)
	return user, nil
}

//save user   向数据库插入用户信息
func SaveUser(username string, password string) error {
	sqlstr := "insert into users(username,password) values(?,?)"
	_, err := utils.Db.Exec(sqlstr, username, password)
	if err != nil {
		return err
	}
	return nil
}
