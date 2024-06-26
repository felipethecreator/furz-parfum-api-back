package comp

type item struct {
	Type      string  `json:"type" bson:"type"`
	Value     float64 `json:"value" bson:"value"`
	Quality   int32   `json:"quality" bson:"quality"`
	Longevity string  `json:"longevity" bson:"longevity"`
	Year      int32   `json:"year" bson:"year"`
}
