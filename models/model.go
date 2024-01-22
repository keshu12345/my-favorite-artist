package models

const artistInfo = "https://ws.audioscrobbler.com/2.0/?method=artist.getinfo&artist=Cher&api_key=2323785c019f7c0509019a8723c35be2&format=json"

type Streamable struct {
	Text      string `json:"#text"`
	Fulltrack string `json:"fulltrack"`
}

type Artist struct {
	Name string `json:"name"`
	Mbid string `json:"mbid"`
	URL  string `json:"url"`
}

type Image struct {
	Text string `json:"#text"`
	Size string `json:"size"`
}

type Track struct {
	Name       string     `json:"name"`
	Duration   string     `json:"duration"`
	Listeners  string     `json:"listeners"`
	Mbid       string     `json:"mbid"`
	URL        string     `json:"url"`
	Streamable Streamable `json:"streamable"`
	Artist     Artist     `json:"artist"`
	Image      []Image    `json:"image"`
	Attr       struct {
		Rank string `json:"rank"`
	} `json:"@attr"`
}

type Tracks struct {
	Track []Track `json:"track"`
}

type TopTracks struct {
	Tracks Tracks `json:"tracks"`
	Attr   struct {
		Country    string `json:"country"`
		Page       string `json:"page"`
		PerPage    string `json:"perPage"`
		TotalPages string `json:"totalPages"`
		Total      string `json:"total"`
	} `json:"@attr"`
}

type ArtistInfo struct {
	Artist struct {
		Name  string `json:"name"`
		MBID  string `json:"mbid"`
		URL   string `json:"url"`
		Image []struct {
			Text string `json:"#text"`
			Size string `json:"size"`
		} `json:"image"`
		Streamable string `json:"streamable"`
		Ontour     string `json:"ontour"`
		Stats      struct {
			Listeners string `json:"listeners"`
			Playcount string `json:"playcount"`
		} `json:"stats"`
		Similar struct {
			Artist []struct {
				Name  string `json:"name"`
				URL   string `json:"url"`
				Image []struct {
					Text string `json:"#text"`
					Size string `json:"size"`
				} `json:"image"`
			} `json:"artist"`
		} `json:"similar"`
	} `json:"artist"`
	Tags struct {
		Tag []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"tag"`
	} `json:"tags"`
	Bio struct {
		Links struct {
			Link struct {
				Text string `json:"#text"`
				Rel  string `json:"rel"`
				Href string `json:"href"`
			} `json:"link"`
		} `json:"links"`
		Published string `json:"published"`
		Summary   string `json:"summary"`
		Content   string `json:"content"`
	} `json:"bio"`
}

type ImageWithImageURL struct {
	Image []Image `json:"image"`
	URL   string
}
