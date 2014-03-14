package dredd

import(
  "fmt"
  "time"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
)

const(
  MONGO_URL = "localhost:27017"
  LOADMON_COLLECTION = "loadmon_test"
  LOADMON_DB = "loadmon"
)

type MongoWriter struct {
  *mgo.Session
  collection *mgo.Collection
}

var MongoConnection MongoWriter

func NewMongoConnection() {
  session,err := mgo.Dial(MONGO_URL)
  if err != nil {
    panic(err)
  }

  session.SetMode(mgo.Monotonic,true)
  collection := session.DB(LOADMON_DB).C(LOADMON_COLLECTION)

  MongoConnection.collection = collection
}


func (m *MongoWriter) UpdateTotal(t time.Time, hour string, minute string, tot int64, inttot int64, outtot int64) {
  query := bson.M{"day":t.YearDay()}
  change := bson.M{"$inc": bson.M{
    "daily":tot,
    "dailyinbound":inttot,
    "dailyoutbound":outtot,
    "hourly."+hour: tot,
    "minute."+hour+"."+minute: tot,
  }}

  m.Write(query,change)
  return
}

func (m *MongoWriter) Write(query bson.M, change bson.M) {
  fmt.Println(query)
  fmt.Println(change)
  /*mongo_err := m.collection.Update(query,change)*/
  //if mongo_err != nil {
    //fmt.Println(mongo_err)
    //panic(mongo_err)
    //return
  /*}*/
}

