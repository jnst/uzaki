package applestore

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Shop struct {
	Head struct {
		Status string `json:"status"`
		Data   struct {
		} `json:"data"`
	} `json:"head"`
	Body struct {
		Stores []struct {
			StoreEmail         string `json:"storeEmail"`
			StoreName          string `json:"storeName"`
			ReservationURL     string `json:"reservationUrl"`
			MakeReservationURL string `json:"makeReservationUrl"`
			State              string `json:"state"`
			StoreImageURL      string `json:"storeImageUrl"`
			Country            string `json:"country"`
			City               string `json:"city"`
			StoreNumber        string `json:"storeNumber"`
			PartsAvailability  struct {
				Z0YQ struct {
					StorePickEligible     bool   `json:"storePickEligible"`
					StoreSearchEnabled    bool   `json:"storeSearchEnabled"`
					StoreSelectionEnabled bool   `json:"storeSelectionEnabled"`
					StorePickupQuote      string `json:"storePickupQuote"`
					PickupSearchQuote     string `json:"pickupSearchQuote"`
					StorePickupLabel      string `json:"storePickupLabel"`
					PartNumber            string `json:"partNumber"`
					PurchaseOption        string `json:"purchaseOption"`
					CtoOptions            string `json:"ctoOptions"`
					StorePickupLinkText   string `json:"storePickupLinkText"`
					PickupDisplay         string `json:"pickupDisplay"`
				} `json:"Z0YQ"`
			} `json:"partsAvailability"`
			PhoneNumber                string `json:"phoneNumber"`
			PickupTypeAvailabilityText string `json:"pickupTypeAvailabilityText"`
			Address                    struct {
				Address    string      `json:"address"`
				Address3   interface{} `json:"address3"`
				Address2   string      `json:"address2"`
				PostalCode string      `json:"postalCode"`
			} `json:"address"`
			HoursURL   string `json:"hoursUrl"`
			StoreHours struct {
				StoreHoursText   string `json:"storeHoursText"`
				BopisPickupDays  string `json:"bopisPickupDays"`
				BopisPickupHours string `json:"bopisPickupHours"`
				Hours            []struct {
					StoreTimings string `json:"storeTimings"`
					StoreDays    string `json:"storeDays"`
				} `json:"hours"`
			} `json:"storeHours"`
			Storelatitude        float64 `json:"storelatitude"`
			Storelongitude       float64 `json:"storelongitude"`
			Storedistance        float64 `json:"storedistance"`
			StoreDistanceVoText  string  `json:"storeDistanceVoText"`
			Storelistnumber      int     `json:"storelistnumber"`
			StoreListNumber      int     `json:"storeListNumber"`
			PickupOptionsDetails struct {
				WhatToExpectAtPickup     string `json:"whatToExpectAtPickup"`
				ComparePickupOptionsLink string `json:"comparePickupOptionsLink"`
				PickupOptions            []struct {
					PickupOptionTitle       string `json:"pickupOptionTitle"`
					PickupOptionDescription string `json:"pickupOptionDescription"`
					Index                   int    `json:"index"`
				} `json:"pickupOptions"`
			} `json:"pickupOptionsDetails"`
		} `json:"stores"`
		OverlayInitiatedFromWarmStart bool   `json:"overlayInitiatedFromWarmStart"`
		ViewMoreHoursLinkText         string `json:"viewMoreHoursLinkText"`
		StoresCount                   string `json:"storesCount"`
		Little                        bool   `json:"little"`
		PickupLocationLabel           string `json:"pickupLocationLabel"`
		PickupLocation                string `json:"pickupLocation"`
		NotAvailableNearby            string `json:"notAvailableNearby"`
		NotAvailableNearOneStore      string `json:"notAvailableNearOneStore"`
		WarmDudeWithAPU               bool   `json:"warmDudeWithAPU"`
		ViewMoreHoursVoText           string `json:"viewMoreHoursVoText"`
	} `json:"body"`
}

func CreateURL() string {
	sec := strconv.FormatInt(time.Now().Unix(), 10)
	return "https://www.apple.com/jp/shop/retail/pickup-message?parts.0=Z0YQ&option.0=MG1A3J%2FA%2CMY7A2FE%2FA&store=R224&_=" + sec
}

func Get(url string) (*Shop, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("request failed: %d", resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	s := &Shop{}
	err = json.Unmarshal(body, s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func Print(s *Shop) {
	for _, v := range s.Body.Stores {
		if v.PartsAvailability.Z0YQ.StoreSelectionEnabled {
			log.Println("In stock.")
		} else {
			log.Println("Out of stock.")
		}
	}
}
