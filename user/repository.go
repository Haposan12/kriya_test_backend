package user

import "database/sql"

type RepositoryPostgres struct {
	postgres *sql.DB
}

func NewRepository(postgres *sql.DB) *RepositoryPostgres {
	return &RepositoryPostgres{postgres: postgres}
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
