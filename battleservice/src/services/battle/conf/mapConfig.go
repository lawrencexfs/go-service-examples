package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/cihub/seelog"
	"github.com/spf13/viper"
)

type MapNodeConfig struct {
	Id     uint32  `json:"id"`
	Type   int     `json:"type"`
	Px     float64 `json:"px"`
	Py     float64 `json:"py"`
	Radius float64 `json:"radius"`
}

type MapConfig struct {
	Title string           `json:"title"`
	Size  float64          `json:"size"`
	Nodes []*MapNodeConfig `json:"nodes"`
}

var _mapConfigDic map[uint64]*MapConfig

func LoadMapConfig(path string) (*MapConfig, bool) {
	config := MapConfig{}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		seelog.Info("LoadMapConfig fail:", err.Error())
		return nil, false
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		seelog.Info("LoadMapConfig ummarshal fail:", err.Error())
		return nil, false
	}

	seelog.Info("LoadMapConfig:", config.Size, " nodes:", len(config.Nodes))
	return &config, true
}

func InitMapConfig() bool {
	_mapConfigDic = make(map[uint64]*MapConfig)
	for _, m := range ConfigMgr_GetMe().Map.Scenes {
		path := fmt.Sprintf("%s%d.json", viper.GetString("global.terraincfg"), m.Id)
		seelog.Info("LoadMapConfig:" + path)
		if config, ok := LoadMapConfig(path); ok {
			_mapConfigDic[m.Id] = config
		} else {
			return false
		}

	}

	return true
}

func GetMapConfigById(sceneID uint64) *MapConfig {
	val := _mapConfigDic[sceneID]
	return val
}
