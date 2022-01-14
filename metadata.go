//go:build metadata
// +build metadata

package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/aerokube/selenoid/event"
	"github.com/aerokube/selenoid/session"
)

const metadataFileExtension = ".json"

func init() {
	mp := &MetadataProcessor{}
	event.AddSessionStoppedListener(mp)
	log.Println("[-] [INIT] [Will save sessions metadata]")
}

type MetadataProcessor struct {
}

func (mp *MetadataProcessor) OnSessionStopped(stoppedSession event.StoppedSession) {
	if logOutputDir != "" {
		meta := session.Metadata{
			ID:           stoppedSession.SessionId,
			Started:      stoppedSession.Session.Started,
			Finished:     time.Now(),
			Capabilities: stoppedSession.Session.Caps,
		}
		data, err := json.MarshalIndent(meta, "", "    ")
		if err != nil {
			log.Printf("[%d] [METADATA] [%s] [Failed to marshal: %v]", stoppedSession.RequestId, stoppedSession.SessionId, err)
			return
		}
		filename := filepath.Join(logOutputDir, stoppedSession.SessionId+metadataFileExtension)
		err = os.WriteFile(filename, data, 0644)
		if err != nil {
			log.Printf("[%d] [METADATA] [%s] [Failed to save to %s: %v]", stoppedSession.RequestId, stoppedSession.SessionId, filename, err)
			return
		}
		log.Printf("[%d] [METADATA] [%s] [%s]", stoppedSession.RequestId, stoppedSession.SessionId, filename)
		createdFile := event.CreatedFile{
			Event: stoppedSession.Event,
			Name:  filename,
			Type:  "metadata",
		}
		event.FileCreated(createdFile)
	}
}
