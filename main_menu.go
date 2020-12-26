package main

import (
   "fmt"
   "strconv"
)

func main() {

    
    valid_choices := []int {1,2,3,4,9}
    var main_choice_S string 
    var main_choice_I int 

    var main_menu string =   
    " \n"+
    " |---------------------------------------------|\n"+
    " |++ MAIN MENU ++                              |\n"+
    " |~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~|\n"+
    " |(1) Create Flightplan                        |\n"+
    " |(2) List current Flightplans                 |\n"+
    " |(3) Adapt Flightplan                         |\n"+
    " |(4) Some other function                      |\n"+
    " |                                             |\n"+
    " |(9) Quit                                     |\n"+
    " |                                             |\n"+
    " |---------------------------------------------|\n";
    
        for
            { 
            fmt.Print(main_menu)
            fmt.Print("Enter choice (confirm with ENTER):    ")
            fmt.Scanln(&main_choice_S)  
            
            fmt.Print("You entered ") 
            fmt.Print(main_choice_S)   
            main_choice_I, _ = strconv.Atoi(string(main_choice_S))
			fmt.Print(" which was converted to ")
			// do not use 0 in the menu as strings are converted to 0  
			fmt.Print(main_choice_I)
            
            if (Contains(valid_choices,main_choice_I)){
                break             
            } else {
                ClearScreen(80)
                
                fmt.Print("Your previous choice ")
                fmt.Print(main_choice_S) 
                fmt.Print(" was not valid, \n")
   
                fmt.Print("\n Valid choices are ")
                fmt.Print(valid_choices)
            }
        }   
        }

        func Contains(list []int, x int) bool {
            for _, item := range list {
                if item == x {
                    return true
                }
            }
            return false
		}
		
		func ClearScreen(n int){
		// Simple function to print 80 empty lines
			for i:=0 ; i<n;i++{
			fmt.Println()
			}
		}
