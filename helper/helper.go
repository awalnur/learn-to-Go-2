package helper

var AppName string = "APP Name" // public var
var appVersion = "APP Version"  // private

func privateFunc(string string) string {
	return "test" + appVersion
}
func SayHello(data string) string {
	//	penamaan method/function harus diawali hurus besar
	//
	return data
}
func Public() string {
	return privateFunc("hello")

}
