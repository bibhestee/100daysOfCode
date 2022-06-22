/**
 * Create a person object that contains three keys: name, age, and country .
 * Use yourself as an example to set the values for name, age, and country.
 * 
 * Create a function, logData(), that uses the person object to create a 
 * string in the following format:
 * "Name is age years old and lives in Country"
 * 
 * Call the logData() function to verify that it works
 */


// Declare a person object
// let person = {
//     name: "Abdulbasit",
//     age: 20,
//     country: "Nigeria",
//     logData: function ()
//     {
//         console.log(person.name + " is "+ person.age+" years old and lives in "+ person.country+".")
//     }
// }

// person.logData()

// Rock Paper Scissors - Generate Random Number test

// let hands = ["rock", "paper", "scissors"]
// let guess = 1;

// // Generate a random number
// function getHand(){
//     let randomNum = Math.floor(Math.random() * 3 ) + 1
//     return randomNum
// }

// Sorting Fruits

let fruit = ["🍊", "🍎", "🍊", "🍎"]
let appleShelf = document.getElementById("apple-shelf")
let orangeShelf = document.getElementById("orange-shelf")

function sortFruit(){
    for(let i = 0; i < fruit.length; i++)
    {
        if(fruit[i] === "🍊")
        {
            orangeShelf.textContent += "🍊"
        }
        else if(fruit[i] === "🍎")
        {
            appleShelf.textContent += "🍎"
        }
        else{

        }
    }
}
