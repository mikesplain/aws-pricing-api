package schema

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/jinzhu/gorm"
)

type rawAmazonCloudDirectory struct {
	FormatVersion	string
	Disclaimer	string
	OfferCode	string
	Version		string
	PublicationDate	string
	Products	map[string]AmazonCloudDirectory_Product
	Terms		map[string]map[string]map[string]rawAmazonCloudDirectory_Term
}


type rawAmazonCloudDirectory_Term struct {
	OfferTermCode string
	Sku	string
	EffectiveDate string
	PriceDimensions map[string]AmazonCloudDirectory_Term_PriceDimensions
	TermAttributes map[string]string
}

func (l *AmazonCloudDirectory) UnmarshalJSON(data []byte) error {
	var p rawAmazonCloudDirectory
	err := json.Unmarshal(data, p)
	if err != nil {
		return err
	}

	products := []AmazonCloudDirectory_Product{}
	terms := []AmazonCloudDirectory_Term{}

	// Convert from map to slice
	for _, pr := range p.Products {
		products = append(products, pr)
	}

	for _, tenancy := range p.Terms {
		// OnDemand, etc.
		for _, sku := range tenancy {
			// Some junk SKU
			for _, term := range sku {
				pDimensions := []AmazonCloudDirectory_Term_PriceDimensions{}
				tAttributes := []AmazonCloudDirectory_Term_Attributes{}

				for _, pd := range term.PriceDimensions {
					pDimensions = append(pDimensions, pd)
				}

				for key, value := range term.TermAttributes {
					tr := AmazonCloudDirectory_Term_Attributes{
						Key: key,
						Value: value,
					}
					tAttributes = append(tAttributes, tr)
				}

				t := AmazonCloudDirectory_Term{
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

type AmazonCloudDirectory struct {
	gorm.Model
	FormatVersion	string
	Disclaimer	string
	OfferCode	string
	Version		string
	PublicationDate	string
	Products	[]AmazonCloudDirectory_Product 	`gorm:"ForeignKey:ID,type:varchar(255)[]"`
	Terms		[]AmazonCloudDirectory_Term	`gorm:"ForeignKey:ID,type:varchar(255)[]"`
}
type AmazonCloudDirectory_Product struct {
	gorm.Model
		Sku	string
	ProductFamily	string
	Attributes	AmazonCloudDirectory_Product_Attributes	`gorm:"ForeignKey:ID,type:varchar(255)[]"`
}
type AmazonCloudDirectory_Product_Attributes struct {
	gorm.Model
		Location	string
	LocationType	string
	Group	string
	GroupDescription	string
	Usagetype	string
	Operation	string
	Servicecode	string
}

type AmazonCloudDirectory_Term struct {
	gorm.Model
	OfferTermCode string
	Sku	string
	EffectiveDate string
	PriceDimensions []AmazonCloudDirectory_Term_PriceDimensions 	`gorm:"ForeignKey:ID,type:varchar(255)[]"`
	TermAttributes []AmazonCloudDirectory_Term_Attributes 	`gorm:"ForeignKey:ID,type:varchar(255)[]"`
}

type AmazonCloudDirectory_Term_Attributes struct {
	gorm.Model
	Key	string
	Value	string
}

type AmazonCloudDirectory_Term_PriceDimensions struct {
	gorm.Model
	RateCode	string
	RateType	string
	Description	string
	BeginRange	string
	EndRange	string
	Unit	string
	PricePerUnit	AmazonCloudDirectory_Term_PricePerUnit 	`gorm:"ForeignKey:ID,type:varchar(255)[]"`
	AppliesTo	[]interface{}
}

type AmazonCloudDirectory_Term_PricePerUnit struct {
	gorm.Model
	USD	string
}
func (a AmazonCloudDirectory) QueryProducts(q func(product AmazonCloudDirectory_Product) bool) []AmazonCloudDirectory_Product{
	ret := []AmazonCloudDirectory_Product{}
	for _, v := range a.Products {
		if q(v) {
			ret = append(ret, v)
		}
	}

	return ret
}
func (a AmazonCloudDirectory) QueryTerms(t string, q func(product AmazonCloudDirectory_Term) bool) []AmazonCloudDirectory_Term{
	ret := []AmazonCloudDirectory_Term{}
	for _, v := range a.Terms {
		if q(v) {
			ret = append(ret, v)
		}
	}

	return ret
}
func (a *AmazonCloudDirectory) Refresh() error {
	var url = "https://pricing.us-east-1.amazonaws.com/offers/v1.0/aws/AmazonCloudDirectory/current/index.json"
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