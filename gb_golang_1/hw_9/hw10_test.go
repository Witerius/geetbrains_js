package hw_9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYamlToStruct(t *testing.T) {
	var result = getConf() //почему-то не срабатывает в тесте считывание файла. при этом, если функцию запусить без теста - работает норм
	asserts := assert.New(t)
	asserts.Equal(8080, result.Port, "Not match")
}
