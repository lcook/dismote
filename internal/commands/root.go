package commands

import (
	"math/rand"
	"reflect"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Commands struct {
	FuncMap map[string]interface{}
	Session *discordgo.Session
	Message *discordgo.MessageCreate
	Prefix  string
}

func (c *Commands) Execute(str string) {
	cmd := strings.Split(strings.ToLower(str), " ")[0]
	cmd = strings.TrimPrefix(cmd, c.Prefix)
	typ := reflect.TypeOf(c.FuncMap[cmd])

	if _, ok := c.FuncMap[cmd]; !ok {
		c.FuncMap["default"].(func())()
		return
	}

	if typ.Kind() != reflect.Func {
		return
	}

	c.FuncMap[cmd].(func())()
}

func (c *Commands) Register(fns map[string]interface{}) {
	for str, fn := range fns {
		cmd := strings.ToLower(str)
		typ := reflect.TypeOf(fn)

		if c.FuncMap[cmd] != nil || typ.Kind() != reflect.Func {
			return
		}

		c.FuncMap[cmd] = fn
	}
}

func New(s *discordgo.Session, m *discordgo.MessageCreate, p string) *Commands {
	return &Commands{
		make(map[string]interface{}), s, m, p,
	}
}

func randomColor() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(16777215-1) + 1
}
