package amadeus

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type HotelView string

func (p HotelView) String() string {
	return string(p)
}

const (
	//HOTELVIEW_NONE geocoordinates, hotel distance
	HOTELVIEW_NONE HotelView = "NONE"
	//HOTELVIEW_LIGHT: NONE view + city name, phone number, fax, address, postal code, country code, state code, ratings, 1 image
	HOTELVIEW_LIGHT = "LIGHT"
	//HOTELVIEW_FULL LIGHT view + hotel description, amenities and facilities
	HOTELVIEW_FULL = "FULL"
)

type HotelSort string

func (p HotelSort) String() string {
	return string(p)
}

const (
	DISTANCE = "DISTANCE"
	PRICE    = "PRICE"
)

type HotelBoardType string

func (p HotelBoardType) String() string {
	return string(p)
}

const (
	ROOM_ONLY     = "RO"  // Room Only
	BREAKFAST     = "BB"  // Bed & Breakfast
	HALF_BOARD    = "DBB" // Diner & Bed & Breakfast (only for Aggregators)
	FULL_BOARD    = "FB"  // Full Board (only for Aggregators)
	ALL_INCLUSIVE = "AI"  // All Inclusive (only for Aggregators)
)

type HotelPaymentPolicy string

func (p HotelPaymentPolicy) String() string {
	return string(p)
}

const (
	DEPOSIT   = "DEPOSIT"
	GUARANTEE = "GUARANTEE"
)

type RadiusUnit string

func (r RadiusUnit) String() string {
	return string(r)
}

const (
	MILE = "MILE"
	KM   = "KM"
)

type ShoppingHotelOffersRequest struct {
	PriceRange    string             `json:"priceRange"`
	Currency      string             `json:"currency"`
	RadiusUnit    RadiusUnit         `json:"radiusUnit"`
	CheckInDate   string             `json:"checkInDate"`
	CheckOutDate  string             `json:"checkOutDate"`
	Sort          HotelSort          `json:"sort"`
	View          HotelView          `json:"view"`
	PaymentPolicy HotelPaymentPolicy `json:"paymentPolicy"`
	HotelName     string             `json:"hotelName"`
	HotelID       string             `json:"hotelId,omitempty"`
	OfferID       string             `json:"offerId,omitempty"`
	BoardType     HotelBoardType     `json:"boardType"`
	CityCode      string             `json:"cityCode,omitempty"`
	Lang          string             `json:"lang"`
	Chains        []string           `json:"chains"`
	RateCodes     []string           `json:"rateCodes"`
	Amenities     []string           `json:"amenities"`
	Ratings       []int              `json:"ratings"`
	ChildAges     []int              `json:"childAges"`
	HotelIDs      []string           `json:"hotelIds"`
	Longitude     float64            `json:"longitude"`
	Radius        int                `json:"radius"`
	Latitude      float64            `json:"latitude"`
	Adults        int                `json:"adults"`
	RoomQuantity  int                `json:"roomQuantity"`
	BestRateOnly  bool               `json:"bestRateOnly"`
	IncludeClosed bool               `json:"includeClosed"`
}

// SetOfferID set offer id
func (sR *ShoppingHotelOffersRequest) SetOfferID(offerID string) *ShoppingHotelOffersRequest {

	sR.OfferID = offerID

	return sR
}

// SetHotelID set hotel id
func (sR *ShoppingHotelOffersRequest) SetHotelID(hotelID string) *ShoppingHotelOffersRequest {

	sR.HotelID = hotelID

	return sR
}

// SetCityCode set city code
func (sR *ShoppingHotelOffersRequest) SetCityCode(cityCode string) *ShoppingHotelOffersRequest {

	sR.CityCode = cityCode

	return sR
}

// SetCheckInDate set checkin date
func (sR *ShoppingHotelOffersRequest) SetCheckInDate(checkInDate string) *ShoppingHotelOffersRequest {

	// check date

	sR.CheckInDate = checkInDate

	return sR
}

// SetCheckOutDate set checkout date
func (sR *ShoppingHotelOffersRequest) SetCheckOutDate(checkOutDate string) *ShoppingHotelOffersRequest {

	// check date

	sR.CheckOutDate = checkOutDate

	return sR
}

// SetAdults set adults
func (sR *ShoppingHotelOffersRequest) SetAdults(adults int) *ShoppingHotelOffersRequest {

	// check date

	sR.Adults = adults

	return sR
}

// AddChildAges add childAges
func (sR *ShoppingHotelOffersRequest) AddChildAges(childAges ...int) *ShoppingHotelOffersRequest {

	sR.ChildAges = childAges

	return sR
}

// SetHotelName set hotel name
func (sR *ShoppingHotelOffersRequest) SetHotelName(hotelName string) *ShoppingHotelOffersRequest {

	sR.HotelName = hotelName

	return sR
}

// SetLatitude set latitude
func (sR *ShoppingHotelOffersRequest) SetLatitude(latitude float64) *ShoppingHotelOffersRequest {

	sR.Latitude = latitude

	return sR
}

// SetLongitude set longitude
func (sR *ShoppingHotelOffersRequest) SetLongitude(longitude float64) *ShoppingHotelOffersRequest {

	sR.Longitude = longitude

	return sR
}

// SetRadius set radius
func (sR *ShoppingHotelOffersRequest) SetRadius(radius int) *ShoppingHotelOffersRequest {

	sR.Radius = radius

	return sR
}

// SetRadiusUnit set radiusUnit | KM or MILE
func (sR *ShoppingHotelOffersRequest) SetRadiusUnit(radiusUnit RadiusUnit) *ShoppingHotelOffersRequest {

	sR.RadiusUnit = radiusUnit

	return sR
}

// AddHotelIDs add hotelIDs
func (sR *ShoppingHotelOffersRequest) AddHotelIDs(hotelIDs ...string) *ShoppingHotelOffersRequest {

	sR.HotelIDs = hotelIDs

	return sR
}

// AddChains add hotel chains filter
func (sR *ShoppingHotelOffersRequest) AddChains(chains ...string) *ShoppingHotelOffersRequest {

	sR.Chains = chains

	return sR
}

// AddRateCodes add rateCodes
func (sR *ShoppingHotelOffersRequest) AddRateCodes(rateCodes ...string) *ShoppingHotelOffersRequest {

	sR.RateCodes = rateCodes

	return sR
}

// AddAmenities add amenities
func (sR *ShoppingHotelOffersRequest) AddAmenities(amenities ...string) *ShoppingHotelOffersRequest {

	sR.Amenities = amenities

	return sR
}

// AddRatings add ratings
func (sR *ShoppingHotelOffersRequest) AddRatings(ratings ...int) *ShoppingHotelOffersRequest {

	sR.Ratings = ratings

	return sR
}

// SetPriceRange set priceRange
func (sR *ShoppingHotelOffersRequest) SetPriceRange(priceRange string) *ShoppingHotelOffersRequest {

	sR.PriceRange = priceRange

	return sR
}

// SetCurrency set currency
func (sR *ShoppingHotelOffersRequest) SetCurrency(currency string) *ShoppingHotelOffersRequest {

	sR.Currency = currency

	return sR
}

// SetPaymentPolicy set payment policy
func (sR *ShoppingHotelOffersRequest) SetPaymentPolicy(paymentPolicy HotelPaymentPolicy) *ShoppingHotelOffersRequest {

	sR.PaymentPolicy = paymentPolicy

	return sR
}

// SetBoardType set board type
func (sR *ShoppingHotelOffersRequest) SetBoardType(BoardType HotelBoardType) *ShoppingHotelOffersRequest {

	sR.BoardType = BoardType

	return sR
}

// IsIncludeClosed include closed
func (sR *ShoppingHotelOffersRequest) IsIncludeClosed(includeClosed bool) *ShoppingHotelOffersRequest {

	if includeClosed {
		sR.IncludeClosed = true
	}

	return sR
}

// IsBestRateOnly best rate only
func (sR *ShoppingHotelOffersRequest) IsBestRateOnly(bestRateOnly bool) *ShoppingHotelOffersRequest {

	if bestRateOnly {
		sR.BestRateOnly = true
	}

	return sR
}

// SetView set view
func (sR *ShoppingHotelOffersRequest) SetView(view HotelView) *ShoppingHotelOffersRequest {

	sR.View = view

	return sR
}

// SetSort set sort
func (sR *ShoppingHotelOffersRequest) SetSort(sort HotelSort) *ShoppingHotelOffersRequest {

	sR.Sort = sort

	return sR
}

// SetLang set lang
func (sR *ShoppingHotelOffersRequest) SetLang(lang string) *ShoppingHotelOffersRequest {

	sR.Lang = lang

	return sR
}

// Helper functions

// SetGeo set latitude & longitude
func (sR *ShoppingHotelOffersRequest) SetGeo(latitude, longitude float64) *ShoppingHotelOffersRequest {

	sR.SetLatitude(latitude)
	sR.SetLongitude(longitude)

	return sR
}

// JoinInts join int slice to string
func (sR ShoppingHotelOffersRequest) JoinInts(intSlice []int) string {

	valuesText := []string{}

	for i := range intSlice {
		number := intSlice[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}

	return strings.Join(valuesText, ",")
}

// GetURL returned key=value format for request on api
func (sR ShoppingHotelOffersRequest) GetURL(baseURL, reqType string) string {

	// set request url
	url := shoppingHotelOffersURL

	// add version
	switch reqType {
	case "GET":

		url = "/v2" + url

		// define query params
		queryParams := []string{}

		// check request values
		if sR.OfferID != "" {

			return baseURL + url + "/" + sR.OfferID

		} else if sR.HotelID != "" {

			url = url + "/by-hotel"

			queryParams = append(queryParams, "hotelId="+sR.HotelID)

		} else if sR.CityCode != "" {
			queryParams = append(queryParams, "cityCode="+sR.CityCode)
		} else if sR.Latitude != 0 && sR.Longitude != 0 {
			queryParams = append(queryParams, "latitude="+fmt.Sprintf("%v", sR.Latitude))
			queryParams = append(queryParams, "longitude="+fmt.Sprintf("%v", sR.Longitude))
		}

		if sR.CheckInDate != "" {
			queryParams = append(queryParams, "checkInDate="+sR.CheckInDate)
		}

		if sR.CheckOutDate != "" {
			queryParams = append(queryParams, "checkOutDate="+sR.CheckOutDate)
		}

		if sR.RoomQuantity != 0 {
			queryParams = append(queryParams, "roomQuantity="+fmt.Sprintf("%v", sR.RoomQuantity))
		}

		if sR.Adults != 0 {
			queryParams = append(queryParams, "adults="+fmt.Sprintf("%v", sR.Adults))
		}

		if len(sR.ChildAges) != 0 {
			queryParams = append(queryParams, "childAges="+sR.JoinInts(sR.ChildAges))
		}

		if sR.HotelName != "" {
			queryParams = append(queryParams, "hotelName="+sR.HotelName)
		}

		if len(sR.HotelIDs) != 0 {
			queryParams = append(queryParams, "hotelIds="+strings.Join(sR.HotelIDs, ","))
		}

		if len(sR.Chains) != 0 {
			queryParams = append(queryParams, "chains="+strings.Join(sR.Chains, ","))
		}

		if len(sR.RateCodes) != 0 {
			queryParams = append(queryParams, "rateCodes="+strings.Join(sR.RateCodes, ","))
		}

		if len(sR.Amenities) != 0 {
			queryParams = append(queryParams, "amenities="+strings.Join(sR.Amenities, ","))
		}

		if len(sR.Ratings) != 0 {
			queryParams = append(queryParams, "ratings="+sR.JoinInts(sR.Ratings))
		}

		if sR.Radius != 0 && sR.RadiusUnit != "" {
			queryParams = append(queryParams, "radius="+fmt.Sprintf("%v", sR.Radius))
			queryParams = append(queryParams, "radiusUnit="+sR.RadiusUnit.String())
		}

		if sR.PriceRange != "" && sR.Currency != "" {
			queryParams = append(queryParams, "priceRange="+sR.PriceRange)
			queryParams = append(queryParams, "currency="+sR.Currency)
		}

		if sR.PaymentPolicy != "" {
			queryParams = append(queryParams, "paymentPolicy="+sR.PaymentPolicy.String())
		}

		if sR.BoardType != "" {
			queryParams = append(queryParams, "boardType="+sR.BoardType.String())
		}

		if sR.IncludeClosed {
			queryParams = append(queryParams, "includeClosed=true")
		}

		if sR.BestRateOnly {
			queryParams = append(queryParams, "bestRateOnly=true")
		}

		if sR.View != "" {
			queryParams = append(queryParams, "view="+sR.View.String())
		}

		if sR.Sort != "" {
			queryParams = append(queryParams, "sort="+sR.Sort.String())
		}

		if sR.Lang != "" {
			queryParams = append(queryParams, "lang="+sR.Lang)
		}

		url = url + "?" + strings.Join(queryParams, "&")

		break
	}

	return baseURL + url
}

// GetBody prepare request body
func (sR ShoppingHotelOffersRequest) GetBody(reqType string) io.Reader {
	return nil
}

type ShoppingHotelOffersResponse struct {
	Dictionaries Dictionaries    `json:"dictionaries,omitempty"`
	Errors       []ErrorResponse `json:"errors,omitempty"`
	Meta         Meta            `json:"meta,omitempty"`
	Data         HotelData       `json:"data,omitempty"`
}

// Decode implement Response interface
func (dR *ShoppingHotelOffersResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

type ShoppingHotelsOffersResponse struct {
	Dictionaries Dictionaries    `json:"dictionaries,omitempty"`
	Data         []HotelData     `json:"data,omitempty"`
	Errors       []ErrorResponse `json:"errors,omitempty"`
	Meta         Meta            `json:"meta,omitempty"`
}

// Decode implement Response interface
func (dR *ShoppingHotelsOffersResponse) Decode(rsp []byte) error {

	err := json.Unmarshal(rsp, &dR)
	if err != nil {
		return err
	}

	return nil
}

// GetOffer return offer from list
func (dR ShoppingHotelsOffersResponse) GetHotel(offerNum int) HotelData {
	return dR.Data[offerNum]
}
