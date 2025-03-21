package repositories

import (
	"context"
	"fmt"
	"itsware/internal/constants"
	"itsware/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	DB *pgxpool.Pool
}

func CreateUser(ctx context.Context, user models.User) error {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return fmt.Errorf("database connection not found in context")
	}

	query := `SELECT users.create(ROW(NULL::int, $1::citext, $2::citext, $3::citext, $4::citext, $5::text, $6::int, $7::int, $8::timestamp, $9::int, NULL::timestamp, NULL::int)::users.user)`
	_, err := conn.Exec(ctx, query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.Password,
		user.RoleID,
		user.TenantID,
		user.CreatedOn,
		user.CreatedBy,
	)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func GetUser(ctx context.Context, id int) (*models.User, error) {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return &models.User{}, fmt.Errorf("database connection not found in context")
	}

	var user models.User
	query := `SELECT * FROM users.get_one($1)`
	row := conn.QueryRow(ctx, query, id)
	err := row.Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password,
		&user.Phone, &user.RoleID, &user.TenantID, &user.CreatedOn, &user.CreatedBy,
		&user.LastModifiedOn, &user.LastModifiedBy,
	)
	if err != nil {
		return &models.User{}, err
	}
	return &user, nil
}

func GetAllUsers(ctx context.Context) ([]models.User, error) {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return []models.User{}, fmt.Errorf("database connection not found in context")
	}

	query := `SELECT * FROM users.get_all()`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Phone,
			&user.Password, &user.RoleID, &user.TenantID, &user.CreatedOn, &user.CreatedBy,
			&user.LastModifiedOn, &user.LastModifiedBy); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func UpdateUser(ctx context.Context, user models.User) error {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return fmt.Errorf("database connection not found in context")
	}

	query := `SELECT users.update(ROW($1::int, $2::citext, $3::citext, $4::citext, $5::text, NULL::citext, NULL::int, NULL::int, NULL::timestamp, NULL::int, NULL::timestamp, NULL::int)::users.user)`
	_, err := conn.Exec(ctx, query,
		user.ID,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
	)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func DeleteUser(ctx context.Context, id int) error {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return fmt.Errorf("database connection not found in context")
	}

	query := `SELECT users.delete($1)`
	_, err := conn.Exec(ctx, query, id)
	return err
}

func (r *User) GetUserByEmailAndPassword(email, password string) (models.User, error) {
	var user models.User
	var role models.Role

	query := `SELECT u.id, u.email, u.password, u.role_id, r.name
              FROM users u
              JOIN roles r ON u.role_id = r.id
              WHERE u.email = $1 AND u.password = $2`

	err := r.DB.QueryRow(context.Background(), query, email, password).Scan(
		&user.ID, &user.Email, &user.Password, &user.RoleID, &role.Name,
	)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *User) GetUserByEmail(email string) (user models.User, err error) {
	query := `SELECT id, email, password, role_id, tenant_id FROM users WHERE email = $1`

	err = r.DB.QueryRow(context.Background(), query, email).Scan(
		&user.ID, &user.Email, &user.Password, &user.RoleID, &user.TenantID,
	)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *User) GetRoleByID(roleID int) (string, error) {
	var role models.Role

	query := `SELECT id, name FROM roles WHERE id = $1`
	err := r.DB.QueryRow(context.Background(), query, roleID).Scan(&role.ID, &role.Name)
	if err != nil {
		return "", err
	}

	return role.Name, nil
}
