package conf

import (
	"github.com/spf13/viper"
	"fmt"
)

type db struct {
	Host      string
	Database  string
	Username  string
	Password  string
	Charset   string
	ParseTime bool
	Loc       string
}

type gen struct {
	Secret string
}

type Conf struct {
	DB db
	General gen
}


func Init(path string, name string) Conf {
	viper.SetConfigName(name)
	viper.AddConfigPath(path)
	err := viper.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	c := Conf{db{}, gen{}}

	err = viper.UnmarshalKey("Database", &c.DB)
	if err != nil {
		panic(fmt.Errorf("Cannot process configuration: %s \n", err))
	}
	err = viper.UnmarshalKey("general", &c.General)
	if err != nil {
		panic(fmt.Errorf("Cannot process configuration: %s \n", err))
	}

	return c
}

func (d db) MarshallString() string {
	return fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=%t&loc=%s",
		d.Username,
		d.Password,
		d.Database,
		d.Charset,
		d.ParseTime,
		d.Loc,
	)
}
