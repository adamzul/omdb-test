package objects

type OmdbAPIMovieObject struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type OmdbAPIListResponseObject struct {
	Search       []OmdbAPIMovieObject `json:"Search"`
	TotalResults string               `json:"totalResults"`
	Response     string               `json:"Response"`
}

type OmdbAPIListRequestObject struct {
	Search string `json:"search"`
	Page   int    `json:"page"`
}
