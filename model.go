package validator

var tags = [...]string{
	"field",  //required
	"type",   //number, text
	"min",    //number of min value
	"max",    // number of max value
	"format", //email
	"startswith",
	"endswith",
	"value_of",
}

const (
	MAIL_FORMAT  = `\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`
	ALPHANUMERIC = `^[A-Za-z0-9 ]*$`
	ALPHABET     = `^[A-Za-z ]*$`
)

type myvalidator interface {
	validate() error
}
type structDetail struct {
	name      string
	val       interface{}
	tag_name  string
	tag_value string
	required  bool
}
type num_error int
type def struct{ *structDetail }
type number struct{ *structDetail }
type text struct{ *structDetail }
type email struct{ *structDetail }
type alphanumeric struct{ *structDetail }
type alphabet struct{ *structDetail }
type required struct{ *structDetail }
type startswith struct{ *structDetail }
type endswith struct{ *structDetail }
type valueOf struct{ *structDetail }
type min struct {
	min    int
	detail *structDetail
}
type max struct {
	max    int
	detail *structDetail
}
