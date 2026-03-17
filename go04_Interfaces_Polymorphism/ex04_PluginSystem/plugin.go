package main

import (
	"fmt"
	"strings"
)

type Command interface {
	Name() string
	Help() string
	Run(args []string) error
}

type Registry struct {
	Commands map[string]Command
}

func (r *Registry) Register(c Command) {
	r.Commands[c.Name()] = c
}

func NewRegistry() *Registry {
	return &Registry{Commands: make(map[string]Command)}
}

func (r *Registry) Execute(input string) error {
	parts := strings.SplitN(input, " ", 2)
	name := parts[0]
	var args []string
	if len(parts) > 1 {
		args = strings.Split(parts[1], " ")
	}
	com, ok := r.Commands[name]
	if ok {
		return com.Run(args)
	}
	return fmt.Errorf("the command %s doesn't exists", name)
}
