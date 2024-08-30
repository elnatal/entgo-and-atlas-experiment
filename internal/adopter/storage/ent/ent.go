package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/elnatal/go-experiment/ent"
	"github.com/elnatal/go-experiment/internal/adopter/config"
	_ "github.com/lib/pq"
)

type Ent struct {
	Client *ent.Client
}

func NewEnt(ctx context.Context, config *config.DB) (*Ent, error) {
	client, err := ent.Open("postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			config.Host,
			config.Port,
			config.User,
			config.Name,
			config.Password),
	)
	if err != nil {
		return nil, err
	}

	return &Ent{
		Client: client,
	}, nil
}

// Migrate runs the database migration
func (ent *Ent) Migrate() error {
	if err := ent.Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return nil
}

// Close closes the database connection
func (ent *Ent) Close() {
	ent.Client.Close()
}
