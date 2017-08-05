package resources

import (
	"fmt"
	"github.com/sreeram-boyapati/razorpay-go/razorpay/constants"
	"net/http"
)

type Card struct {
	base BaseResource
}

func (card *Card) fetch(
	card_id string,
	data map[string]interface{},
	options map[string]string) (*http.Response, error) {

	url := fmt.Sprintf("%s/%s", constants.CARD_URL, card_id)

	return card.base.getUrl(url, data, options)
}
