package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//ConfigLoader will read the mapping.json file and retrun a map of type ArtNode
func ConfigLoader(configFile string) map[string]ArtMap {

	mapFile, err := ioutil.ReadFile(configFile)

	if err != nil {
		log.Fatal("Critical error :: Cannot open Map file  @ ./data/mapping.json ", err)
	}

	configMap := make(map[string]ArtMap)
	mapData := ArtNodes{}
	if err = json.Unmarshal(mapFile, &mapData); err != nil {
		panic(err)
	}

	//loop through the said data and transfer to map

	for i := 0; i < len(mapData.ArtNodes); i++ {
		configMap[mapData.ArtNodes[i].ArtTitle] = mapData.ArtNodes[i]
	}

	return configMap

}
