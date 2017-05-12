package schema

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/jinzhu/gorm"
)

type rawAmazonQuickSight struct {
	FormatVersion	string
	Disclaimer	string
	OfferCode	string
	Version		string
	PublicationDate	string
	Products	map[string]AmazonQuickSight_Product
	Terms		map[string]map[string]map[string]rawAmazonQuickSight_Term
}


type rawAmazonQuickSight_Term struct {
	OfferTermCode string
	Sku	string
	EffectiveDate string
	PriceDimensions map[string]AmazonQuickSight_Term_PriceDimensions
	TermAttributes map[string]string
}

func (l *AmazonQuickSight) UnmarshalJSON(data []byte) error {
	var p rawAmazonQuickSight
	err := json.Unmarshal(data, p)
	if err != nil {
		return err
	}

	products := []AmazonQuickSight_Product{}
	terms := []AmazonQuickSight_Term{}

	// Convert from map to slice
	for _, pr := range p.Products {
		products = append(products, pr)
	}

	for _, tenancy := range p.Terms {
		// OnDemand, etc.
		for _, sku := range tenancy {
			// Some junk SKU
			for _, term := range sku {
				pDimensions := []AmazonQuickSight_Term_PriceDimensions{}
				tAttributes := []AmazonQuickSight_Term_Attributes{}

				for _, pd := range term.PriceDimensions {
					pDimensions = append(pDimensions, pd)
				}

				for key, value := range term.TermAttributes {
					tr := AmazonQuickSight_Term_Attributes{
						Key: key,
						Value: value,
					}
					tAttributes = append(tAttributes, tr)
				}

				t := AmazonQuickSight_Term{
					OfferTermCode: term.OfferTermCode,
					Sku: term.Sku,
					EffectiveDate: term.EffectiveDate,
					TermAttributes: tAttributes,
					PriceDimensions: pDimensions,
				}

				terms = append(terms, t)
			}
		}
	}

	l.FormatVersion = p.FormatVersion
	l.Disclaimer = p.Disclaimer
	l.OfferCode = p.OfferCode
	l.Version = p.Version
	l.PublicationDate = p.PublicationDate
	l.Products = products
	l.Terms = terms
	return nil
}

type AmazonQuickSight struct {
	gorm.Model
	FormatVersion	string
	Disclaimer	string
	OfferCode	string
	Version		string
	PublicationDate	string
	Products	[]AmazonQuickSight_Product 	`gorm:"ForeignKey:ID,type:varchar(255)[]"`
	Terms		[]AmazonQuickSight_Term	`gorm:"ForeignKey:ID,type:varchar(255)[]"`
}
type AmazonQuickSight_Product struct {
	gorm.Model
		Sku	string
	ProductFamily	string
	Attributes	AmazonQuickSight_Product_Attributes	`gorm:"ForeignKey:ID,type:varchar(255)[]"`
}
type AmazonQuickSight_Product_Attributes struct {
	gorm.Model
		Location	string
	Group	string
	GroupDescription	string
	Servicecode	string
	Usagetype	string
	Operation	string
	Edition	string
	SubscriptionType	string
	LocationType	string
}

type AmazonQuickSight_Term struct {
	gorm.Model
	OfferTermCode string
	Sku	string
	EffectiveDate string
	PriceDimensions []AmazonQuickSight_Term_PriceDimensions 	`gorm:"ForeignKey:ID,type:varchar(255)[]"`
	TermAttributes []AmazonQuickSight_Term_Attributes 	`gorm:"ForeignKey:ID,type:varchar(255)[]"`
}

type AmazonQuickSight_Term_Attributes struct {
	gorm.Model
	Key	string
	Value	string
}

type AmazonQuickSight_Term_PriceDimensions struct {
	gorm.Model
	RateCode	string
	RateType	string
	Description	string
	BeginRange	string
	EndRange	string
	Unit	string
	PricePerUnit	AmazonQuickSight_Term_PricePerUnit 	`gorm:"ForeignKey:ID,type:varchar(255)[]"`
	AppliesTo	[]interface{}
}

type AmazonQuickSight_Term_PricePerUnit struct {
	gorm.Model
	USD	string
}
func (a AmazonQuickSight) QueryProducts(q func(product AmazonQuickSight_Product) bool) []AmazonQuickSight_Product{
	ret := []AmazonQuickSight_Product{}
	for _, v := range a.Products {
		if q(v) {
			ret = append(ret, v)
		}
	}

	return ret
}
func (a AmazonQuickSight) QueryTerms(t string, q func(product AmazonQuickSight_Term) bool) []AmazonQuickSight_Term{
	ret := []AmazonQuickSight_Term{}
	for _, v := range a.Terms {
		if q(v) {
			ret = append(ret, v)
		}
	}

	return ret
}
func (a *AmazonQuickSight) Refresh() error {
	var url = "https://pricing.us-east-1.amazonaws.com/offers/v1.0/aws/AmazonQuickSight/current/index.json"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, a)
	if err != nil {
		return err
	}

	return nil
}