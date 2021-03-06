package Miscs

import (
	"math/rand"
	"strconv"

	"github.com/genesixx/coalibot/Struct"
	"github.com/genesixx/coalibot/Utils"
)

var tab []int

func Roulette(option string, event *Struct.Message) bool {
	if tab == nil || len(tab) == 0 {
		for i := 0; i < 6; i++ {
			tab = append(tab, 0)
		}
		tab[rand.Intn(6)] = 1
		event.API.PostMessage(event.Channel, "On recharge le revolver!", Struct.SlackParams)
	}
	var count = 6 - len(tab) + 1
	if tab[0] == 1 {
		tab = nil
		event.API.PostMessage(event.Channel, "<@"+event.User+">: Bang ( "+strconv.Itoa(count)+" / 6 )", Struct.SlackParams)
		Utils.HandleRouletteStat(event)
	} else {
		tab = tab[1:]
		event.API.PostMessage(event.Channel, "<@"+event.User+">: Click ( "+strconv.Itoa(count)+" / 6 )", Struct.SlackParams)
	}
	return true
}
