package auth

import "database/sql"

type RepositoryPostgres struct {
	postgres *sql.DB
}

func NewRepository(postgres *sql.DB) *RepositoryPostgres {
	return &RepositoryPostgres{postgres: postgres}
}

func (repo RepositoryPostgres) GetUsernameAndPasswordAdmin(username, password string) (isAdmin bool, err error) {
	query := `SELECT EXISTS(
				SELECT 1 FROM users u
				INNER JOIN roles r ON r.id = u.role_id
				WHERE u.data ->> 'username' = $1 AND u.data ->> 'password' = $2
				AND r.data ->> 'role_name' = 'Admin'
			)`

	err = repo.postgres.QueryRow(query, username, password).Scan(&isAdmin)

	return
}
