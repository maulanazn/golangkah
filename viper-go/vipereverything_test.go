package viper_go_test

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestConfigJSON(t *testing.T) {
	var viper *viper.Viper = viper.New()
	t.Run("json", func (t *testing.T) {
		viper.SetConfigFile("config.json")	
		viper.AddConfigPath(".")
	
		err := viper.ReadInConfig()
		assert.Nil(t, err)
	})
	t.Run("readjson", func (t *testing.T) {
		viper.SetConfigFile("config.json")
		viper.AddConfigPath(".")

		err := viper.ReadInConfig()
		assert.Nil(t, err)

		assert.Exactly(t, "Documented", viper.GetString("app.name"))
		assert.Exactly(t, "1.0.0", viper.GetString("app.version"))
		assert.Exactly(t, "Maulana ZN", viper.GetString("app.author"))
	})
}

func TestConfigYAML(t *testing.T) {
	var viper *viper.Viper = viper.New()
	t.Run("yaml", func (t *testing.T) {
		viper.SetConfigFile("config.yaml")	
		viper.AddConfigPath(".")
	
		err := viper.ReadInConfig()
		assert.Nil(t, err)
	})
	t.Run("readyaml", func (t *testing.T) {
		viper.SetConfigFile("config.yaml")
		viper.AddConfigPath(".")

		err := viper.ReadInConfig()
		assert.Nil(t, err)

		assert.Exactly(t, "Documented", viper.GetString("app.name"))
		assert.Exactly(t, "1.0.0", viper.GetString("app.version"))
		assert.Exactly(t, "Maulana ZN", viper.GetString("app.author"))
	})
}

func TestConfigENV(t *testing.T) {
	var viper *viper.Viper = viper.New()
	t.Run("env", func (t *testing.T) {
		viper.SetConfigFile("local.env")	
		viper.AddConfigPath(".")
	
		err := viper.ReadInConfig()
		assert.Nil(t, err)
	})
	t.Run("readenv", func (t *testing.T) {
		viper.SetConfigFile("local.env")
		viper.AddConfigPath(".")

		err := viper.ReadInConfig()
		assert.Nil(t, err)

		assert.Exactly(t, "localhost", viper.GetString("DATABASE_HOST"))
		assert.Exactly(t, 5432, viper.GetInt("DATABASE_PORT"))
		assert.Exactly(t, "maulanazn", viper.GetString("DATABASE_NAME"))
	})
}
