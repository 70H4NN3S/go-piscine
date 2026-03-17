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

type Echo struct{}

func (e Echo) Name() string { return "echo" }
func (e Echo) Help() string { return "echo <text> - prints the text" }
func (e Echo) Run(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("only one argument is allowed")
	}
	fmt.Println(args[0])
	return nil
}

type Upper struct{}

func (u Upper) Name() string { return "upper" }
func (u Upper) Help() string { return "upper <text> - prints the text in upper case" }
func (u Upper) Run(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("needs at least one argument")
	}
	fmt.Println(strings.ToUpper(strings.Join(args, " ")))
	return nil
}
