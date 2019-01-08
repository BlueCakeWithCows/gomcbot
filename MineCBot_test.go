package gomcbot

import (
	"testing"
)

func TestPingAndList(t *testing.T) {
	resp, err := PingAndList("localhost", 25565)
	if err != nil {
		t.Errorf("ping and list server fail: %v", err)
	}
	t.Log("Status:" + resp)
}

func TestJoinServer(t *testing.T) {
	p := Auth{
		Name: "Mi_Xi_Xi",
		UUID: "ff7a038f-265c-4d42-b0cf-04c575896469",
		AsTk: "",
	}
	g, err := p.JoinServer("localhost", 25565)
	if err != nil {
		t.Fatal(err)
	}
	err = g.HandleGame()
	if err != nil {
		t.Fatal(err)
	}
}
