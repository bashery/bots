package main

import "github.com/sonyarouje/simdb/db"

type Customer struct {
	CustID string `json:"custid"`
	Name string `json:"name"`
	Address string `json:"address"`
	Contact Contact
}

type Contact struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
}

//ID any struct that needs to persist should implement this function defined 
//in Entity interface.
func (c Customer) ID() (jsonField string, value interface{}) {
	value=c.CustID
	jsonField="custid"
	return
}

func main() {
    driver, err:=db.New("data")
    if(err!=nil){
      panic(err)
    }
  
    customer:=Customer {
      CustID:"CUST1",
      Name:"sarouje",
      Address: "address",
      Contact: Contact {
        Phone:"45533355",
        Email:"someone@gmail.com",
      },
    }    
   }
   err=driver.Insert(customer)
    if(err!=nil){
      panic(err)
    }
   var customers []Customer
  err=driver.Open(Customer{}).Where("name","=","sarouje").Get().AsEntity(&customers)
  if(err!=nil){
    panic(err)
  }
  
  //GET ONE Customer
  //First() will return the first record from the results 
  //AsEntity takes a pointer to Customer variable (not an array pointer)
  var customerFrist Customer
  err=driver.Open(Customer{}).Where("custid","=","CUST1").First().AsEntity(&customerFrist)
  if(err!=nil){
    panic(err)
  }
  
  //Update function uses the ID() to get the Id field/value to find the record and update the data.
  customerFrist.Name="Sony Arouje"
  err=driver.Update(customerFrist)
  if(err!=nil){
    panic(err)
  }
  
  //Delete
  toDel:=Customer{
     CustID:"CUST1",
  }
  err=driver.Delete(toDel)


func tambah(tambahan string) {
  //creates a new Customer file inside the directory passed as the parameter to New()
  //if the Customer file already exist then insert operation will add the customer data to the array
  
  
  //GET ALL Customer
  //opens the customer json file and filter all the customers with name sarouje.
  //AsEntity takes a pointer to Customer array and fills the result to it.
  //we can loop through the customers array and retireve the data.
  
}