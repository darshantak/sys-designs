// {"item1":[{},{}], "item2":[{},{},{}]}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type ItemType string

const (
	Product ItemType = "product"
	Movie   ItemType = "movie"
)

type Rating struct {
	UserId string
	Rating int
}

type RatingManager struct {
	ItemRatings map[string][]Rating
	mu          sync.RWMutex
}

func newRatingManager(id string) *RatingManager {
	return &RatingManager{
		ItemRatings: make(map[string][]Rating),
	}
}
func (rm *RatingManager) listRatings() {

	fmt.Println(rm.ItemRatings)

}
func (rm *RatingManager) createNewItem(itemId string) bool {

	rm.ItemRatings[itemId] = make([]Rating, 0)
	fmt.Println("New item added to the manager")
	return true
}

func (rm *RatingManager) rateItem(userId string, itemId string, rating int) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	if _, exists := rm.ItemRatings[itemId]; !exists {
		isExists := rm.createNewItem(itemId)
		if isExists {
			rm.rateItem(userId, itemId, rating)
			return
		} else {
			fmt.Println("Something wrong")
		}
	}
	//If user has already rated the item
	for _, ratingStruct := range rm.ItemRatings[itemId] {
		if ratingStruct.UserId == userId {
			fmt.Printf("User %s has already rated. Please use update command to update the rating. \n", userId)
			return
		}
	}
	newRating := Rating{
		UserId: userId,
		Rating: rating,
	}
	rm.ItemRatings[itemId] = append(rm.ItemRatings[itemId], newRating)
	fmt.Printf("Rating %d for item %s added \n", rating, itemId)
}
func (rm *RatingManager) updateRating(userId, itemId string, rating int) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	if _, exists := rm.ItemRatings[itemId]; !exists {
		fmt.Println("This item doesn't exist")
		return
	}
	flag := false
	for _, ratingStruct := range rm.ItemRatings[itemId] {
		if ratingStruct.UserId == userId {
			flag = true
			newRating := Rating{
				UserId: userId,
				Rating: rating,
			}
			ratingStruct = newRating
		}
	}
	if !flag {
		fmt.Println("User doesn't exist for this item")
	} else {
		fmt.Println("New rating has been updated")
	}
}
func (rm *RatingManager) deleteRating(userId, itemId string) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	if _, exists := rm.ItemRatings[itemId]; !exists {
		fmt.Println("This item does not exists")
		return
	}
	ratingsArr := rm.ItemRatings[itemId]
	for i, r := range ratingsArr {
		if r.UserId == userId {
			rm.ItemRatings[itemId] = append(ratingsArr[:i], ratingsArr[i+1:]...)
		}
	}
	fmt.Printf("Rating removed for user %s \n", userId)
}
func RatingServiceCli() {
	newRM := newRatingManager("RM123")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()

		parts := strings.Fields(input)
		command, args := parts[0], parts[1:]
		fmt.Println(args)
		switch command {
		case "help":
			fmt.Println("Usage \n add <userId> <itemId> <rating>")
			fmt.Println("update <userId> <itemId> <rating>")
			fmt.Println("delete <userId> <itemId>")
		case "list":
			newRM.listRatings()
		case "add":
			rating, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Println("Rating must be an integer.")
				continue
			}
			newRM.rateItem(args[0], args[1], rating)
		case "update":
			rating, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Println("Rating must be an integer.")
				continue
			}
			newRM.updateRating(args[0], args[1], rating)
		case "delete":
			newRM.deleteRating(args[0], args[1])
		default:
			break
		}
	}
}

// Sample queries for this tool

// add user1 item1 2
// add user2 item1 4
// add user1 item2 1
// add user1 item3 5

// add user1 item1 2
// update user1 item1 4
// update user1 item4 5

//delete user1 item1
