package dao

import (
	"Demo/goweb/model"
	"Demo/goweb/utils"
	"fmt"
)

/*
	Query a record from the database according to the user name and password.
	Check whether the user name and password are correct when you log in

    'username':用户名
    password:密码
	*model.User:用户信息
	error:错误信息
*/
func CheckLogin(username string, password string) (*model.User, error) {
	//根据用户名和密码从数据库中查询一条记录，登陆时校验用户名和密码是否正确
	//写sql语句
	sqlStr := "select user_id,nickname,username,password,email from user where username = ? and password = ?"
	//执行
	row := utils.DB.QueryRow(sqlStr, username, password)
	user := &model.User{}
	err := row.Scan(&user.Userid, &user.Nickname, &user.Username, &user.Password, &user.Email)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}

/*
   Query a record from the database based on the user name and password,
   and check whether the user name exists during registration
*/
func CheckUsername(username string) (*model.User, error) {
	//根据用户名和密码从数据库中查询一条记录,注册时校验用户名是否存在
	//写sql语句
	sqlStr := "select user_id,username,password,email from users where username = ?"
	//执行
	row := utils.DB.QueryRow(sqlStr, username)
	user := &model.User{}
	err := row.Scan(&user.Userid, &user.Nickname, &user.Username, &user.Password, &user.Email)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, err
}

/*
	Insert user information into the database and save the user information at registration time
*/
func SaveUser(username string, password string, email string) error {
	//向数据库中插入用户信息，注册时保存用户信息
	//写sql语句
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	//执行
	_, err := utils.DB.Exec(sqlStr, username, password, email)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
