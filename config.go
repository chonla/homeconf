package homeconf

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Conf struct for HomeConf
type Conf struct {
	filename string
}

// NewConf creates homeconf instance
func NewConf(filename string) (*Conf, error) {
	fullFilename, e := ensure(filename)
	if e != nil {
		return nil, e
	}
	return &Conf{
		filename: fullFilename,
	}, nil
}

// Read to read configuration and unmarshal it to given interface
func (c *Conf) Read(out interface{}) error {
	b, e := ioutil.ReadFile(c.filename)
	if e != nil {
		return e
	}

	e = yaml.Unmarshal(b, out)
	return e
}

// Write to write interface to configuration file
func (c *Conf) Write(in interface{}) error {
	b, e := yaml.Marshal(in)
	if e != nil {
		return e
	}

	e = ioutil.WriteFile(c.filename, b, 0666)
	return e
}

// ensure is called to ensure file existence. return file name with full paths
func ensure(filename string) (string, error) {
	confPath, e := os.UserHomeDir()
	if e != nil {
		log.Fatal("unable to get home directory")
		return "", e
	}

	confFileName := fmt.Sprintf(`%s%s%s`, confPath, string(os.PathSeparator), filename)

	fd, e := os.OpenFile(confFileName, os.O_RDONLY|os.O_CREATE, 0666)
	if e != nil {
		return "", e
	}
	defer fd.Close()

	return confFileName, nil
}
