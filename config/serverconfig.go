package config

import (
	"encoding/xml"
	"io/ioutil"
)

var serverconfigInst *ServerConfig

type ServerConfig struct {
	xmlconf XmlConfig
}
type XmlConfig struct {
	Server   XmlServer   `xml:"Server"`
	Database XmlDatabase `xml:"Database"`
}

type XmlServer struct {
	Ip           string `xml:"Ip,attr"`
	Port         uint16 `xml:"Port,attr"`
	AccessPtMode int    `xml:"AccessPtMode,attr"`
	AccessPtCode string `xml:"AccessPtCode,attr"`
}

type XmlDatabase struct {
	Ip   string `xml:"Ip,attr"`
	Port uint16 `xml:"Port,attr"`
	Db   string `xml:"Db,attr"`
	User string `xml:"User,attr"`
	Pwd  string `xml:"Pwd,attr"`
}

func GetServerConfig() *ServerConfig {
	if serverconfigInst == nil {
		serverconfigInst = &ServerConfig{}
	}

	return serverconfigInst
}

func (this *ServerConfig) InitConfig() bool {
	//TODO

	return this.parseXml()
}

func (this *ServerConfig) parseXml() bool {
	xmlcontent, err := ioutil.ReadFile("serverconfig.xml")
	if err != nil {
		panic(err)
		return false
	}
	err = xml.Unmarshal(xmlcontent, &this.xmlconf)
	if err != nil {
		panic(err)
		return false
	}
	return true
}

func (this *ServerConfig) GetXmlConf() *XmlConfig {
	return &this.xmlconf
}
