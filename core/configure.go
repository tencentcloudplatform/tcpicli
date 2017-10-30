package core

import (
	"errors"
	"github.com/go-ini/ini"
	"log"
	"os"
	"os/user"
)

type ConfigFile struct {
	*ini.File
}

type Configure struct {
	SecretId      string
	SecretKey     string
	DefaultRegion string
	Internal      bool
}

var (
	configFile ConfigFile
	config     *Configure = new(Configure)
	cuser, _              = user.Current()
	configDir             = cuser.HomeDir + "/.tcpicli"
	configLoc             = configDir + "/config"
)

func ListProfiles() []*ini.Section {
	var profiles []*ini.Section
	for _, v := range configFile.Sections() {
		if v.Name() != "DEFAULT" {
			profiles = append(profiles, v)
		}
	}
	return profiles
}

func DefaultRegion() string {
	return config.DefaultRegion
}
func Internal() bool {
	return config.Internal
}

func CurrentProfile() string {
	return configFile.Section("DEFAULT").Key("current").String()
}

func CreateProfile(name string, conf *Configure) error {
	for _, v := range ListProfiles() {
		if name == v.Name() {
			return errors.New("config already exist")
		}
	}
	configFile.Section(name).Key("secretid").SetValue(conf.SecretId)
	configFile.Section(name).Key("secretkey").SetValue(conf.SecretKey)
	configFile.Section(name).Key("defaultregion").SetValue(conf.DefaultRegion)
	internal := "false"
	if conf.Internal {
		internal = "true"
	}
	configFile.Section(name).Key("internal").SetValue(internal)
	return configFile.SaveTo(configLoc)
}

func SwitchProfile(name string) error {
	for _, v := range ListProfiles() {
		if v.Name() == name {
			configFile.Section("DEFAULT").Key("current").SetValue(name)
			configFile.SaveTo(configLoc)
			return nil
		}
	}
	return errors.New("profile not found")
}

func selectProfile() {
	secretId := os.Getenv("SECRETID")
	secretKey := os.Getenv("SECRETKEY")
	defaultRegion := os.Getenv("DEFAULTREGION")
	internal := os.Getenv("INTERNAL")
	if len(secretId) > 0 {
		config.SecretId = secretId
		config.SecretKey = secretKey
		config.DefaultRegion = defaultRegion
		if internal == "true" {
			config.Internal = true
		}
	} else {
		profile := CurrentProfile()
		config.SecretId = configFile.Section(profile).Key("secretid").String()
		config.SecretKey = configFile.Section(profile).Key("secretkey").String()
		config.DefaultRegion = configFile.Section(profile).Key("defaultregion").String()
		config.Internal, _ = configFile.Section(profile).Key("internal").Bool()
	}
}

func chkConf() {
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		os.Mkdir(configDir, 0744)
	}
	if _, err := os.Stat(configLoc); os.IsNotExist(err) {
		os.Create(configLoc)
		os.Chmod(configLoc, 0700)
		loadConfig()
		configFile.Section("DEFAULT").Key("current").SetValue("default")
		configFile.Section("default").Key("secretid").SetValue("input your SecretId here")
		configFile.Section("default").Key("secretkey").SetValue("input your SecretKey here")
		configFile.Section("default").Key("defaultregion").SetValue("input your DefaultRegion here")
		configFile.Section("default").Key("internal").SetValue("false")
		configFile.SaveTo(configLoc)
	}
}

func loadConfig() {
	cfg, err := ini.Load(configLoc)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	configFile = ConfigFile{cfg}
}

func init() {
	chkConf()
	loadConfig()
	selectProfile()
}
