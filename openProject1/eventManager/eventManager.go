package eventManager

import (
	"fmt"
	"time"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"os"
)

type ServInfo struct {
	Addr		string
	MgoSession	*mgo.Session
}

type gameInfo struct {
	id					int
	eventRunning		bool
	lastEventCompletion	string
}

func retrieveGames() []gameInfo {
	currentGames := []gameInfo{
		{123456789, false, time.Now().String()},
		{987654321, false, time.Monday.String()},
		{789456123, true, time.Thursday.String()},
		{321654987, true, time.Friday.String()}}
	return currentGames
}

func retrieveGamesFromMongoDB() []gameInfo {
	currentGames := []gameInfo{
		{123456789, false, time.Now().String()},
		{987654321, false, time.Monday.String()},
		{789456123, true, time.Thursday.String()},
		{321654987, true, time.Friday.String()}}

	return currentGames
}

func ParseGamesForEvents(simulation bool, mongo ServInfo) {
	if mongo.MgoSession == nil && !simulation {
		sess, err := mgo.Dial(mongo.Addr)
		mongo.MgoSession = sess
		if err != nil {
			fmt.Printf("err: %#v", err)
			os.Exit(-1)
		}
	}

	var games []gameInfo
	if simulation {
		games = retrieveGames()
	} else {
		games = retrieveGamesFromMongoDB()
	}

	for i, elem := range games {
		fmt.Printf("Games #%d with ID: %d", i, elem.id)
		if elem.eventRunning {
			fmt.Printf(" is in an event.")
		} else {
			fmt.Printf(" was in an event that ended %s", elem.lastEventCompletion)
		}
		fmt.Println("");
	}
}
