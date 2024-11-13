"""
    FOOD ORDRING SYSTEM
"""
# Set the price variable
price = 0

# This function gets the order.
def getOrder(order, _price=price):
    # The product dictionary
    prods = {
        "what": []
    }
    # Iterate through the orders
    for i in range(len(order)):
        # Check the orders
        match int(order[i]):
           case 1:
               quantity = int(input("Burger: Enter quantity: "))
               prods["order"] = "Burger"
               _price += 5*quantity
               prods["what"].append(f"{quantity}pc {prods['order']} (${5*quantity})")
           case 2: 
               quantity = int(input("Fries: Enter quantity: "))
               prods["order"] = "Fries"
               _price += 3*quantity
               prods["what"].append(f"{quantity}pc {prods['order']} (${3*quantity})")
           case 3:
               quantity = int(input("Salad: Enter quantity: "))
               prods["order"] = "Salad"
               _price += 4.5*quantity
               prods["what"].append(f"{quantity}pc {prods['order']} (${4.5*quantity})")
           case 4:
               quantity = int(input("Pizza Slice: Enter quantity: "))
               prods["order"] = "Pizza Slice"
               _price += 2.5*quantity
               prods["what"].append(f"{quantity}pc {prods['order']} (${2.5*quantity})")
           case 5:
               quantity = int(input("Soda: Enter quantity: "))
               prods["order"] = "Soda"
               _price += 1.5*quantity
               prods["what"].append(f"{quantity}pc {prods['order']} (${1.5*quantity})")
           case 6:
               quantity = int(input("Milkshake: Enter quantity: "))
               prods["order"] = "Milkshake"
               _price += 3.5*quantity
               prods["what"].append(f"{quantity}pc {prods['order']} (${3.5*quantity})")
           case 0:
               print("Thank you!")
               exit(0)
           case _:
               print("Illegal order!")
               showMenu()
    # Display what you ordered 
    print("\nHere's what you ordered: ")
    # Iterate through the order list
    for w in prods["what"]:
        print(f"\t{w}")
    print("---------------------------------------------------")
    # Check if price is ltoe (less than or equal to) 20
    if (_price >= 20):
        orig_price = _price
        _price = _price*0.10
        print(f"total: ${orig_price} with 10% discount: ${round(_price, 2)}")
    else:
        print(f"total: ${_price}")

# This shows the menu and prompts you
def showMenu():
    print("Welcome!")
    print("Choose a food below (mulitple choices are separated with commas, 0 for exit.) ")
    print("1. Burger       - $5")
    print("2. Fries        - $3")
    print("3. Salad        - $4.50")
    print("4. Pizza Slice  - $2.50")
    print("5. Soda         - $1.50")
    print("6. Milkshake    - $3.50")
    choice = input(">> ") # This is the prompt

    choices = choice.split(",") # This splits the orders by commas

    getOrder(choices) # This gets the order. see getOrder in line 8.


showMenu() # we will show the menu by default.
