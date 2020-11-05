package server

import (
	"encoding/json"
	"io/ioutil"
	"sync"
	"time"

	"github.com/empovit/assisted-agent-simulator/server/models"
	log "github.com/sirupsen/logrus"
)

const reloadInterval = 5 * time.Second

// CommandReader reads commands from a file and returns them one-by-one
type CommandReader struct {
	sync.Mutex
	file     string
	count    int
	commands []models.Command
}

// NewCommandReader returns a new instance that will watch the file passed in the argument
func NewCommandReader(file string) *CommandReader {

	r := &CommandReader{file: file, count: 0}
	r.loadCommands()

	// TODO: Replace with a file watcher
	go func() {
		time.Sleep(reloadInterval)
		r.loadCommands()
	}()

	return r
}

// NextCommand advances to the next command to be returned to the command runner
func (r *CommandReader) NextCommand() *models.Command {

	r.Lock()
	defer r.Unlock()

	if len(r.commands) == 0 {
		return nil
	}

	c := r.commands[r.count%len(r.commands)]
	r.count++
	return &c
}

// AllCommands returns a copy of all currently loaded commands
func (r *CommandReader) AllCommands() []models.Command {
	r.Lock()
	defer r.Unlock()
	return r.commands
}

func (r *CommandReader) loadCommands() {

	log.Infof("Loading commands from %q", r.file)

	f, err := ioutil.ReadFile(r.file)
	if err != nil {
		log.WithError(err).Errorf("Failed reading from %q: %s", r.file, err)
		return
	}

	r.Lock()
	defer r.Unlock()

	if err := json.Unmarshal([]byte(f), &r.commands); err != nil {
		log.WithError(err).Errorf("Failed un-marshaling commands from JSON file %q: %s", r.file, err)
	}
}
