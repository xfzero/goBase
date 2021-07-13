package main

import (
	"flag"
	"encoding/xml"
	"io/ioutil"
	"fmt"
)

// type xmlLog struct {
// 	Config string `xml:"config"`
// }

type xmlEtcd struct {
	Addr string `xml:"addr"`
	Node string `xml:"node"`
}

type xmlConfig struct {
	Etcd *xmlEtcd `xml:"etcd"`
	Etcd1 *xmlEtcd `xml:"etcd1"`
	Etcd2 *xmlEtcd `xml:"etcd2"`
}

var (
	gConfig = new(xmlConfig)
)

var configFile = flag.String("config", "./gm_config.xml","")

func LoadConfig(filename string, v interface{}) error {
	if contents, err := ioutil.ReadFile(filename); err != nil {
		return err
	} else {
		if err = xml.Unmarshal(contents, v); err != nil {
			return err
		}
		return nil
	}
}

func main() {
	if err := LoadConfig(*configFile, gConfig); err != nil {
		panic("panic 111 err")
	}

	fmt.Println("gConfig:",gConfig)

	if gConfig.Etcd.Addr != "" {
		fmt.Println("Addr:",gConfig.Etcd.Addr)
	}

	if gConfig.Etcd1.Addr != "" {
		fmt.Println("Addr1:",gConfig.Etcd1.Addr)
	}

	// if gConfig.Etcd2.Addr != "" {
	// 	fmt.Println("Addr1:",gConfig.Etcd1.Addr)
	// }
}