package model

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net"
	"sync"
	"time"

	"github.com/globalsign/mgo"
	"github.com/sirupsen/logrus"
)

const (
	MongoRunMode     = "Strong"
	MongoPoolMax     = 4096
	MongoSyncTimeout = 5
)

type Model struct {
	mux        sync.RWMutex
	session    *mgo.Session
	collection *mgo.Collection

	config  *Config
	logger  *logrus.Logger
	indexes map[string]bool
}

func NewModel(cfg *Config, log *logrus.Logger) *Model {
	dsn := "mongodb://"

	user := cfg.GetUser()
	pwd := cfg.GetPasswd()

	if user != "" && pwd != "" {
		dsn += user + ":" + pwd + "@"
	}

	dsn += cfg.Host
	if cfg.Database != "" {
		dsn += "/" + cfg.Database
	}

	dialInfo, err := mgo.ParseURL(dsn)
	if err != nil {
		log.Panic(err)
	}

	if cfg.PEMFILE != "" {
		rootPem, err := ioutil.ReadFile(cfg.PEMFILE)
		if err != nil {
			log.Panic(err)
		}

		roots := x509.NewCertPool()
		if !roots.AppendCertsFromPEM(rootPem) {
			log.Panic("fail to parse root certificate")
		}

		tlsCfg := &tls.Config{
			RootCAs:            roots,
			InsecureSkipVerify: true,
		}

		dialInfo.DialServer = func(addr *mgo.ServerAddr) (conn net.Conn, e error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsCfg)
			if err != nil {
				log.Println(err)
			}
			return conn, err
		}
	}

	//	set server dial time out
	dialInfo.Timeout = cfg.Timeout * time.Second

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Panic(err)
	}

	if err := session.Ping(); err != nil {
		log.Panic(err)
	}

	//	set session mode
	switch cfg.Mode {
	case "Strong":
		session.SetMode(mgo.Strong, true)
	case "Monotonic":
		session.SetMode(mgo.Monotonic, true)
	case "Eventual":
		session.SetMode(mgo.Eventual, true)
	default:
		session.SetMode(mgo.Strong, true)
	}

	//	set session safe
	session.SetSafe(&mgo.Safe{
		W:        1,
		WTimeout: 200,
	})

	//	 set pool size
	if cfg.Pool > 0 {
		if cfg.Pool > MongoPoolMax {
			cfg.Pool = MongoPoolMax
		}

		session.SetPoolLimit(cfg.Pool)
	}

	//	 set op response timeout
	if cfg.Timeout == 0 {
		cfg.Timeout = MongoSyncTimeout
	}
	session.SetSyncTimeout(cfg.Timeout * time.Second)

	if err := session.Ping(); err != nil {
		panic(err)
	}

	return &Model{
		session: session,
		config:  cfg,
		logger:  log,
		indexes: make(map[string]bool),
	}
}

func (model *Model) Use(database string) *Model {
	model.config.Database = database
	return model
}

func (model *Model) Copy() *Model {
	return &Model{
		config:  model.config,
		session: model.session,
		logger:  model.logger,
	}
}

func (model *Model) Database() string {
	return model.config.Database
}

//copy db
func (model *Model) C(db string) *Model {
	copiedDb := model.Copy()
	copiedDb.collection = copiedDb.session.DB(model.Database()).C(db)
	return copiedDb
}

func (model *Model) Config() *Config {
	return model.config
}

func (model *Model) Session() *mgo.Session {
	return model.session
}

func (model *Model) Collection() *mgo.Collection {
	return model.collection
}

func (model *Model) Query(collectionName string, collectionIndexes []mgo.Index, query func(*mgo.Collection)) {
	copiedDb := model.C(collectionName)
	defer copiedDb.Close()

	copidCollection := copiedDb.Collection()

	if !model.indexes[collectionName] {
		model.mux.Lock()
		if !model.indexes[collectionName] {
			for _, index := range collectionIndexes {
				if err := copidCollection.EnsureIndex(index); err != nil {
					model.indexes[collectionName] = false
					model.logger.Printf("Ensure index of %s (%#v) : %v", collectionName, index, err)
				}
			}
			model.indexes[collectionName] = true
		}
		model.mux.Unlock()
	}

	query(copidCollection)
}

func (model *Model) Close() {
	model.session.Close()
}
