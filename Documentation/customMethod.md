create aliases for types to enable custom methods

CUSTOM TYPES!

```golang
type str string

func (text str) log(){
    fmt.Println(text)
}

func main(){
    var name str = "Will"  //need to explicitly type or Go would infer normal string type

    name.log()
}
```
