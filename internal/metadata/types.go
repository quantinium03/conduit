package metadata

type Versions struct {
    Info int32 `json:"info"`
    Extract int32 `json:"extract"`
    Thumbnail int32 `json:"thumbnail"`
    Keyframes int32 `json:"keyframes"`
}

type Video struct {
    Index uint32 `json:"index"`
    Title *string `json:"title"`
    Language *string `json:"language"`
    Codec string `json:"codec"`
    MimeCodec *string `json:"mimeCodec"`
    Width uint32 `json:"width"`
    Height uint32 `json:"height"`
    Bitrate uint32 `json:"bitrate"`
    IsDefault bool `json:"IsDefault"`
    Keyframes *Keyframe `json:"-"`
};

type MediaInfo struct {
    Sha string `json:"sha"` // sha of the file
    Path string `json:"path"` // full path of the file
    Extension string `json:"extension"` // extensinon of the file
    MimeCodec *string `json:"mimeCoded"` // mimeype (defined as RFC 6301) ex: video/mp4
    Size uint64 `json:"size"`
    Duration float64 `json:"duration"`
    Container *string `json:"container"`
    Versions Versions `json:"versions"`
    Video Video `json:"video"`
}
