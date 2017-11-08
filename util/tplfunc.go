package util
import (
	"github.com/astaxie/beego"
)
func head(in string)(out string){
	out = StripTags(in)
    return out[:150]+"..."
}
func init(){
	beego.AddFuncMap("head",head)
}