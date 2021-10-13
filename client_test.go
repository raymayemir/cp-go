package cp_go

import (
	"testing"
	"time"
)

func TestClient_Ping(t *testing.T) {

	client := NewClient(Config{
		ApiSecret: "YOUR_API_SECRET",
		PublicId:  "YOUR_PUBLIC_ID",
		Timeout:   30 * time.Second,
	})

	response, err := client.Ping()

	if err != nil || response["Success"] != true {
		t.Error(err)
	}
}
