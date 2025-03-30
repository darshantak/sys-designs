package main

type Booking struct {
	CustomerId string
	Items      map[string]int
}

func Book(custId string, items map[string]int) Booking {
	newBooking := Booking{
		CustomerId: custId,
		Items:      items,
	}
	//Check with the restaurants that the quantity is matched or not. 
	return newBooking
}


// book (‘cust1’
// 	, {‘bendi
// 	macaroni’: 3,
// 	_
// 	‘samosa
// 	_pizza’: 2’})
