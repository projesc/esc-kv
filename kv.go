package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/projesc/esc"
	"github.com/yuin/gopher-lua"
	"log"
	"strings"
	"time"
)

var kv *cache.Cache

func luaGet(vm *lua.LState) int {
	key := vm.ToString(1)
	value := Get(key)
	vm.Push(lua.LString(value))
	return 1
}

func luaSet(vm *lua.LState) int {
	key := vm.ToString(1)
	value := vm.ToString(2)
	Set(key, value)
	return 0
}
func Get(key string) string {
	v, found := kv.Get(key)
	if !found {
		log.Println(key, "not found")
		return ""
	}
	return v.(string)
}

func Set(key string, value string) {
	kv.Set(key, value, cache.NoExpiration)
	esc.Send("*", "set", fmt.Sprintf("%s,%s", key, value))
}

func setEvt(msg *esc.Message) {
	if msg.From != esc.Self() {
		parts := strings.SplitN(msg.Payload, ",", 2)
		if len(parts) == 2 {
			log.Printf("Set %s = %s", parts[0], parts[1])
			kv.Set(parts[0], parts[1], cache.NoExpiration)
		}
	}
}

func syncKv(msg *esc.Message) {
	if msg.Payload == esc.Self() {
		return
	}
	for k, v := range kv.Items() {
		esc.Send("*", "set", fmt.Sprintf("%s,%s", k, v.Object.(string)))
	}
}

func Start(_ *esc.EscConfig) {
	kv = cache.New(6*time.Hour, 1*time.Hour)
	esc.On("*", "set", setEvt)
	esc.On(esc.Self(), "connected", syncKv)
}

func Script(script *esc.Script) {
	script.Lua.SetGlobal("set", script.Lua.NewFunction(luaSet))
	script.Lua.SetGlobal("get", script.Lua.NewFunction(luaGet))
}

func Stop() {
}
