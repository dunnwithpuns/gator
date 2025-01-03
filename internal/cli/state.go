package cli

import (
	"github.com/dunnwithpuns/gator/internal/config"
	"github.com/dunnwithpuns/gator/internal/database"
)

type State struct {
	db *database.Queries
	cfg *config.Config
}
