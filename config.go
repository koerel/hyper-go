package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const TYPE_SEQUENCE = "sequence"
const TYPE_RUN_OR_RAISE = "run-or-raise"
const TYPE_COMMAND = "command"
const TYPE_TYPE = "type"

type Settings struct {
	DefaultModifier string `yaml:"default_modifier"`
}

type Mapping struct {
	Key      string `yaml:"key"`
	Mapping  string `yaml:"mapping"`
	Modifier string `yaml:"modifier"`
	Type     string `yaml:"type"`
	Value    string `yaml:"value"`
	Class    string `yaml:"class"`
	Command  string `yaml:"command"`
	Text     string `yaml:"text"`
}

type Config struct {
	Settings Settings  `yaml:"settings"`
	Mappings []Mapping `yaml:"mappings"`
	Keys     []Mapping `yaml:"keys"`
}

func loadConfig(path string) Config {
	c := Config{}
	dat, err := ioutil.ReadFile(path)
	handle(err)
	err = yaml.Unmarshal(dat, &c)
	handle(err)
	for k, m := range c.Keys {
		if m.Modifier == "" {
			m.Modifier = c.Settings.DefaultModifier
		}
		m.Mapping = m.Modifier + "-" + m.Key
		c.Keys[k] = m
	}

	return c
}
