package gomcbot

import (
	"fmt"
	"testing"
)

func TestPingAndList(t *testing.T) {
	resp, err := PingAndList("localhost", 25565)
	if err != nil {
		t.Errorf("ping and list server fail: %v", err)
	}
	t.Log("Status:" + resp)
}

func TestJoinServerOffline(t *testing.T) {
	p := Auth{
		Name: "Mi_Xi_Xi",
		UUID: "ff7a038f265c4d42b0cf04c575896469",
		AsTk: "",
	}
	g, err := p.JoinServer("localhost", 25565)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Login success")
	events := g.GetEvents()
	go g.HandleGame()

	g.SetChatCallBack(func(msg string, pos byte) {
		if pos == 0 {
			fmt.Println(msg)
		}
	})

	for e := range events {
		switch e {
		case PlayerSpawnEvent:
			fmt.Println(g.player.X, g.player.Y, g.player.Z)
		case PlayerDeadEvent:
			fmt.Println("Player Dead")
		default:
			fmt.Println(e)
		}
	}
}
