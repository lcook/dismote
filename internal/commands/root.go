package commands

import (
	"math/rand"
	"reflect"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Modules map[string]interface{}

type Handler struct {
	CmdMap  Modules
	Session *discordgo.Session
	Message *discordgo.MessageCreate
	Prefix  string
}

func (h *Handler) Execute(str string) {
	cmd := strings.Split(strings.ToLower(str), " ")[0]
	cmd = strings.TrimPrefix(cmd, h.Prefix)
	typ := reflect.TypeOf(h.CmdMap[cmd])

	if _, ok := h.CmdMap[cmd]; !ok {
		h.CmdMap["default"].(func())()
		return
	}

	if typ.Kind() != reflect.Func {
		return
	}

	h.CmdMap[cmd].(func())()
}

func (h *Handler) Register(fns Modules) {
	for str, fn := range fns {
		cmd := strings.ToLower(str)
		typ := reflect.TypeOf(fn)

		if h.CmdMap[cmd] != nil || typ.Kind() != reflect.Func {
			return
		}

		h.CmdMap[cmd] = fn
	}
}

func New(s *discordgo.Session, m *discordgo.MessageCreate, p string) *Handler {
	return &Handler{make(Modules), s, m, p}
}

func randomColor() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(16777215-1) + 1
}
