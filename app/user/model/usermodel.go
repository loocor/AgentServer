package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	userFieldNames          = builderx.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	UserModel interface {
		Insert(data User) (sql.Result, error)
		FindOne(id int64) (*User, error)
		FindOneByPhone(phone string) (*User, error)
		FindOneByIdNumber(idNumber string) (*User, error)
		Update(data User) error
		Delete(id int64) error
	}

	defaultUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	User struct {
		Id         int64        `db:"id"`
		Kind       int64        `db:"kind"`       // 账户类型 0渠道经纪人 1渠道部门主管 2渠道机构主管 3甲方项目主管 4甲方集团主管
		Role       string       `db:"role"`       // 账户角色
		From       string       `db:"from"`       // 注册来源
		Phone      string       `db:"phone"`      // 账户号码
		Name       string       `db:"name"`       // 账户姓名
		Nickname   string       `db:"nickname"`   // 账户昵称
		Gender     string       `db:"gender"`     // 账户性别
		Password   string       `db:"password"`   // 账户密码
		IdNumber   string       `db:"id_number"`  // 身份证号
		OpenId     string       `db:"open_id"`    // 微信ID
		Organize   string       `db:"organize"`   // 机构
		Department string       `db:"department"` // 部门
		JobTitle   string       `db:"job_title"`  // 职务
		Avatar     string       `db:"avatar"`     // 头像
		Address    string       `db:"address"`    // 地址
		CreateTime time.Time    `db:"create_time"`
		UpdateTime time.Time    `db:"update_time"`
		DeleteTime sql.NullTime `db:"delete_time"`
	}
)

func NewUserModel(conn sqlx.SqlConn) UserModel {
	return &defaultUserModel{
		conn:  conn,
		table: "`user`",
	}
}

func (m *defaultUserModel) Insert(data User) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Kind, data.Role, data.From, data.Phone, data.Name, data.Nickname, data.Gender, data.Password, data.IdNumber, data.OpenId, data.Organize, data.Department, data.JobTitle, data.Avatar, data.Address, data.DeleteTime)
	return ret, err
}

func (m *defaultUserModel) FindOne(id int64) (*User, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	var resp User
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByPhone(phone string) (*User, error) {
	var resp User
	query := fmt.Sprintf("select %s from %s where `phone` = ? limit 1", userRows, m.table)
	err := m.conn.QueryRow(&resp, query, phone)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByIdNumber(idNumber string) (*User, error) {
	var resp User
	query := fmt.Sprintf("select %s from %s where `id_number` = ? limit 1", userRows, m.table)
	err := m.conn.QueryRow(&resp, query, idNumber)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Update(data User) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Kind, data.Role, data.From, data.Phone, data.Name, data.Nickname, data.Gender, data.Password, data.IdNumber, data.OpenId, data.Organize, data.Department, data.JobTitle, data.Avatar, data.Address, data.DeleteTime, data.Id)
	return err
}

func (m *defaultUserModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
