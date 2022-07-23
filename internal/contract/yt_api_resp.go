package contract

import "time"

type APIData struct {
	Kind          string                 `json:"kind"`
	Etag          string                 `json:"etag"`
	NextPageToken string                 `json:"nextPageToken"`
	PrevPageToken string                 `json:"prevPageToken,omitempty"`
	RegionCode    string                 `json:"regionCode"`
	PageInfo      map[string]interface{} `json:"pageInfo"`
	Items         []*VideoItem           `json:"items"`
}

type VideoItem struct {
	Kind    string        `json:"kind"`
	Etag    string        `json:"etag"`
	ID      *ItemID       `json:"id"`
	Snippet *VideoSnippet `json:"snippet"`
}

type VideoSnippet struct {
	PublishedAt          time.Time        `json:"publishedAt"`
	ChannelID            string           `json:"channelId"`
	Title                string           `json:"title"`
	Description          string           `json:"description"`
	Thumbnails           *VideoThumbnails `json:"thumbnails"`
	ChannelTitle         string           `json:"channelTitle"`
	LiveBroadcastContest interface{}      `json:"liveBroadcastContest"`
	PublishTime          time.Time        `json:"publishTime"`
}

type VideoThumbnails struct {
	Default *Thumbnail `json:"default"`
	Medium  *Thumbnail `json:"medium"`
	High    *Thumbnail `json:"high"`
}

type Thumbnail struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

type ItemID struct {
	Kind    string `json:"kind"`
	VideoID string `json:"VideoId"`
}
