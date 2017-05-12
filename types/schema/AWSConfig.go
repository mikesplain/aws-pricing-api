package schema

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/jinzhu/gorm"
)

type rawAWSConfig struct {
	FormatVersion	string
	Disclaimer	string
	OfferCode	string
	Version		string
	PublicationDate	string
	Products	map[string]AWSConfig_Product
	Terms		map[string]map[string]map[string]rawAWSConfig_Term
}


type rawAWSConfig_Term struct {
	OfferTermCode string
	Sku	string
	EffectiveDate string
	PriceDimensions map[string]AWSConfig_Term_PriceDimensions
	TermAttributes map[string]string
}

func (l *AWSConfig) UnmarshalJSON(data []byte) error {
	var p rawAWSConfig
	err := json.Unmarshal(data, p)
	if err != nil {
		return err
	}

	products := []AWSConfig_Product{}
	terms := []AWSConfig_Term{}

	// Convert from map to slice
	for _, pr := range p.Products {
		products = append(products, pr)
	}

	for _, tenancy := range p.Terms {
		// OnDemand, etc.
		for _, sku := range tenancy {
			// Some junk SKU
			for _, term := range sku {
				pDimensions := []AWSConfig_Term_PriceDimensions{}
				tAttributes := []AWSConfig_Term_Attributes{}

				for _, pd := range term.PriceDimensions {
					pDimensions = append(pDimensions, pd)
				}

				for key, value := range term.TermAttributes {
					tr := AWSConfig_Term_Attributes{
						Key: key,
						Value: value,
					}
					tAttributes = append(tAttributes, tr)
				}

				t := AWSConfig_Term{
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

type AWSConfig struct {
	gorm.Model
	FormatVersion	string
	Disclaimer	string
	OfferCode	string
	Version		string
	PublicationDate	string
	Products	[]AWSConfig_Product 	`gorm:"ForeignKey:ID,type:varchar(255)[]"`
	Terms		[]AWSConfig_Term	`gorm:"ForeignKey:ID,type:varchar(255)[]"`
}
type AWSConfig_Product struct {
	gorm.Model
		Sku	string
	ProductFamily	string
	Attributes	AWSConfig_Product_Attributes	`gorm:"ForeignKey:ID,type:varchar(255)[]"`
}
type AWSConfig_Product_Attributes struct {
	gorm.Model
		Servicecode	string
	Location	string
	LocationType	string
	Usagetype	string
	Operation	string
}

type AWSConfig_Term struct {
	gorm.Model
	OfferTermCode string
	Sku	string
	EffectiveDate string
	PriceDimensions []AWSConfig_Term_PriceDimensions 	`gorm:"ForeignKey:ID,type:varchar(255)[]"`
	TermAttributes []AWSConfig_Term_Attributes 	`gorm:"ForeignKey:ID,type:varchar(255)[]"`
}

type AWSConfig_Term_Attributes struct {
	gorm.Model
	Key	string
	Value	string
}

type AWSConfig_Term_PriceDimensions struct {
	gorm.Model
	RateCode	string
	RateType	string
	Description	string
	BeginRange	string
	EndRange	string
	Unit	string
	PricePerUnit	AWSConfig_Term_PricePerUnit 	`gorm:"ForeignKey:ID,type:varchar(255)[]"`
	AppliesTo	[]interface{}
}

type AWSConfig_Term_PricePerUnit struct {
	gorm.Model
	USD	string
}
func (a AWSConfig) QueryProducts(q func(product AWSConfig_Product) bool) []AWSConfig_Product{
	ret := []AWSConfig_Product{}
	for _, v := range a.Products {
		if q(v) {
			ret = append(ret, v)
		}
	}

	return ret
}
func (a AWSConfig) QueryTerms(t string, q func(product AWSConfig_Term) bool) []AWSConfig_Term{
	ret := []AWSConfig_Term{}
	for _, v := range a.Terms {
		if q(v) {
			ret = append(ret, v)
		}
	}

	return ret
}
func (a *AWSConfig) Refresh() error {
	var url = "https://pricing.us-east-1.amazonaws.com/offers/v1.0/aws/AWSConfig/current/index.json"
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