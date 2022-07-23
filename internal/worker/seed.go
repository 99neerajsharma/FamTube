package worker

import (
	"encoding/json"
	"fmt"
	"github.com/99neerajsharma/FamTube/internal/contract"
	"github.com/99neerajsharma/FamTube/internal/model"
	"github.com/99neerajsharma/FamTube/internal/utility"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"time"
)

func SeedVideoData(pg *gorm.DB) {
	for {
		publishedAfterTime := time.Now().UTC().Add(-1 * time.Hour)
		seedPageData(publishedAfterTime, "", pg)
		time.Sleep(10 * time.Second)
	}
}

func seedPageData(publishedAfterTime time.Time, nextPageToken string, pg *gorm.DB) {
	data := fetchData(publishedAfterTime, nextPageToken)
	var filteredData []*model.Video
	if data != nil {
		filteredData = filterData(*data, publishedAfterTime, pg)
	}
	if len(filteredData) > 0 {
		seedIntoDB(filteredData, pg)
	}
	if len(filteredData) == 50 {
		seedPageData(publishedAfterTime, data.NextPageToken, pg)
	}
}

func seedIntoDB(data []*model.Video, pg *gorm.DB) {
	if err := pg.Model(model.Video{}).Create(&data).Error; err != nil {
		fmt.Println(err)
	}
}

func filterData(data contract.APIData, publishedAfter time.Time, pg *gorm.DB) []*model.Video {
	var filteredData []*model.Video
	var videoIDs []string

	if err := pg.Model(model.Video{}).Where("published_at >= ?", publishedAfter).Select("id").Find(&videoIDs).Error; err != nil {
		fmt.Println(err)
		return filteredData
	}

	videoIDMap := utility.GetBoolMapFromStringSlice(videoIDs)

	for _, item := range data.Items {
		if _, ok := videoIDMap[item.ID.VideoID]; !ok && item.Snippet.PublishedAt.After(publishedAfter.Add(-100*time.Hour)) {
			video := model.Video{ID: item.ID.VideoID, Title: item.Snippet.Title, Description: item.Snippet.Description,
				DefaultThumbnailURL: item.Snippet.Thumbnails.Default.URL, MediumThumbnailURL: item.Snippet.Thumbnails.Medium.URL,
				HighThumbnailURL: item.Snippet.Thumbnails.High.URL, ChannelName: item.Snippet.ChannelTitle, PublishedAt: item.Snippet.PublishedAt}
			filteredData = append(filteredData, &video)
		}
	}
	return filteredData
}

func fetchData(publishedAfterTime time.Time, nextPageToken string) *contract.APIData {
	publishedAfterFormatTime := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02dZ",
		publishedAfterTime.Year(), publishedAfterTime.Month(), publishedAfterTime.Day(),
		publishedAfterTime.Hour(), publishedAfterTime.Minute(), publishedAfterTime.Second())
	url := fmt.Sprintf("https://youtube.googleapis.com/youtube/v3/search?part=snippet&maxResults=50&key=%v&"+
		"type=video&publishedAfter=%v&q=surfing&order=date",
		"", publishedAfterFormatTime)
	if nextPageToken != "" {
		url += fmt.Sprintf("&pageToken=%v", nextPageToken)
	}
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	var result contract.APIData

	if err != nil {
		fmt.Println(err)
		return nil
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer func() {
		_ = res.Body.Close()
	}()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if res.StatusCode == 403 {
		fmt.Println("quota error: " + string(body))
		return nil
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("json unmarshal error: ", err)
		return nil
	}
	return &result
}
