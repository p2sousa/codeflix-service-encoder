package main

import (
	v "encoder/domain"
	log "github.com/sirupsen/logrus"
)

func main()  {
	var video v.Video

	data := []byte("{\"uuid\":\"123\", \"Path\":\"video.mp4\", \"status\":\"pending\"}")
	video.Unmarshal(data)

	log.Info(video.Uuid)
}