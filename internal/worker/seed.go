package worker

import (
	"encoding/json"
	"fmt"
	"github.com/99neerajsharma/FamTube/internal/contract"
	"github.com/99neerajsharma/FamTube/internal/model"
	"github.com/99neerajsharma/FamTube/internal/utility"
	"go.uber.org/config"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type APIKeys struct {
	index int
	count int
	keys  []string
}

func SeedVideoData(pg *gorm.DB, configYAML *config.YAML) {
	YAMLKeys := configYAML.Get("worker.yt_keys")
	query := configYAML.Get("worker.query").String()
	var parsedKeys []string
	err := YAMLKeys.Populate(&parsedKeys)
	if err != nil {
		log.Println(err)
		return
	}
	keys := APIKeys{index: 0, count: len(parsedKeys), keys: parsedKeys}
	for {
		publishedAfterTime := time.Now().UTC().Add(-1 * time.Hour)
		seedPageData(publishedAfterTime, "", pg, &keys, query)
		time.Sleep(10 * time.Second)
	}
}

func seedPageData(publishedAfterTime time.Time, nextPageToken string, pg *gorm.DB, keys *APIKeys, query string) {
	data := fetchData(publishedAfterTime, nextPageToken, keys, query)
	var filteredData []*model.Video
	if data != nil {
		filteredData = filterData(*data, publishedAfterTime, pg)
	}
	if len(filteredData) > 0 {
		seedIntoDB(filteredData, pg)
	}
	if len(filteredData) == 50 {
		seedPageData(publishedAfterTime, data.NextPageToken, pg, keys, query)
	}
	log.Println("Data seeded for publish after: ", publishedAfterTime)
}

func seedIntoDB(data []*model.Video, pg *gorm.DB) {
	if err := pg.Model(model.Video{}).Create(&data).Error; err != nil {
		log.Println(err)
	}
}

func filterData(data contract.APIData, publishedAfter time.Time, pg *gorm.DB) []*model.Video {
	var filteredData []*model.Video
	var videoIDs []string

	if err := pg.Model(model.Video{}).Where("published_at >= ?", publishedAfter).Select("id").Find(&videoIDs).Error; err != nil {
		log.Println(err)
		return filteredData
	}

	videoIDMap := utility.GetBoolMapFromStringSlice(videoIDs)

	for _, item := range data.Items {
		if _, ok := videoIDMap[item.ID.VideoID]; !ok && item.Snippet.PublishedAt.After(publishedAfter) {
			video := model.Video{ID: item.ID.VideoID, Title: item.Snippet.Title, Description: item.Snippet.Description,
				DefaultThumbnailURL: item.Snippet.Thumbnails.Default.URL, MediumThumbnailURL: item.Snippet.Thumbnails.Medium.URL,
				HighThumbnailURL: item.Snippet.Thumbnails.High.URL, ChannelName: item.Snippet.ChannelTitle, PublishedAt: item.Snippet.PublishedAt}
			filteredData = append(filteredData, &video)
		}
	}
	return filteredData
}

func fetchData(publishedAfterTime time.Time, nextPageToken string, keys *APIKeys, query string) *contract.APIData {
	publishedAfterFormatTime := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02dZ",
		publishedAfterTime.Year(), publishedAfterTime.Month(), publishedAfterTime.Day(),
		publishedAfterTime.Hour(), publishedAfterTime.Minute(), publishedAfterTime.Second())
	url := fmt.Sprintf("https://youtube.googleapis.com/youtube/v3/search?part=snippet&maxResults=50&key=%v&"+
		"type=video&publishedAfter=%v&q=%v&order=date",
		keys.keys[keys.index], publishedAfterFormatTime, query)
	if nextPageToken != "" {
		url += fmt.Sprintf("&pageToken=%v", nextPageToken)
	}
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	var result contract.APIData

	if err != nil {
		log.Println(err)
		return nil
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer func() {
		_ = res.Body.Close()
	}()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil
	}
	if res.StatusCode == 403 {
		keys.index = (keys.index + 1) % keys.count
		log.Println("key changed to index: ", keys.index)
		time.Sleep(2 * time.Second)
		return fetchData(publishedAfterTime, nextPageToken, keys, query)
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("json unmarshal error: ", err)
		return nil
	}
	return &result
}
