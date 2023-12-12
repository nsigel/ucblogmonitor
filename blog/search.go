package blog

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func (c *Client) Search() (*Post, error) {
	req, err := http.NewRequest(http.MethodGet, "https://collegeadmissions.uchicago.edu/uncommon-blog", nil)
	if err != nil {
		return nil, err
	}

	res, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var post Post

	postSelection := document.Find(".views-row").First()

	post.Title = postSelection.Find(".post-title").Text()
	post.Content = postSelection.Find(".summary").Text()
	post.Byline = postSelection.Find(".author-date").Text()
	post.Url, _ = postSelection.Find(".post-title a").Attr("href")

	return &post, nil
}

type Post struct {
	Content string
	Byline  string
	Title   string
	Url     string
}
