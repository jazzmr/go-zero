package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tbUserInfoFieldNames          = builder.RawFieldNames(&TbUserInfo{})
	tbUserInfoRows                = strings.Join(tbUserInfoFieldNames, ",")
	tbUserInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(tbUserInfoFieldNames, "`create_time`", "`update_time`"), ",")
	tbUserInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(tbUserInfoFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheTbUserInfoIdPrefix = "cache:tbUserInfo:id:"
)

type (
	TbUserInfoModel interface {
		Insert(ctx context.Context, data *TbUserInfo) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*TbUserInfo, error)
		Update(ctx context.Context, data *TbUserInfo) error
		Delete(ctx context.Context, id string) error
	}

	defaultTbUserInfoModel struct {
		sqlc.CachedConn
		table string
	}

	TbUserInfo struct {
		Id         string         `db:"id"`         // 主键
		NickName   sql.NullString `db:"nickName"`   // nickName
		Sex        int64          `db:"sex"`        // 性别:0-未知 1-男性 2-女性
		Country    sql.NullString `db:"country"`    // 国家
		Province   sql.NullString `db:"province"`   // 省份
		City       sql.NullString `db:"city"`       // 城市
		Avatar     sql.NullString `db:"avatar"`     // 头像
		CreateBy   string         `db:"createBy"`   // 创建人
		CreateTime time.Time      `db:"createTime"` // 创建时间
		UpdateBy   sql.NullString `db:"updateBy"`   // 修改人
		UpdateTime sql.NullTime   `db:"updateTime"` // 修改时间
	}
)

func NewTbUserInfoModel(conn sqlx.SqlConn, c cache.CacheConf) TbUserInfoModel {
	return &defaultTbUserInfoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`tb_user_info`",
	}
}

func (m *defaultTbUserInfoModel) Insert(ctx context.Context, data *TbUserInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, tbUserInfoRowsExpectAutoSet)
	ret, err := m.ExecNoCacheCtx(ctx, query, data.Id, data.NickName, data.Sex, data.Country, data.Province, data.City, data.Avatar, data.CreateBy, data.UpdateBy)

	return ret, err
}

func (m *defaultTbUserInfoModel) FindOne(ctx context.Context, id string) (*TbUserInfo, error) {
	tbUserInfoIdKey := fmt.Sprintf("%s%v", cacheTbUserInfoIdPrefix, id)
	var resp TbUserInfo
	err := m.QueryRowCtx(ctx, &resp, tbUserInfoIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tbUserInfoRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *defaultTbUserInfoModel) Update(ctx context.Context, data *TbUserInfo) error {
	tbUserInfoIdKey := fmt.Sprintf("%s%v", cacheTbUserInfoIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tbUserInfoRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.NickName, data.Sex, data.Country, data.Province, data.City, data.Avatar, data.CreateBy, data.UpdateBy, data.Id)
	}, tbUserInfoIdKey)
	return err
}

func (m *defaultTbUserInfoModel) Delete(ctx context.Context, id string) error {
	tbUserInfoIdKey := fmt.Sprintf("%s%v", cacheTbUserInfoIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, tbUserInfoIdKey)
	return err
}

func (m *defaultTbUserInfoModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTbUserInfoIdPrefix, primary)
}

func (m *defaultTbUserInfoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tbUserInfoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}
