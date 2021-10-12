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
func TestNewWallet(t *testing.T) {
	a:=assert.New(t)
	s:=Service{}
	wallet:=s.newWallet(1)
	a.Equal(data.Money(0),wallet.Balance)
	a.Equal(1,wallet.WalletId)
}
//check it again with deposit function
//func TestGetBalance(t *testing.T) {
//	a:=assert.New(t)
//	s:=Service{}
//	s.userRegister("sina","sinakhorsandi@gmail.com")
//	wal,err:=s.deposit(1,data.Money(1000))
//	if err != nil {
//		t.Errorf("Account Not Exist")
//	}
//	balance,err:=s.getBalance(wal.WalletId)
//	if err != nil {
//		t.Errorf("Account Not Exist")
//	}
//	a.Equal(data.Money(1000),balance)
//}

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


func TestDeposit(t *testing.T) {
	a:=assert.New(t)
	s:=Service{}
	s.userRegister("sina","sinakhorsandi@gmail.com")
	userID,_:=s.findIDByName("sina")
	walletID,_:=s.findWalletID(userID)
	wal,_:=s.deposit(userID,walletID, data.Money(1000000))
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
func TestFindIDByEmail(t *testing.T) {
	a:=assert.New(t)
	s:=Service{}
	s.userRegister("sina1","sinakhorsandi@gmail.com")
	s.userRegister("sina2","eng.sinakhorsandi@gmail.com")
	s.userRegister("sina3","sina76khorsandi@gmail.com")
	id,err:=s.findIDByEmail("eng.sinakhorsandi@gmail.com")
	if err != nil {
		t.Errorf("Account not exist")
	}
	a.Equal(2,id)
}

func TestFindIDByName(t *testing.T) {
	a:=assert.New(t)
	s:=Service{}
	s.userRegister("sina1","sinakhorsandi@gmail.com")
	s.userRegister("sina2","eng.sinakhorsandi@gmail.com")
	s.userRegister("sina3","sina76khorsandi@gmail.com")
	id,err:=s.findIDByName("sina2")
	if err != nil {
		t.Errorf("Account not exist")
	}
	a.Equal(2,id)
}

