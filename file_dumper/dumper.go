package dumper

import (
	"log"
	"os"
)

// Dumper -
type Dumper struct {
	Config *Config
	File   *os.File
}

// NewDumper -
func NewDumper(topics []string, fileName string, proto bool, json bool) *Dumper {
	log.Println("Setup new Dumper...")

	return &Dumper{
		Config: NewConfig(topics, fileName, proto, json),
	}
}

// Open -
func (d *Dumper) Open() error {
	log.Println("Open Dumper...")

	if d.Config.Proto {
		f, err := os.OpenFile(d.Config.FileName+".proto-test", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		d.File = f
	}
	if d.Config.JSON {
		f, err := os.OpenFile(d.Config.FileName+".encoding-test", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		d.File = f
	}
	return nil
}

// Write -
func (d *Dumper) Write(data []byte) error {
	log.Println("Write message to dumper: %s", data)

	_, err := d.File.Write(data)
	return err
}
