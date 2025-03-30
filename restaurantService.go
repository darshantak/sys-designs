package main

type Menu struct {
	Items map[string]int
}

type Restaurant struct {
	Name          string
	Menu          map[string]int
	TotalCapacity int
	CapacityInUse int
}

type Restaurants struct {
	RestaurantList []Restaurant
}

func RestaurantService() *Restaurants {
	return &Restaurants{
		RestaurantList: make([]Restaurant, 0),
	}
}
func (rs *Restaurants) AddRestaurant(restName string, menu map[string]int, capacity int) Restaurant {
	newRest := Restaurant{
		Name:          restName,
		Menu:          menu,
		TotalCapacity: capacity,
		CapacityInUse: capacity,
	}
	rs.RestaurantList = append(rs.RestaurantList, newRest)
	return newRest
}

func (rs *Restaurants) GetRestaunts() []Restaurant {
	return rs.RestaurantList
}

