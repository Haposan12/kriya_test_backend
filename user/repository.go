package user

import (
	"database/sql"
	"fmt"
	"kriya_test_backend/shared/times"
)

type RepositoryPostgres struct {
	postgres *sql.DB
}

func NewRepository(postgres *sql.DB) *RepositoryPostgres {
	return &RepositoryPostgres{postgres: postgres}
}

func (repo RepositoryPostgres) InsertUser(data CreateUserModel) (err error) {
	now := times.Now(times.TimeGmt, times.TimeFormat)

	query := `INSERT INTO users(id, data, role_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`

	_, err = repo.postgres.Query(query, data.UUID, data.Data, data.RoleID, now, now)

	return
}

func (repo RepositoryPostgres) UpdateUser(data UpdateUserRequestModel) (err error)  {
	now := times.Now(times.TimeGmt, times.TimeFormat)
	email := fmt.Sprintf("%s%s%s", `"`, data.Email, `"`)
	username := fmt.Sprintf("%s%s%s", `"`, data.Username, `"`)
	password := fmt.Sprintf("%s%s%s", `"`, data.Password, `"`)

	query := `UPDATE users
			SET data = jsonb_set(
				jsonb_set(
					jsonb_set(data, '{email}', $1), '{username}', $2
				), '{pasword}', $3), updated_at = $4
				
			WHERE id = $5`

	_, err = repo.postgres.Exec(query, email, username, password, now, data.UUID)

	return
}

func (repo RepositoryPostgres) DeleteUser(uuid string) (err error) {
	query := `DELETE FROM users WHERE id = $1`

	_, err = repo.postgres.Exec(query, uuid)

	return
}

func (repo RepositoryPostgres) GetUserDetail(uuid string) (userDetail GetUserDetailResponse, err error) {
	query := `SELECT u.id, u.data ->> 'username' AS username, u.data ->> 'email' AS email, r.data ->> 'role_name' AS role_name
			FROM users u 
			INNER JOIN roles r ON r.id = u.role_id
			WHERE u.id = $1`

	err = repo.postgres.QueryRow(query, uuid).Scan(&userDetail.UserID, &userDetail.Username, &userDetail.Email, &userDetail.RoleName)

	return
}

func (repo RepositoryPostgres) GetListUser(offset int) (listUser []GetUserListResponse, err error) {
	var user UserModel
	var userResponse GetUserListResponse
	var isActiveModel IsActiveModel

	query := `SELECT data ->> 'email' AS email, data ->> 'username' AS nama, data -> 'status' ->> 'is_active' AS is_active
			FROM users ORDER BY created_at desc LIMIT 5 OFFSET $1`

	rows, err := repo.postgres.Query(query, offset)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&user.Email,
			&user.UserName,
			&user.IsActive)

		if err != nil {
			return
		}

		userResponse.Username = user.UserName
		userResponse.Email = user.Email
		isActiveModel.IsActive = user.IsActive
		userResponse.Status = isActiveModel

		listUser = append(listUser, userResponse)

	}

	return
}
