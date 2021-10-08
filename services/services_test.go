package wallet

import (
	"github.com/sinakhorsandi/e-wallet/db"
	"github.com/sinakhorsandi/testify/assert"
	"testing"
)

func TestCheckCustomerExist(t *testing.T) {
	a:=assert.New(t)
	s:=Service{}
	s.userRegister("sina","sinakhorsandi@gmail.com")
	b:=s.checkCustomerExist("sinakhorsandi@gmail.com")
	a.Equal(true,b)
}
func TestUserRegistration(t *testing.T) {
	a:=assert.New(t)
	s:=Service{}
	s.userRegister("sina","sinakhorsandi@gmail.com")
	a.Equal(1,len(s.personalAccounts))
}
func TestCheckPersonalAccountById(t *testing.T) {
	a:=assert.New(t)
	s:=Service{}
	s.userRegister("sina","sinakhorsandi@gmail.com")
	account,_:=s.checkPersonalAccountById(1)
	a.Equal("sina",account.UserName)
}
func TestCheckPersonalWallet(t *testing.T) {
	a:=assert.New(t)
	s:=Service{}
	s.userRegister("sina","sinakhorsandi@gmail.com")
	account,_:=s.checkPersonalWallet(1)
	a.Equal(1,account.WalletId)
}
func TestNewWallet(t *testing.T) {
	a:=assert.New(t)
	s:=Service{}
	fWallet:=s.newWallet(1)
	sWallet:=s.newWallet(2)
	a.Equal(db.Money(0),fWallet.Balance)
	a.Equal(1,fWallet.WalletId)
	a.Equal(2,sWallet.WalletId)
}
func TestDeposit(t *testing.T) {
	a:=assert.New(t)
	s:=Service{}
	s.userRegister("sina","sinakhorsandi@gmail.com")
	wal,_:=s.deposit(1,db.Money(1000000))
	a.Equal(db.Money(1000000),wal.Balance)
}


