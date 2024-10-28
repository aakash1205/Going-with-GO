package main

import "fmt"

func main() {
	var name string = "Aakash"
	fmt.Println(name);
	fmt.Printf("variable is of type: %T \n",name);

	var isLoggedIn bool =false
	fmt.Println(isLoggedIn);
	fmt.Printf("variable is of type: %T \n",isLoggedIn);
   // checking uint8 =2^8-1 == 0-255 
	var small uint8 =255 
	fmt.Println(small);
	fmt.Printf("variable is of type: %T \n",small);


	// float32 gives 5 decimals 

	var  decimals float32=45.4567890875;
	fmt.Println(decimals);
	fmt.Printf("variable is of type: %T \n",decimals);

	// checking default initialization of GO
	var  num int; // here it initializes num as 0 
	fmt.Println(num);
	fmt.Printf("variable is of type: %T \n",num);

	var  s string; // here it initializes s as nothing 
	fmt.Println("string is initialized as: ",s);
	fmt.Printf("variable is of type: %T \n",s);

	//implicit style (not specifying the type)

	var website="https://localhost:3000";
	fmt.Println(website);
	fmt.Printf("its type is: %T\n",website)

	// no var style

	newvar := 30000

	fmt.Println("newvar value is",newvar)
	fmt.Printf("type is: %T \n",newvar)

}
