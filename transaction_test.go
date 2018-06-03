package ncoin_wallet

import (
	"testing"
)

func TestSignTransactionFull(t *testing.T) {

	transaction := NewTransaction(12318464, 5468421, 25.6, 1.4)
	err = transaction.Sign(privateKey)
	if err != nil {
		t.Errorf("Fail in sign transaction %s", err.Error())
	}

	err = transaction.Verify(publicKey)
	if err != nil {
		t.Errorf("Fail to verify transaction %s", err.Error())
	}

}

func TestSignTransactionEmpty(t *testing.T) {

	transaction := Transaction{}
	err := transaction.Sign(privateKey)
	if err != nil {
		t.Errorf("Fail in sign transaction %s", err.Error())
	}

	err = transaction.Verify(publicKey)
	if err != nil {
		t.Errorf("Fail to verify transaction %s", err.Error())
	}

}

func TestFailSignByChange(t *testing.T) {

	transaction := NewTransaction(12318464, 5468421, 25.6, 1.4)
	err := transaction.Sign(privateKey)
	if err != nil {
		t.Errorf("Fail in sign transaction %s", err.Error())
	}
	transaction.fee = 8.6
	err = transaction.Verify(publicKey)
	if err == nil {
		t.Errorf("Should fail because fee atribute is change: %s", err.Error())
	}

}

func TestFailSignByBadCredentials(t *testing.T) {

	transaction := NewTransaction(12318464, 5468421, 25.6, 1.4)
	err := transaction.Sign(privateKey)
	if err != nil {
		t.Errorf("Fail in sign transaction %s", err.Error())
	}
	_, publicKeyTest, _ := GenerateKeys()
	err = transaction.Verify(publicKeyTest)
	if err == nil {
		t.Errorf("Should fail because the public key its another; %s", err.Error())
	}

}
