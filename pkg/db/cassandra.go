package db

import (
	"fmt"
	"github.com/gocql/gocql"
	"sync"
)

var (
	cassandraDB     *CassandraDB
	cassandraDBOnce sync.Once
)

type CassandraDB struct {
	Instance *gocql.Session
	config   *DBConfig
}

func GetDB() *CassandraDB {
	return cassandraDB
}

func Init(conf *DBConfig) {
	cassandraDBOnce.Do(func() {
		cassandraDB = &CassandraDB{
			config: conf,
		}
		cassandraDB.connect()
	})
}

func (db *CassandraDB) connect() {
	cluster := gocql.NewCluster(db.config.Host)
	cluster.Keyspace = db.config.Keyspace
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()
	if err != nil {
		panic(fmt.Sprintf("Could not connect to Cassandra: %v", err))
	}

	db.Instance = session
}

func (db *CassandraDB) Close() {
	if db.Instance != nil {
		db.Instance.Close()
	}
}
