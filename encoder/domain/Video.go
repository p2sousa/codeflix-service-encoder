package domain

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	"cloud.google.com/go/storage"
	log "github.com/sirupsen/logrus"
)

type Video struct {
	Uuid string `json:"uuid"`
	Path string `json:"path"`
	Status string `json:"status"`
}

func (video *Video) Unmarshal(payload []byte) Video  {
	err := json.Unmarshal(payload, &video)
	if err != nil {
		panic(err)
	}

	return *video
}

func (video *Video) Download(bucket string, path string) (Video, error)  {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Error(err.Error())
		video.Status = "error"
		return *video, err
	}

	bkt := client.Bucket(bucket)
	obj := bkt.Object(video.Path)
	r, err := obj.NewReader(ctx)
	if err != nil {
		log.Error(err.Error())
		video.Status = "error"
		return *video, err
	}

	defer r.Close()

	body, err := ioutil.ReadAll(r)
	if err != nil {
		log.Error(err.Error())
		video.Status = "error"
		return *video, err
	}

	f, err := os.Create(path + "/" + video.Uuid + "mp4")
	if err != nil {
		log.Error(err.Error())
		video.Status = "error"
		return *video, err
	}

	_, err := f.Write(body)
	if err != nil {
		log.Error(err.Error())
		video.Status = "error"
		return *video, err
	}

	defer f.Close()

	log.Info("Video: " + video.Uuid + "has been storage")

	return *video, nil
}