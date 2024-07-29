package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/kociumba/SkyDriver/env"
)

type Auctions struct {
	Success  bool `json:"success,omitempty"`
	Auctions []struct {
		UUID        string   `json:"uuid,omitempty"`
		Auctioneer  string   `json:"auctioneer,omitempty"`
		ProfileID   string   `json:"profile_id,omitempty"`
		Coop        []string `json:"coop,omitempty"`
		Start       int64    `json:"start,omitempty"`
		End         int64    `json:"end,omitempty"`
		ItemName    string   `json:"item_name,omitempty"`
		ItemLore    string   `json:"item_lore,omitempty"`
		Extra       string   `json:"extra,omitempty"`
		Category    string   `json:"category,omitempty"`
		Tier        string   `json:"tier,omitempty"`
		StartingBid int      `json:"starting_bid,omitempty"`
		ItemBytes   struct {
			Type int    `json:"type,omitempty"`
			Data string `json:"data,omitempty"`
		} `json:"item_bytes,omitempty"`
		Claimed          bool  `json:"claimed,omitempty"`
		ClaimedBidders   []any `json:"claimed_bidders,omitempty"`
		HighestBidAmount int   `json:"highest_bid_amount,omitempty"`
		Bids             []struct {
			AuctionID string `json:"auction_id,omitempty"`
			Bidder    string `json:"bidder,omitempty"`
			ProfileID string `json:"profile_id,omitempty"`
			Amount    int    `json:"amount,omitempty"`
			Timestamp int64  `json:"timestamp,omitempty"`
		} `json:"bids,omitempty"`
	} `json:"auctions,omitempty"`
}

// auctions seem to be kinda broken or just nobot has any auctions up 💀
func GetAuctions(username string) Auctions {
	uuid := ConvertUserToUUID(username)

	// Construct the URL with the uuid and the environment variable KEY
	url := fmt.Sprintf("https://api.hypixel.net/v2/skyblock/auction?uuid=%v&key=%v", uuid, env.KEY)

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Unmarshal the JSON response
	var collections Auctions
	if err := json.NewDecoder(resp.Body).Decode(&collections); err != nil {
		log.Error("Error decoding JSON:", "err", err, "resp", resp)
	}

	return collections
}
