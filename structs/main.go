package main
import "fmt"
type contactInfo struct{
    email string
    zip int
}
type person struct{
    firstName string
    lastName string
    contact contactInfo

}

func main(){

    jim := person{
        firstName: "Jim",
        lastName: "Parson",
        contact: contactInfo{
            email: "jim@t.com",
            zip: 111,

        },
    }
    jim.updatePerson("Yooo")
    jim.print()
}
func (p person) print(){
    
    fmt.Println("%+v",p)

}

func (p *person) updatePerson(name string){
    (*p).firstName = name
    fmt.Println(*p)
}
