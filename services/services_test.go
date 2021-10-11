package wallet

import (
	"github.com/sinakhorsandi/testify/assert"
	"testing"
)

func TestCheckCustomerExist(t *testing.T) {
	a:=assert.New(t)
	s:=Service{}
	s.userRegister("sina","sinakhorsandi.dev@gmail.com")
	b:=s.checkCustomerExist("sinakhorsandi.dev@gmail.com")
	a.Equal(false,b)
}
//check this test again
func TestUserRegistration(t *testing.T) {
	s:=Service{}
	s.userRegister("sina","sinakhorsandi@gmail.com")
	account,_:=s.FindEmailByID(1)
	a:=assert.New(t)
	a.Equal("sinakhorsandi@gmail.com",account.Email)
}
//func TestCheckPersonalAccountById(t *testing.T) {
//	a:=assert.New(t)
//	s:=Service{}
//	s.userRegister("sina","sinakhorsandi@gmail.com")
//	account,_:=s.checkPersonalAccountById(1)
//	a.Equal("sina",account.UserName)
//}
//func TestCheckPersonalWallet(t *testing.T) {
//	a:=assert.New(t)
//	s:=Service{}
//	s.userRegister("sina","sinakhorsandi@gmail.com")
//	account,_:=s.checkPersonalWallet(1)
//	a.Equal(1,account.WalletId)
//}
//func TestNewWallet(t *testing.T) {
//	a:=assert.New(t)
//	s:=Service{}
//	fWallet:=s.newWallet(1)
//	sWallet:=s.newWallet(2)
//	a.Equal(data.Money(0),fWallet.Balance)
//	a.Equal(1,fWallet.WalletId)
//	a.Equal(2,sWallet.WalletId)
//}
//func TestDeposit(t *testing.T) {
//	a:=assert.New(t)
//	s:=Service{}
//	s.userRegister("sina","sinakhorsandi@gmail.com")
//	wal,_:=s.deposit(1, data.Money(1000000))
//	a.Equal(data.Money(1000000),wal.Balance)
//}
//
//
