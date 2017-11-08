package models
import(
	"log"
    "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/astaxie/beego"
)
type Post struct {
    Name     string `bson:"name"`
    Category string `bson:"category"`
    Address  string `bson:"address"`
	Site     string `bson:"site"`
	Cover    string `bson:"cover"`
	Publish  bool   `bson:"publish"`
	Content  map[string]PostContent `bson:"content"`
}

type PostContent struct{
	Title    string `bson:"title"`
    Location string `bson:"location"`
	Date     string `bson:"date"`
	Content  string `bson:"content"`
}

func AddPost(){

}

func GetPosts(page int,pageSize int)[]*Post{
	session, err := mgo.Dial(beego.AppConfig.String("db-host"))
	if err != nil {
			panic(err)
	}

	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(beego.AppConfig.String("db-name")).C("posts")
	var posts []*Post
	err = c.Find(bson.M{}).Skip(page*pageSize).Limit(pageSize).Iter().All(&posts)
	beego.Debug(posts)
	if err != nil {
			log.Fatal(err)
	}
	return posts
}

func GetPostsCount()int{
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
			panic(err)
	}
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("travel").C("posts")
	var count int
	count,err = c.Find(bson.M{}).Count()
	if err != nil {
			log.Fatal(err)
	}
	return count
}


func GetPost(name string)(*Post,error){
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
			panic(err)
	}
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("travel").C("posts")
	var post Post
	err = c.Find(bson.M{"name": name}).One(&post)
	// if err != nil {
	// 		log.Fatal(err)
	// }
	return &post,err
}

func Search(text string,page int,pageSize int)[]*Post{
	session, err := mgo.Dial(beego.AppConfig.String("db-host"))
	if err != nil {
			panic(err)
	}

	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(beego.AppConfig.String("db-name")).C("posts")
	var posts []*Post
	err = c.Find(bson.M{}).Skip(page*pageSize).Limit(pageSize).Iter().All(&posts)
	beego.Debug(posts)
	if err != nil {
			log.Fatal(err)
	}
	return posts
}