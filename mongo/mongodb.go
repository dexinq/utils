package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Mongodb struct {
	url string
	database string
	collection string
}

func NewMongodb(url, database, collection string) *Mongodb {
	return &Mongodb{url:url, database:database, collection:collection}
}

func (m Mongodb)getSession() (*mgo.Session, error) {
	session, err := mgo.Dial(m.url)
	if err!=nil{
		return nil, err
	}
	return session, err
}

func (m *Mongodb)InsertMongo(data interface{}) error {
	session, err:=(*m).getSession()
	defer session.Close()

	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	db := session.DB(m.database)
	collection := db.C(m.collection)
	err = collection.Insert(data)
	if err!=nil{
		log.Fatalf(err.Error())
		return err
	}
	return nil
}

func (m *Mongodb)UpdateMongo(data interface{}) error {
	session, err:=(*m).getSession()
	defer session.Close()

	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	db := session.DB(m.database)
	collection := db.C(m.collection)
	query := bson.M{"_id":data}
	_, err = collection.Upsert(query, data)
	if err!=nil{
		log.Fatalf(err.Error())
	}

	return nil
}

func (m *Mongodb)GetMongo(res *[]interface{}) error{
	session, err:=(*m).getSession()
	defer session.Close()
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}

	db := session.DB(m.database)
	collection := db.C(m.collection)
	err = collection.Find(nil).All(res)

	if err!=nil{
		log.Fatalf(err.Error())
		return err
	}
	return nil
}

func (m *Mongodb)GetOne(query map[string]string, res *[]interface{}) error {
	session, err:=(*m).getSession()
	defer session.Close()
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}

	db := session.DB(m.database)
	collection := db.C(m.collection)
	err = collection.Find(query).All(res)

	if err!=nil{
		log.Fatalf(err.Error())
		return err
	}
	return nil
}
