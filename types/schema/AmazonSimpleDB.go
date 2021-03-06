package schema

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
)

type rawAmazonSimpleDB struct {
	FormatVersion   string
	Disclaimer      string
	OfferCode       string
	Version         string
	PublicationDate string
	Products        map[string]AmazonSimpleDB_Product
	Terms           map[string]map[string]map[string]rawAmazonSimpleDB_Term
}

type rawAmazonSimpleDB_Term struct {
	OfferTermCode   string
	Sku             string
	EffectiveDate   string
	PriceDimensions map[string]AmazonSimpleDB_Term_PriceDimensions
	TermAttributes  map[string]string
}

func (l *AmazonSimpleDB) UnmarshalJSON(data []byte) error {
	var p rawAmazonSimpleDB
	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}

	products := []*AmazonSimpleDB_Product{}
	terms := []*AmazonSimpleDB_Term{}

	// Convert from map to slice
	for i, _ := range p.Products {
		pr := p.Products[i]
		products = append(products, &pr)
	}

	for _, tenancy := range p.Terms {
		// OnDemand, etc.
		for _, sku := range tenancy {
			// Some junk SKU
			for _, term := range sku {
				pDimensions := []*AmazonSimpleDB_Term_PriceDimensions{}
				tAttributes := []*AmazonSimpleDB_Term_Attributes{}

				for _, pd := range term.PriceDimensions {
					pDimensions = append(pDimensions, &pd)
				}

				for key, value := range term.TermAttributes {
					tr := AmazonSimpleDB_Term_Attributes{
						Key:   key,
						Value: value,
					}
					tAttributes = append(tAttributes, &tr)
				}

				t := AmazonSimpleDB_Term{
					OfferTermCode:   term.OfferTermCode,
					Sku:             term.Sku,
					EffectiveDate:   term.EffectiveDate,
					TermAttributes:  tAttributes,
					PriceDimensions: pDimensions,
				}

				terms = append(terms, &t)
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

type AmazonSimpleDB struct {
	gorm.Model
	FormatVersion   string
	Disclaimer      string
	OfferCode       string
	Version         string
	PublicationDate string
	Products        []*AmazonSimpleDB_Product `gorm:"ForeignKey:AmazonSimpleDBID"`
	Terms           []*AmazonSimpleDB_Term    `gorm:"ForeignKey:AmazonSimpleDBID"`
}
type AmazonSimpleDB_Product struct {
	gorm.Model
	AmazonSimpleDBID uint
	Sku              string
	ProductFamily    string
	Attributes       AmazonSimpleDB_Product_Attributes `gorm:"ForeignKey:AmazonSimpleDB_Product_AttributesID"`
}
type AmazonSimpleDB_Product_Attributes struct {
	gorm.Model
	AmazonSimpleDB_Product_AttributesID uint
	ToLocation                          string
	ToLocationType                      string
	Usagetype                           string
	Operation                           string
	Servicecode                         string
	TransferType                        string
	FromLocation                        string
	FromLocationType                    string
}

type AmazonSimpleDB_Term struct {
	gorm.Model
	OfferTermCode    string
	AmazonSimpleDBID uint
	Sku              string
	EffectiveDate    string
	PriceDimensions  []*AmazonSimpleDB_Term_PriceDimensions `gorm:"ForeignKey:AmazonSimpleDB_TermID"`
	TermAttributes   []*AmazonSimpleDB_Term_Attributes      `gorm:"ForeignKey:AmazonSimpleDB_TermID"`
}

type AmazonSimpleDB_Term_Attributes struct {
	gorm.Model
	AmazonSimpleDB_TermID uint
	Key                   string
	Value                 string
}

type AmazonSimpleDB_Term_PriceDimensions struct {
	gorm.Model
	AmazonSimpleDB_TermID uint
	RateCode              string
	RateType              string
	Description           string
	BeginRange            string
	EndRange              string
	Unit                  string
	PricePerUnit          *AmazonSimpleDB_Term_PricePerUnit `gorm:"ForeignKey:AmazonSimpleDB_Term_PriceDimensionsID"`
	// AppliesTo	[]string
}

type AmazonSimpleDB_Term_PricePerUnit struct {
	gorm.Model
	AmazonSimpleDB_Term_PriceDimensionsID uint
	USD                                   string
}

func (a *AmazonSimpleDB) Refresh() error {
	var url = "https://pricing.us-east-1.amazonaws.com/offers/v1.0/aws/AmazonSimpleDB/current/index.json"
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
