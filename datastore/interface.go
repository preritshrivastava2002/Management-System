package datastore

import (
	"github.com/preritshrivastava2002/Management-System/model"
	"gofr.dev/pkg/gofr"
)

type House interface {
	StartHandler(c *gofr.Context) error
	AllHomeHandler(c *gofr.Context) ([]model.House, error)
	CreateHomeHandler(c *gofr.Context, model *model.House) (*model.House, error)
	GetHomeHandler(c *gofr.Context, id string) (*model.House, error)
	UpdateHomeHandler(c *gofr.Context, model *model.House) error
	DeleteHomeHandler(c *gofr.Context, id string) error
}
