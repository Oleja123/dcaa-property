package propertydb

import (
	"context"
	"fmt"
	"time"

	"github.com/Oleja123/dcaa-property/internal/domain/property"
	"github.com/Oleja123/dcaa-property/pkg/client/postgresql"
	"github.com/jackc/pgx/v5/pgconn"
)

type repository struct {
	client postgresql.Client
}

func (r *repository) Create(ctx context.Context, p property.Property) (int, error) {
	q := `
		INSERT INTO properties (addr, price, info, category_id, property_name, last_update) 
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	err := r.client.QueryRow(ctx, q, p.Addr, p.Price, p.Info, p.CategoryId, p.Name, time.Now()).Scan(&p.Id)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			return 0, fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where)
		}
		return 0, err
	}
	return p.Id, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	q := `
		DELETE FROM properties WHERE id = $1
	`
	_, err := r.client.Exec(ctx, q, id)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			return fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where)
		}
		return err
	}
	return nil
}

func (r *repository) FindAll(ctx context.Context) ([]property.Property, error) {
	q := `
		SELECT id, addr, price, info, category_id, last_update, property_name FROM properties
	`

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	properties := make([]property.Property, 0)
	for rows.Next() {
		var p property.Property

		err := rows.Scan(&p.Id, &p.Addr, &p.Price, &p.Info, &p.CategoryId, &p.LastUpdate, &p.Name)
		if err != nil {
			return nil, err
		}

		properties = append(properties, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return properties, nil
}

func (r *repository) FindOne(ctx context.Context, id int) (property.Property, error) {
	q := `
		SELECT id, addr, price, info, category_id, last_update, property_name FROM properties
		WHERE id = $1
	`

	var p property.Property
	err := r.client.QueryRow(ctx, q, id).Scan(&p.Id, &p.Addr, &p.Price, &p.Info, &p.CategoryId, &p.LastUpdate, &p.Name)
	if err != nil {
		return property.Property{}, err
	}

	return p, nil
}

func (r *repository) Update(ctx context.Context, p property.Property) error {
	q := `
		UPDATE properties SET addr = $1, property_name = $2, price = $3, info = $4, category_id = $5, last_update = $6
		WHERE id = $7
	`

	_, err := r.client.Exec(ctx, q, p.Addr, p.Name, p.Price, p.Info, p.CategoryId, time.Now(), p.Id)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			return fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where)
		}
		return err
	}
	return nil
}

func NewRepository(client postgresql.Client) property.Repository {
	return &repository{
		client: client,
	}
}
