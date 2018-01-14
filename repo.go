package main

import mgo "gopkg.in/mgo.v2"
import "log"
import "gopkg.in/mgo.v2/bson"

// OpenSession : opens session in db
func OpenSession() (*mgo.Session, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return session, err
	}

	return session, nil
}

// CreateOrder : creates order in db
func CreateOrder(order Order) (bson.ObjectId, error) {

	session, err := OpenSession()
	if err != nil {
		log.Panic(err)
		return "", err
	}

	order._id = bson.NewObjectId()

	err = session.DB("main").C("orders").Insert(order)
	log.Print(err)
	defer session.Close()

	if err != nil {
		log.Panic(err)
		return "", err
	}

	if len(order.User) > 0 {
		SendMail(order.User)
	}

	return order._id, nil
}

// GetReservedSeats :
func GetReservedSeats(routeID string) ([]Seat, error) {

	var result []Seat

	session, err := OpenSession()
	if err != nil {
		log.Panic(err)
		return result, err
	}

	log.Print(routeID)

	var orders []Order
	gameQuery := bson.M{"routeid": routeID}
	err = session.DB("main").C("orders").Find(gameQuery).All(&orders)
	if err != nil {
		log.Panic(err)
		return result, err
	}

	log.Print(orders)
	for _, elem := range orders {
		result = append(result, elem.Seats...)
	}

	defer session.Close()

	return result, nil
}
