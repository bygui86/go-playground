package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
)

const (
	cassandraHost     = "localhost"
	cassandraPort     = 9042
	cassandraKeyspace = "example"

	insertQuery         = "INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)"
	selectSpecificQuery = "SELECT id, text FROM tweet WHERE timeline = ? LIMIT 1"
	selectAllQuery      = "SELECT id, text FROM tweet WHERE timeline = ?"
	truncateQuery       = "truncate example.tweet;"

	johnTimeline = "john"
	johnText1    = "john doe thinks something"
	johnText2    = "john doe thinks something else now"
	janeTimeline = "jane"
	janeText     = "jane doe thinks something different"
)

var session *gocql.Session

/*
	Before you execute the program, Launch `cqlsh` and execute:
		create keyspace example with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
		create table example.tweet(timeline text, id UUID, text text, PRIMARY KEY(id));
		create index on example.tweet(timeline);
*/
func main() {
	createConnection()
	defer session.Close()

	insertTweet(johnTimeline, johnText1)
	insertTweet(johnTimeline, johnText2)
	insertTweet(janeTimeline, janeText)

	searchAndPrintSpecificTweet(johnTimeline)

	searchAndPrintAllTweets(johnTimeline)

	time.Sleep(1)
	cleanup()
}

func createConnection() {
	clusterCfg := gocql.NewCluster(cassandraHost)
	// clusterCfg := gocql.NewCluster("127.0.0.1")
	clusterCfg.Port = cassandraPort
	clusterCfg.Keyspace = cassandraKeyspace
	// clusterCfg.Consistency = gocql.Quorum
	// clusterCfg.Consistency = gocql.EachQuorum
	// clusterCfg.Consistency = gocql.LocalQuorum
	clusterCfg.ConnectTimeout = 5 * time.Second // initial connection timeout, used during initial dial to server (default: 600ms)
	clusterCfg.Timeout = 5 * time.Second        // connection timeout (default: 600ms)
	// clusterCfg.NumConns = 5 // number of connections per host (default: 2)

	var err error
	session, err = clusterCfg.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
}

func insertTweet(timeline, text string) {
	gocql.RandomUUID()
	err := session.Query(insertQuery, timeline, gocql.TimeUUID(), text).Exec()
	if err != nil {
		log.Fatal(err)
	}
}

/*
	Search for a specific set of records whose 'timeline' column matches the input parameter.
	The secondary index that we created earlier will be used for optimizing the search.
*/
func searchAndPrintSpecificTweet(timeline string) {
	var id gocql.UUID
	var text string
	err := session.Query(selectSpecificQuery, timeline).Consistency(gocql.One).Scan(&id, &text)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Tweet [%v]: %s \n", id, text)
}

func searchAndPrintAllTweets(timeline string) {
	var id gocql.UUID
	var text string

	cqlIter := session.Query(selectAllQuery, timeline).Iter()
	for cqlIter.Scan(&id, &text) {
		fmt.Printf("Tweet [%v]: %s \n", id, text)
	}
	err := cqlIter.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func cleanup() {
	err := session.Query(truncateQuery).Exec()
	if err != nil {
		log.Fatal(err)
	}
}
