package api

import (
	"github.com/backedrum/lipz/models"
)

// API -
type API struct {
	users  *models.UserManager
	quotes *models.QuoteManager
}

// NewAPI -
func NewAPI(db *models.DB) *API {

	usermgr, _ := models.NewUserManager(db)
	quotemgr, _ := models.NewQuoteManager(db)

	return &API{
		users:  usermgr,
		quotes: quotemgr,
	}
}
