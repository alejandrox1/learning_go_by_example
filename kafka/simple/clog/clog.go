package clog

import (
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const configFile = "config.json"

var (
	logDirFlag = flag.String("log-dir", "tmp/clog-logs",
		"Comma separated list of directories under which to store log files")
)

type Clog struct {
	LogDir string
}

type TopicConfig struct {}


// Create log directories from a slice of strings and return a Clog struct.
func NewClog(logDir string) (*Clog, error) {
	ld, err := os.Stat(logDir)

	// Create logDir if it doesn't exist, return error if it is a file.
	if os.IsNotExist(err) {
		err = os.Mkdir(logDir, 0755)
		if err != nil {
			return nil, err
		}
	} else if !ld.IsDir() {
		return nil, errors.New("log dir is not a directory")
	}

	return &Clog{LogDir: logDir}, nil
}


//
func (c *Clog) Register(name string, topic Topic) error {

}

// 
func (c *Clog) initTopic(name string) error {
	topicPath := filepath.Join(c.LogDir, name)
	configPath := filepath.Join(topicPath, configFile)

	f, err := os.OpenFile(configPath, os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	var config TopicConfig
	if err = json.NewDecoder(f).Decode(&config); err != nil {
		return err
	}

	topic := newTopic(config)
	err = c.Register(name, topic)
	return err
}

// 
func (c *Clog) initTopics() error {
	files, err := ioutil.ReadDir(c.LogDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		if err = c.initTopic(file.Name()); err !=nil {
			break
		}
	}
	return err
}

