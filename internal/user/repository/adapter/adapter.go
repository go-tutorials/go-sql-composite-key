package adapter

import (
	"context"
	"database/sql"
	"reflect"
	"strings"

	s "github.com/core-go/sql"

	"go-service/internal/user/model"
)

func NewUserAdapter(db *sql.DB, buildQuery func(*model.UserFilter) (string, []interface{})) (*UserAdapter, error) {
	userType := reflect.TypeOf(model.User{})
	fieldsIndex, _, jsonColumnMap, keys, _, _, buildParam, _, err := s.Init(userType, db)
	if err != nil {
		return nil, err
	}
	return &UserAdapter{DB: db, Map: fieldsIndex, Keys: keys, JsonColumnMap: jsonColumnMap, BuildParam: buildParam, BuildQuery: buildQuery}, nil
}

type UserAdapter struct {
	DB            *sql.DB
	Map           map[string]int
	Keys          []string
	JsonColumnMap map[string]string
	BuildParam    func(int) string
	BuildQuery    func(*model.UserFilter) (string, []interface{})
}

func (r *UserAdapter) All(ctx context.Context) ([]model.User, error) {
	query := `
		select
			company_id,
			user_id,
			username,
			email,
			phone,
			date_of_birth
		from company_users`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var user model.User
		err = rows.Scan(
			&user.CompanyId,
			&user.UserId,
			&user.Username,
			&user.Email,
			&user.Phone,
			&user.DateOfBirth)
		users = append(users, user)
	}
	return users, nil
}

func (r *UserAdapter) Load(ctx context.Context, id model.UserId) (*model.User, error) {
	query := `
		select
			company_id,
			user_id,
			username,
			email,
			phone,
			date_of_birth
		from company_users where company_id = $1 and user_id = $2`
	rows, err := r.DB.QueryContext(ctx, query, id.CompanyId, id.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user model.User
		err = rows.Scan(
			&user.CompanyId,
			&user.UserId,
			&user.Username,
			&user.Email,
			&user.Phone,
			&user.DateOfBirth)
		return &user, nil
	}
	return nil, nil
}

func (r *UserAdapter) Create(ctx context.Context, user *model.User) (int64, error) {
	query := `
		insert into company_users (
			company_id,
			user_id,
			username,
			email,
			phone,
			date_of_birth)
		values (
			$1,
			$2,
			$3, 
			$4,
			$5,
			$6)`
	res, err := r.DB.ExecContext(ctx, query,
		user.CompanyId,
		user.UserId,
		user.Username,
		user.Email,
		user.Phone,
		user.DateOfBirth)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return 0, err
		}
		return -1, err
	}
	return res.RowsAffected()
}

func (r *UserAdapter) Update(ctx context.Context, user *model.User) (int64, error) {
	query := `
		update company_users 
		set
			username = $1,
			email = $2,
			phone = $3,
			date_of_birth = $4
		where company_id = $5 and user_id = $6`
	res, err := r.DB.ExecContext(ctx, query,
		user.Username,
		user.Email,
		user.Phone,
		user.DateOfBirth,
		user.CompanyId,
		user.UserId)
	if err != nil {
		return -1, err
	}
	count, err := res.RowsAffected()
	if count == 0 {
		return count, nil
	}
	return count, err
}

func (r *UserAdapter) Patch(ctx context.Context, user map[string]interface{}) (int64, error) {
	colMap := s.JSONToColumns(user, r.JsonColumnMap)
	query, args := s.BuildToPatch("company_users", colMap, r.Keys, r.BuildParam)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (r *UserAdapter) Delete(ctx context.Context, id model.UserId) (int64, error) {
	query := "delete from company_users where company_id = $1 and user_id = $2"
	res, err := r.DB.ExecContext(ctx, query, id.CompanyId, id.UserId)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (r *UserAdapter) Search(ctx context.Context, filter *model.UserFilter, limit int64, offset int64) ([]model.User, int64, error) {
	var users []model.User
	if filter.Limit <= 0 {
		return users, 0, nil
	}
	query, params := r.BuildQuery(filter)
	pagingQuery := s.BuildPagingQuery(query, limit, offset)
	countQuery := s.BuildCountQuery(query)

	row := r.DB.QueryRowContext(ctx, countQuery, params...)
	if row.Err() != nil {
		return users, 0, row.Err()
	}
	var total int64
	err := row.Scan(&total)
	if err != nil || total == 0 {
		return users, total, err
	}

	err = s.Query(ctx, r.DB, r.Map, &users, pagingQuery, params...)
	return users, total, err
}
