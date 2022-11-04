package amadeus

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestShoppingHotelsOffers(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestShoppingHotelsOffers", func(t *testing.T) {

		// create client
		amadeus, err := New(
			os.Getenv("AMADEUS_CLIENT_ID"),
			os.Getenv("AMADEUS_CLIENT_SECRET"),
			os.Getenv("AMADEUS_ENV"),
		)
		if err != nil {
			t.Fatal("not expected error while creating client", err)
		}

		// get flights destinations
		req, resp, err := amadeus.NewRequest(ShoppingHotelsOffers)

		// set flights destination request params
		req.(*ShoppingHotelOffersRequest).SetCityCode("LON")

		// send request flight destinations
		err = amadeus.Do(req, &resp, "GET")
		if err != nil {
			t.Fatal("not expected error while geting destination data", err)
		}

		// get flight destination response
		respData := resp.(*ShoppingHotelsOffersResponse)

		// check if reponse exist
		if len(respData.Data) == 0 {
			t.Error("return 0 results in offer search request")
		}

	})

}

func TestShoppingHotelOffers(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestShoppingHotelOffers", func(t *testing.T) {

		// create client
		amadeus, err := New(
			os.Getenv("AMADEUS_CLIENT_ID"),
			os.Getenv("AMADEUS_CLIENT_SECRET"),
			os.Getenv("AMADEUS_ENV"),
		)
		if err != nil {
			t.Fatal("not expected error while creating client", err)
		}

		// get flights destinations
		req, resp, err := amadeus.NewRequest(ShoppingHotelOffers)

		// set flights destination request params
		req.(*ShoppingHotelOffersRequest).SetHotelID("HILONBE3")

		// send request flight destinations
		err = amadeus.Do(req, &resp, "GET")
		if err != nil {
			t.Fatal("not expected error while geting destination data", err)
		}

		// get flight destination response
		respData := resp.(*ShoppingHotelOffersResponse)

		// check if reponse exist
		if len(respData.Data.Offers) == 0 {
			t.Error("return 0 results in offer search request")
		}

	})

}

func TestShoppingHotelOffer(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not found .env file")
	}

	t.Run("TestShoppingHotelOffer", func(t *testing.T) {

		// create client
		amadeus, err := New(
			os.Getenv("AMADEUS_CLIENT_ID"),
			os.Getenv("AMADEUS_CLIENT_SECRET"),
			os.Getenv("AMADEUS_ENV"),
		)
		if err != nil {
			t.Fatal("not expected error while creating client", err)
		}

		// get flights destinations
		req, resp, err := amadeus.NewRequest(ShoppingHotelOffers)

		// set flights destination request params
		req.(*ShoppingHotelOffersRequest).SetOfferID("3442F83BF3BF9482A7B058D67959FE807FB69B9344F85A4B3F1893DD903E1791")

		// send request flight destinations
		err = amadeus.Do(req, &resp, "GET")
		if err != nil {
			t.Fatal("not expected error while geting destination data", err)
		}

		// get flight destination response
		respData := resp.(*ShoppingHotelOffersResponse)

		// check if reponse exist
		if len(respData.Data.Offers) == 0 {
			t.Error("return 0 results in offer search request")
		}

	})

}
