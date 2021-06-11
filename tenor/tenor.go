package tenor

import (
	"fmt"
	"regexp"
	"shitposter-bot/shared"

	gotenor "github.com/jkershaw2000/go-tenor"
)

var reg = regexp.MustCompile(`(\d{2,})`)
var client *gotenor.Tenor

func Start(token string) {
	client = gotenor.NewTenor(token)
	fmt.Println("Created Tenor")
}

func GetGIFbyURL(url string) (string, bool) {
	fmt.Println("Getting tenor GIF for", url)
	id := reg.FindString(url)
	if len(id) > 0 {
		res, err := client.GetById(id)
		if shared.CheckError(err) || len(res.Results) < 1 {
			return "", false
		}

		return gotenor.GetGifURL(*res), true
	}

	return "", false
}
