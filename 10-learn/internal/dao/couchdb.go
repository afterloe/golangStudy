package dao

import (
	"context"
	"fmt"
	"github.com/go-kivik/kivik/v4"
	_ "github.com/go-kivik/kivik/v4/couchdb"
	"github.com/google/uuid"
	"log"
	"onenet/internal/util"
	"reflect"
	"time"
)

var (
	driverName = "couch"
	Couchdb    *_couchDB
)

type _couchDB struct {
	protocol  string
	username  string
	password  string
	url       string
	dbName    string
	client    *kivik.Client
	db        *kivik.DB
	failCount int
}

func Register(config util.Config) {
	Couchdb = &_couchDB{
		protocol:  "http",
		username:  config.DBUser,
		password:  config.DBPassword,
		url:       config.DBUrl,
		dbName:    config.DBName,
		failCount: 0,
	}
}

func (that *_couchDB) generatorURL() string {
	return fmt.Sprintf("%s://%s:%s@%s/", that.protocol, that.username, that.password, that.url)
}

func (that *_couchDB) Find(dbName string, query interface{}, options ...kivik.Option) *kivik.ResultSet {
	db := that.GetDB(dbName)
	return db.Find(context.TODO(), query, options...)
}

func (that *_couchDB) Put(dbName string, entity interface{}, options ...kivik.Option) (string, error) {
	db := that.GetDB(dbName)
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	idStr := id.String()
	reflectValue := reflect.ValueOf(entity)
	reflectValue.Elem().FieldByName("Id").SetString(idStr)
	//reflectValue.Elem().FieldByName("CreatedBy").SetString()
	reflectValue.Elem().FieldByName("CreatedAt").Set(reflect.ValueOf(time.Now()))
	reflectValue.Elem().FieldByName("Deleted").SetBool(false)
	rev, err := db.Put(context.TODO(), idStr, entity, options...)
	if err != nil {
		return "", err
	}
	log.Printf("album inserted with revision %s", rev)
	return idStr, nil
}

func (that *_couchDB) GetDB(dbNames ...string) *kivik.DB {
	client := that.GetClient()
	if len(dbNames) == 0 {
		return client.DB(that.dbName)
	} else {
		return client.DB(dbNames[0])
	}
}

func (that *_couchDB) GetClient() *kivik.Client {
register:
	if that.client == nil {
		client, err := kivik.New(driverName, that.generatorURL())
		if err != nil {
			log.Fatal(err)
		}
		that.client = client
	}

	if flag, err := that.client.Ping(context.TODO()); nil != err || flag == false {
		log.Println("client is close")
		goto register
	}

	return that.client
}
