package main

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
	Product struct {
		ID             int64     `json:"id"`
		Title          string    `json:"title"`
		BodyHTML       string    `json:"body_html"`
		Vendor         string    `json:"vendor"`
		ProductType    string    `json:"product_type"`
		CreatedAt      time.Time `json:"created_at"`
		Handle         string    `json:"handle"`
		UpdatedAt      time.Time `json:"updated_at"`
		PublishedAt    time.Time `json:"published_at"`
		TemplateSuffix string    `json:"template_suffix"`
		PublishedScope string    `json:"published_scope"`
		Tags           string    `json:"tags"`
		Variants       []struct {
			ID                   int64       `json:"id"`
			ProductID            int64       `json:"product_id"`
			Title                string      `json:"title"`
			Price                string      `json:"price"`
			Sku                  string      `json:"sku"`
			Position             int         `json:"position"`
			InventoryPolicy      string      `json:"inventory_policy"`
			CompareAtPrice       string      `json:"compare_at_price"`
			FulfillmentService   string      `json:"fulfillment_service"`
			InventoryManagement  string      `json:"inventory_management"`
			Option1              string      `json:"option1"`
			Option2              string      `json:"option2"`
			Option3              string      `json:"option3"`
			CreatedAt            time.Time   `json:"created_at"`
			UpdatedAt            time.Time   `json:"updated_at"`
			Taxable              bool        `json:"taxable"`
			Barcode              interface{} `json:"barcode"`
			Grams                int         `json:"grams"`
			ImageID              int64       `json:"image_id"`
			Weight               float32     `json:"weight"`
			WeightUnit           string      `json:"weight_unit"`
			InventoryQuantity    int         `json:"inventory_quantity"`
			OldInventoryQuantity int         `json:"old_inventory_quantity"`
			RequiresShipping     bool        `json:"requires_shipping"`
		} `json:"variants"`
		Options []struct {
			ID        int64    `json:"id"`
			ProductID int64    `json:"product_id"`
			Name      string   `json:"name"`
			Position  int      `json:"position"`
			Values    []string `json:"values"`
		} `json:"options"`
		Images []struct {
			ID         int64     `json:"id"`
			ProductID  int64     `json:"product_id"`
			Position   int       `json:"position"`
			CreatedAt  time.Time `json:"created_at"`
			UpdatedAt  time.Time `json:"updated_at"`
			Alt        string    `json:"alt"`
			Width      int       `json:"width"`
			Height     int       `json:"height"`
			Src        string    `json:"src"`
			VariantIds []int64   `json:"variant_ids"`
		} `json:"images"`
		Image struct {
			ID         int64     `json:"id"`
			ProductID  int64     `json:"product_id"`
			Position   int       `json:"position"`
			CreatedAt  time.Time `json:"created_at"`
			UpdatedAt  time.Time `json:"updated_at"`
			Alt        string    `json:"alt"`
			Width      int       `json:"width"`
			Height     int       `json:"height"`
			Src        string    `json:"src"`
			VariantIds []int64   `json:"variant_ids"`
		} `json:"image"`
	} `json:"product"`
}

func main() {
	for {
		url := createURL()
		log.Printf("requesting to %s", url)

		s, err := get(url)
		if err != nil {
			log.Println(s)
		}

		print(s)

		time.Sleep(6 * time.Minute)
	}
}

func createURL() string {
	sec := strconv.FormatInt(time.Now().Unix(), 10)
	return "https://yamatomichi.myshopify.com/products/dw-5-pocket-shorts-m.json?_=" + sec
}

func get(url string) (*Shop, error) {
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

func print(s *Shop) {
	for _, v := range s.Product.Variants {
		if v.InventoryQuantity > 0 {
			log.Printf("title: %s, quantity: %d, old_quantity: %d, updated_at: %v\n", v.Title, v.InventoryQuantity, v.OldInventoryQuantity, v.UpdatedAt)
		}
	}
}
