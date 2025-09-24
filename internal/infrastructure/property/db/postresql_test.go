package propertydb_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Oleja123/dcaa-property/internal/domain/property"
	propertydb "github.com/Oleja123/dcaa-property/internal/infrastructure/property/db"
	"github.com/Oleja123/dcaa-property/pkg/client"
	"github.com/Oleja123/dcaa-property/pkg/client/postgresql"
	"github.com/Oleja123/dcaa-property/pkg/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) client.Client {
	cfg, err := config.LoadConfig("../../../../test_config .yaml")
	require.NoError(t, err)

	client, err := postgresql.NewClient(t.Context(), cfg)
	require.NoError(t, err)

	_, err = client.Exec(context.Background(), "TRUNCATE TABLE properties RESTART IDENTITY CASCADE")
	require.NoError(t, err)

	return client
}

func TestPropertyRepository_CRUD(t *testing.T) {
	ctx := context.Background()
	client := setupTestDB(t)
	repo := propertydb.NewRepository(client)

	p := property.Property{
		Name:       "Villa",
		Addr:       "Main St 123",
		Price:      sql.NullFloat64{Float64: 500000, Valid: true},
		Info:       sql.NullString{String: "Nice house", Valid: true},
		CategoryId: 1,
	}

	id, err := repo.Create(ctx, p)
	require.NoError(t, err)
	assert.Greater(t, id, 0)

	got, err := repo.FindOne(ctx, id)
	require.NoError(t, err)
	assert.Equal(t, "Villa", got.Name)
	assert.Equal(t, "Main St 123", got.Addr)
	assert.Equal(t, 500000.0, got.Price.Float64)
	assert.True(t, got.Info.Valid)
	assert.Equal(t, "Nice house", got.Info.String)

	got.Name = "Updated Villa"
	got.Price = sql.NullFloat64{Float64: 600000, Valid: true}
	err = repo.Update(ctx, got)
	require.NoError(t, err)

	updated, err := repo.FindOne(ctx, id)
	require.NoError(t, err)
	assert.Equal(t, "Updated Villa", updated.Name)
	assert.Equal(t, 600000.0, updated.Price.Float64)

	all, err := repo.FindAll(ctx)
	require.NoError(t, err)
	assert.Len(t, all, 1)

	err = repo.Delete(ctx, id)
	require.NoError(t, err)

	_, err = repo.FindOne(ctx, id)
	assert.Error(t, err)
}
