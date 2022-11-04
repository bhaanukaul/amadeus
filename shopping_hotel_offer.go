package amadeus

import "time"

type HotelData struct {
	Type      string        `json:"type,omitempty"`
	Self      string        `json:"self,omitempty"`
	Offers    []HotelOffers `json:"offers,omitempty"`
	Hotel     Hotel         `json:"hotel,omitempty"`
	Available bool          `json:"available,omitempty"`
}

// GetOfferID return offerID from list
func (dR HotelData) GetOfferID(offerNum int) string {
	return dR.Offers[offerNum].ID
}

type Hotel struct {
	Address       HotelAddress  `json:"address,omitempty"`
	Contact       HotelContact  `json:"contact,omitempty"`
	Description   Description   `json:"description,omitempty"`
	HotelDistance HotelDistance `json:"hotelDistance,omitempty"`
	HotelID       string        `json:"hotelId,omitempty"`
	ChainCode     string        `json:"chainCode,omitempty"`
	BrandCode     string        `json:"brandCode,omitempty"`
	DupeID        string        `json:"dupeId,omitempty"`
	Name          string        `json:"name,omitempty"`
	Rating        string        `json:"rating,omitempty"`
	Type          string        `json:"type,omitempty"`
	CityCode      string        `json:"cityCode,omitempty"`
	Amenities     []string      `json:"amenities,omitempty"`
	Media         []Media       `json:"media,omitempty"`
	Longitude     float64       `json:"longitude,omitempty"`
	Latitude      float64       `json:"latitude,omitempty"`
}

type Description struct {
	Lang string `json:"lang,omitempty"`
	Text string `json:"text,omitempty"`
}
type Media struct {
	URI      string `json:"uri,omitempty"`
	Category string `json:"category,omitempty"`
}
type HotelDistance struct {
	DistanceUnit string  `json:"distanceUnit,omitempty"`
	Distance     float64 `json:"distance,omitempty"`
}
type HotelAddress struct {
	PostalCode  string   `json:"postalCode,omitempty"`
	CityName    string   `json:"cityName,omitempty"`
	CountryCode string   `json:"countryCode,omitempty"`
	StateCode   string   `json:"stateCode,omitempty"`
	Lines       []string `json:"lines,omitempty"`
}
type HotelContact struct {
	Phone string `json:"phone,omitempty"`
	Fax   string `json:"fax,omitempty"`
	Email string `json:"email,omitempty"`
}

type HotelOffers struct {
	Policies            Policies            `json:"policies,omitempty"`
	Room                Room                `json:"room,omitempty"`
	Commission          Commission          `json:"commission,omitempty"`
	RateFamilyEstimated RateFamilyEstimated `json:"rateFamilyEstimated,omitempty"`
	Description         Description         `json:"description,omitempty"`
	Self                string              `json:"self,omitempty"`
	RateCode            string              `json:"rateCode,omitempty"`
	CheckOutDate        string              `json:"checkOutDate,omitempty"`
	Category            string              `json:"category,omitempty"`
	CheckInDate         string              `json:"checkInDate,omitempty"`
	BoardType           string              `json:"boardType,omitempty"`
	ID                  string              `json:"id,omitempty"`
	Type                string              `json:"type,omitempty"`
	Price               HotelPrice          `json:"price,omitempty"`
	Guests              Guests              `json:"guests,omitempty"`
	RoomQuantity        int                 `json:"roomQuantity,omitempty"`
}

type RateFamilyEstimated struct {
	Code string `json:"code,omitempty"`
	Type string `json:"type,omitempty"`
}
type Commission struct {
	//Percentage  float64     `json:"percentage,omitempty"`
	Amount      string      `json:"amount,omitempty"`
	Description Description `json:"description,omitempty"`
}

type Room struct {
	Description   Description   `json:"description,omitempty"`
	Type          string        `json:"type,omitempty"`
	TypeEstimated TypeEstimated `json:"typeEstimated,omitempty"`
}

type TypeEstimated struct {
	Category string `json:"category,omitempty"`
	BedType  string `json:"bedType,omitempty"`
	Beds     int    `json:"beds,omitempty"`
}

type Guests struct {
	ChildAges []int `json:"childAges,omitempty"`
	Adults    int   `json:"adults,omitempty"`
}

type HotelPrice struct {
	Currency   string            `json:"currency,omitempty"`
	Total      string            `json:"total,omitempty"`
	Base       string            `json:"base,omitempty"`
	Taxes      []HotelPriceTaxes `json:"taxes,omitempty"`
	Variations Variations        `json:"variations,omitempty"`
}

type HotelPriceTaxes struct {
	Currency         string `json:"currency,omitempty"`
	Amount           string `json:"amount,omitempty"`
	Code             string `json:"code,omitempty"`
	Description      string `json:"description,omitempty"`
	PricingFrequency string `json:"pricingFrequency,omitempty"`
	PricingMode      string `json:"pricingMode,omitempty"`
	Included         bool   `json:"included,omitempty"`
}

type Variations struct {
	Average Average   `json:"average,omitempty"`
	Changes []Changes `json:"changes,omitempty"`
}

type Average struct {
	Base  string `json:"base,omitempty"`
	Total string `json:"total,omitempty"`
}

type Changes struct {
	StartDate string `json:"startDate,omitempty"`
	EndDate   string `json:"endDate,omitempty"`
	Base      string `json:"base,omitempty"`
	Total     string `json:"total,omitempty"`
}

type Policies struct {
	CheckInOut   CheckInOut   `json:"checkInOut,omitempty"`
	Cancellation Cancellation `json:"cancellation,omitempty"`
	PaymentType  string       `json:"paymentType,omitempty"`
	HoldTime     HoldTime     `json:"holdTime,omitempty"`
	Deposit      Deposit      `json:"deposit,omitempty"`
	Prepay       Prepay       `json:"prepay,omitempty"`
	Guarantee    Guarantee    `json:"guarantee,omitempty"`
}

type Guarantee struct {
	Description      Description      `json:"description,omitempty"`
	AcceptedPayments AcceptedPayments `json:"acceptedPayments,omitempty"`
}
type AcceptedPayments struct {
	CreditCards []string `json:"creditCards,omitempty"`
	Methods     []string `json:"methods,omitempty"`
}

type Deposit struct {
	Amount string `json:"amount,omitempty"`
	//Deadline         time.Time        `json:"deadline,omitempty"`
	Deadline         string           `json:"deadline,omitempty"`
	Description      Description      `json:"description,omitempty"`
	AcceptedPayments AcceptedPayments `json:"acceptedPayments,omitempty"`
}
type Prepay struct {
	Amount           string           `json:"amount,omitempty"`
	Deadline         time.Time        `json:"deadline,omitempty"`
	Description      Description      `json:"description,omitempty"`
	AcceptedPayments AcceptedPayments `json:"acceptedPayments,omitempty"`
}
type HoldTime struct {
	//Deadline time.Time `json:"deadline,omitempty"`
	Deadline string `json:"deadline,omitempty"`
}
type Cancellation struct {
	Deadline       time.Time   `json:"deadline,omitempty"`
	Description    Description `json:"description,omitempty"`
	Type           string      `json:"type,omitempty"`
	Amount         string      `json:"amount,omitempty"`
	Percentage     string      `json:"percentage,omitempty"`
	NumberOfNights int         `json:"numberOfNights,omitempty"`
}

type CheckInOut struct {
	CheckIn             string              `json:"checkIn,omitempty"`
	CheckInDescription  CheckInDescription  `json:"checkInDescription,omitempty"`
	CheckOut            string              `json:"checkOut,omitempty"`
	CheckOutDescription CheckOutDescription `json:"checkOutDescription,omitempty"`
}

type CheckInDescription struct {
	Lang string `json:"lang,omitempty"`
	Text string `json:"text,omitempty"`
}
type CheckOutDescription struct {
	Lang string `json:"lang,omitempty"`
	Text string `json:"text,omitempty"`
}
