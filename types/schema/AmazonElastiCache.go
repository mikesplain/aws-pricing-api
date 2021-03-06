package schema

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
)

type rawAmazonElastiCache struct {
	FormatVersion   string
	Disclaimer      string
	OfferCode       string
	Version         string
	PublicationDate string
	Products        map[string]AmazonElastiCache_Product
	Terms           map[string]map[string]map[string]rawAmazonElastiCache_Term
}

type rawAmazonElastiCache_Term struct {
	OfferTermCode   string
	Sku             string
	EffectiveDate   string
	PriceDimensions map[string]AmazonElastiCache_Term_PriceDimensions
	TermAttributes  map[string]string
}

func (l *AmazonElastiCache) UnmarshalJSON(data []byte) error {
	var p rawAmazonElastiCache
	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}

	products := []*AmazonElastiCache_Product{}
	terms := []*AmazonElastiCache_Term{}

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
				pDimensions := []*AmazonElastiCache_Term_PriceDimensions{}
				tAttributes := []*AmazonElastiCache_Term_Attributes{}

				for _, pd := range term.PriceDimensions {
					pDimensions = append(pDimensions, &pd)
				}

				for key, value := range term.TermAttributes {
					tr := AmazonElastiCache_Term_Attributes{
						Key:   key,
						Value: value,
					}
					tAttributes = append(tAttributes, &tr)
				}

				t := AmazonElastiCache_Term{
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

type AmazonElastiCache struct {
	gorm.Model
	FormatVersion   string
	Disclaimer      string
	OfferCode       string
	Version         string
	PublicationDate string
	Products        []*AmazonElastiCache_Product `gorm:"ForeignKey:AmazonElastiCacheID"`
	Terms           []*AmazonElastiCache_Term    `gorm:"ForeignKey:AmazonElastiCacheID"`
}
type AmazonElastiCache_Product struct {
	gorm.Model
	AmazonElastiCacheID uint
	Attributes          AmazonElastiCache_Product_Attributes `gorm:"ForeignKey:AmazonElastiCache_Product_AttributesID"`
	Sku                 string
	ProductFamily       string
}
type AmazonElastiCache_Product_Attributes struct {
	gorm.Model
	AmazonElastiCache_Product_AttributesID uint
	Vcpu                                   string
	Memory                                 string
	NetworkPerformance                     string
	Location                               string
	LocationType                           string
	InstanceType                           string
	CurrentGeneration                      string
	InstanceFamily                         string
	Operation                              string
	Servicename                            string
	Servicecode                            string
	CacheEngine                            string
	Usagetype                              string
}

type AmazonElastiCache_Term struct {
	gorm.Model
	OfferTermCode       string
	AmazonElastiCacheID uint
	Sku                 string
	EffectiveDate       string
	PriceDimensions     []*AmazonElastiCache_Term_PriceDimensions `gorm:"ForeignKey:AmazonElastiCache_TermID"`
	TermAttributes      []*AmazonElastiCache_Term_Attributes      `gorm:"ForeignKey:AmazonElastiCache_TermID"`
}

type AmazonElastiCache_Term_Attributes struct {
	gorm.Model
	AmazonElastiCache_TermID uint
	Key                      string
	Value                    string
}

type AmazonElastiCache_Term_PriceDimensions struct {
	gorm.Model
	AmazonElastiCache_TermID uint
	RateCode                 string
	RateType                 string
	Description              string
	BeginRange               string
	EndRange                 string
	Unit                     string
	PricePerUnit             *AmazonElastiCache_Term_PricePerUnit `gorm:"ForeignKey:AmazonElastiCache_Term_PriceDimensionsID"`
	// AppliesTo	[]string
}

type AmazonElastiCache_Term_PricePerUnit struct {
	gorm.Model
	AmazonElastiCache_Term_PriceDimensionsID uint
	USD                                      string
}

func (a *AmazonElastiCache) Refresh() error {
	var url = "https://pricing.us-east-1.amazonaws.com/offers/v1.0/aws/AmazonElastiCache/current/index.json"
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
