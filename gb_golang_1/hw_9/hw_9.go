package hw_9

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"regexp"
)

type conf struct {
	Port  int    `yaml:"Port"`
	DbUrl string `yaml:"DbUrl"`
}

func getConf() conf {

	var c = conf{}

	yamlFile, err := ioutil.ReadFile("gb_golang_1/hw_9/conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func main() {
	var conf = getConf()

	log.Println(conf.DbUrl)
	log.Println(conf.Port)

	urlMathCondition, err := regexp.MatchString("(http[s]?:\\/\\/)?([^\\/\\s]+\\/)(.*)", conf.DbUrl)
	if err != nil {
		log.Printf("Url match error: %v", err)
	}

	log.Printf("Pattern match condition: %v", urlMathCondition)

}
