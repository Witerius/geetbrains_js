package main

import (
    "fmt"
    "log"
    "time"
    "gopkg.in/yaml.v2"
    "io/ioutil"

    "github.com/kelseyhightower/envconfig"
)

type conf struct {
    port int64 `yaml:"port"`
    db_url string `yaml:"db_url"`
}

func (c *conf) getConf() *conf {

    yamlFile, err := ioutil.ReadFile("conf.yaml")
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }

    return c
}

func main() {
    var c conf
    c.getConf()

}