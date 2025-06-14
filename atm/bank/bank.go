package bank

import "example.com/atm/account"

type Bank struct {
	IFSCCode   string
	Name       string
	Branch     string
	AccountMap map[string]*account.Account
}
