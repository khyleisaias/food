package main

import (
    "fmt"
    "strings"
    "strconv"
    "log"
    "os"
)

var (
    printf = fmt.Printf
    scanf = fmt.Scanln

    price float64
)

func getOrder(ch []string) {
    var prods = make(map[string]string)
    var prods2 = make(map[string][]string)
    var _price = &price
    
    for i := 0; i < len(ch); i++ {
        c, e := strconv.Atoi(ch[i])
        if e != nil {
            log.Fatal(e)
        }
        // printf("%d\n", c)
        switch c {
            case 1:
                var p float64 = 5.0
                var quantity int 
                prods["order"] = "Burger"
                printf("Burger: Enter quantity: ")
                scanf(&quantity)
                var f = strconv.FormatFloat((p*float64(quantity)), 'f', 2, 64)
                
                *_price += p*float64(quantity)
                // prods2["what"] = append(prods2["what"], "%spc %s (%s)", strconv.Itoa(quantity), prods["order"], string(f))
                prods2["what"] = append(prods2["what"], strconv.Itoa(quantity) + "pc " +  prods["order"] + " ($" + f + ")")
            case 2:
                var p float64 = 3.0
                var quantity int
                prods["order"] = "Fries"
                printf("Fries: Enter quantity: ")
                scanf(&quantity)
                var f = strconv.FormatFloat((p*float64(quantity)), 'f', -1, 64)
                
                *_price += p*float64(quantity)
                prods2["what"] = append(prods2["what"], strconv.Itoa(quantity) + "pc " +  prods["order"] + " ($" + f + ")")
            case 3:
                var p float64 = 4.5
                var quantity int
                prods["order"] = "Salad"
                printf("Salad: Enter quantity: ")
                scanf(&quantity)
                var f = strconv.FormatFloat((p*float64(quantity)), 'f', -1, 64)
                
                *_price += 4.5*float64(quantity)
                prods2["what"] = append(prods2["what"], strconv.Itoa(quantity) + "pc " +  prods["order"] + " ($" + f + ")")
            case 4:
                var p float64 = 2.5
                var quantity int
                prods["order"] = "Pizza Slice"
                printf("Pizza Slice: Enter quantity: ")
                scanf(&quantity)
                var f = strconv.FormatFloat((p*float64(quantity)), 'f', -1, 64)
                
                *_price += p*float64(quantity)
                prods2["what"] = append(prods2["what"], strconv.Itoa(quantity) + "pc " +  prods["order"] + " ($" + f + ")")
            case 5:
                var p float64 = 1.5
                var quantity int
                prods["order"] = "Soda"
                printf("Soda: Enter quantity: ")
                scanf(&quantity)
                var f = strconv.FormatFloat((p*float64(quantity)), 'f', -1, 64)
                
                *_price += p*float64(quantity)
                prods2["what"] = append(prods2["what"], strconv.Itoa(quantity) + "pc " +  prods["order"] + " ($" + f + ")")
            case 6:
                var p float64 = 3.5
                var quantity int 
                prods["order"] = "Milkshake"
                printf("Milkshake: Enter quantity: ")
                scanf(&quantity)
                var f = strconv.FormatFloat((p*float64(quantity)), 'f', -1, 64)
                
                *_price += p*float64(quantity)
                prods2["what"] = append(prods2["what"], strconv.Itoa(quantity) + "pc " +  prods["order"] + " ($" + f + ")")
            case 0:
                printf("Thank you!!\n")
                os.Exit(0)
            default:
                printf("Illegal Order!!\n")
                    showMenu()
        }
    }
    printf("\nHere's what you ordered:\n")
    for j := 0; j < len(prods2["what"]); j++ {
        //printf("%d\n", j)
        printf("\t%s\n", prods2["what"][j])
    }
    printf("----------------------------------------------------\n")
    if *_price >= 20 {
        var orig_price = *_price
        *_price = orig_price*0.90
        printf("total: $%.2f with 10%% discount: $%.2f\n", orig_price, *_price)
    }
}

func showMenu() {
    printf("Welcome!\n")
    printf("Choose a food below (mulitple choices are separated with commas, 0 for exit.)\n")
    printf("1. Burger       - $5\n")
    printf("2. Fries        - $3\n")
    printf("3. Salad        - $4.50\n")
    printf("4. Pizza Slice  - $2.50\n")
    printf("5. Soda         - $1.50\n")
    printf("6. Milkshake    - $3.50\n")
    var choice string
    printf(">> ") // This is the prompt
    scanf(&choice)

    var choices = strings.Split(choice, ",") // This splits the orders by commas
    // printf("%s\n", choices);
    getOrder(choices) // This gets the order. see getOrder in line 8.

}

func main() {
    showMenu()
}
