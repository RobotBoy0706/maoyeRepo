package main

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestDecodeCrypto(t *testing.T)  {
	//cost,err:=bcrypt.Cost([]byte("$2a$10$pm2P/WZt6Me9gPD8quTw1e7zDuCGKqbLl/I34KHpcPihNMqSJO46m"))
	//if err != nil {
	//	t.Errorf("cost fail err=%v",err)
	//	return
	//}

	pw,err:=bcrypt.GenerateFromPassword([]byte("16d7a4fca7442dda3ad93c9a726597e4"),3)
	if err != nil {
		t.Errorf("GenerateFromPassword fail err=%v",err)
		return
	}

	t.Logf("pw=%s\n",string(pw))
	err = bcrypt.CompareHashAndPassword([]byte("$2a$10$Iq0/nF9OWRvvld8MPwzcW.sxRqKVa5avLaOOBCNKB6PXWzlskiFZq"), []byte("test1234"))
	if err != nil {
		t.Errorf("CompareHashAndPassword fail err=%v",err)
		return
	}

}
