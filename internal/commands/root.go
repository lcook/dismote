package commands

import (
	"math/rand"
	"reflect"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"gitlab.com/lcook/dismote/internal/config"
)

type Data struct {
	Callback   interface{}
	Permission int
}

type Modules map[string]Data

type Handler struct {
	ModuleMap Modules
	Session   *discordgo.Session
	Message   *discordgo.MessageCreate
	Config    *config.Config
}

func (h *Handler) Register(mods Modules) {
	for str, data := range mods {
		cmd := strings.ToLower(str)
		typ := reflect.TypeOf(data.Callback)

		if h.ModuleMap[cmd].Callback != nil || typ.Kind() != reflect.Func {
			return
		}

		h.ModuleMap[cmd] = data
	}
}

func (h *Handler) Execute(str string) {
	cmd := strings.Split(strings.ToLower(str), " ")[0]
	cmd = strings.TrimPrefix(cmd, h.Config.Prefix)
	typ := reflect.TypeOf(h.ModuleMap[cmd].Callback)

	if _, ok := h.ModuleMap[cmd]; !ok || !strings.HasPrefix(str, h.Config.Prefix) {
		h.ModuleMap["default"].Callback.(func())()
		return
	}

	if typ.Kind() != reflect.Func {
		return
	}

	switch h.ModuleMap[cmd].Permission {
	case PermAll:
		h.ModuleMap[cmd].Callback.(func())()
	case PermOwner:
		if h.Message.Author.ID == h.Config.Owner {
			h.ModuleMap[cmd].Callback.(func())()
		}
	default:
		return
	}
}

func New(s *discordgo.Session, m *discordgo.MessageCreate, c *config.Config) *Handler {
	return &Handler{make(Modules), s, m, c}
}

func randomColor() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(16777215-1) + 1
}
