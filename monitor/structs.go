package monitor

type Region string
const (
	EUROPE = "EU"
	UNITED_STATES = "US"
	CANADA = "CA"
	GERMANY = "DE"
	UNITED_KINGDOM = "UK"
)

type Stores string
const (
	Amazon 		= "AMAZON"
	Target 		= "TARGET"
	Walmart		= "WALMART"
	BestBuy 	= "BESTBUY"
	NewEgg 		= "NEWEGG"
	Lego 		= "LEGO"
)

type BaseProduct struct {
	Title    string `json:"title" bson:"title"`
	Image    string `json:"image" bson:"image"`
	Link     string `json:"link"  bson:"link"`
	Price    string `json:"price" bson:"price"`
	SKU      string `json:"sku"   bson:"sku"`
	Store    Stores `json:"store" bson:"store"`
	StoreURL string `json:"store_url" bson:"store_url"`
	Region   Region `json:"region" bson:"region"`

	//stuff for Amazon
	OfferID  	string `json:"offer_id" bson:"offer_id"`

	AvailableQuantity int `json:"available_quantity" bson:"-"`

	//stuff for Microsoft
	AvailabilityId string `json:"availability_id" bson:"availability_id"`
	SKUId		   string `json:"sku_id" bson:"sku_id"`
}