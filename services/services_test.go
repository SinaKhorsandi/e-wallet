package wallet

import (
	"github.com/sinakhorsandi/e-wallet/data"
	"github.com/sinakhorsandi/testify/assert"
	"testing"
)

func TestCheckCustomerExist(t *testing.T) {
	a:=assert.New(t)
	s:=Service{}
	s.userRegister("sina1","sinakhorsandi.dev@gmail.com")
	s.userRegister("sina2","eng.sinakhorsandi@gmail.com")
	b:=s.checkCustomerExist("sinakhorsandi.dev@gmail.com")
	a.Equal(true,b)
}

func TestUserRegistration(t *testing.T) {
	s:=Service{}
	s.userRegister("sina","sinakhorsandi@gmail.com")
	s.userRegister("sina","sinakhorsandii@gmail.com")
	s.userRegister("sina","eng.sinakhorsandi@gmail.com")
	account,_:=s.findIDByEmail("sinakhorsandii@gmail.com")
	a:=assert.New(t)
	a.Equal(2,account)
}

//check it again with deposit function
func TestGetBalance(t *testing.T) {
	a:=assert.New(t)
	s:=Service{}
	s.userRegister("sina","sinakhorsandi@gmail.com")
	account,_:=s.findIDByName("sina")
	balance,_:=s.getBalance(account)
	a.Equal(data.Money(0),balance)
}

func TestGetWallet(t *testing.T) {
	s:=Service{}
	s.userRegister("sina","sinakhorsandi@gmail.com")
	userID,_:=s.findIDByName("sina")
	walletID,_:=s.findWalletID(userID)
	wal,_:=s.getWallet(walletID)
	if wal == nil {
		t.Errorf("fail")
	}
}

func TestNewWallet(t *testing.T) {
	a:=assert.New(t)
	s:=Service{}
	wallet:=s.newWallet(1)
	a.Equal(data.Money(0),wallet.Balance)
	a.Equal(1,wallet.WalletId)
}
func TestDeposit(t *testing.T) {
	a:=assert.New(t)
	s:=Service{}
	s.userRegister("sina","sinakhorsandi@gmail.com")
	userID,_:=s.findIDByName("sina")
	walletID,_:=s.findWalletID(userID)
	wal,_:=s.deposit(walletID, data.Money(1000000))
	a.Equal(data.Money(1000000),wal.Balance)
}

//problem
func TestAccountCount(t *testing.T) {
	a:=assert.New(t)
	s:=Service{}
	s.userRegister("sina1","sinakhorsandi@gmail.com")
	s.userRegister("sina2","eng.sinakhorsandi@gmail.com")
	s.userRegister("sina3","sina76khorsandi@gmail.com")

	n:=s.countAccount()
	a.Equal(3,n)
}



