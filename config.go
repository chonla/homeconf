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

func (c *Conf) Read(out interface{}) error {
	b, e := ioutil.ReadFile(c.filename)
	if e != nil {
		return e
	}

	e = yaml.Unmarshal(b, out)
	if e != nil {
		return e
	}
	return nil
}

// ensure is called to ensure file existence. return file name with full paths
func ensure(filename string) (string, error) {
	confPath, e := os.UserHomeDir()
	if e != nil {
		log.Fatal("unable to get home directory")
		return "", e
	}

	confFileName := fmt.Sprintf(`%s%s%s`, confPath, string(os.PathSeparator), filename)

	log.Printf("ensuring %s\n", confFileName)
	fd, e := os.OpenFile(confFileName, os.O_RDONLY|os.O_CREATE, 0666)
	if e != nil {
		return "", e
	}
	defer fd.Close()

	return confFileName, nil
}
