package clog

import (
	"io"
)


type Topic struct {
	name   string
	config TopicConfig
	log    *Log
	writer *io.Writer
}

func newTopic(config TopicConfig) *Topic {
	return &Topic{}
}
