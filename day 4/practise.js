//Evening Practise

/*
const name = "ABDULLAHI ABDULBASIT";
const age = 20;
const phoneNumber = "08184757979"

let output = `My name is ${name} and I'm ${age} years old. You can contact me on ${phoneNumber}.`;

console.log(output);
*/

//Create an object for the above code

let personalProfile = {
    name: "ABDULLAHI ABDULBASIT",
    age: 20,
    phoneNumber: "08184757979",
    address: "No 31 Erubu lane, off Emir's road, Ilorin, Kwara State.",
    getContact: function()
    {
        let message = `You can reach me on ${this.phoneNumber} or my address: ${this.address}`;
        return message;
    },
    getFullDetails: function()
    {
        let details = `My name is ${this.name} and I'm ${this.age} years old. You can contact me on ${this.phoneNumber} or my address @ ${this.address}`;
        return details;
    }
}

// Render the getContact and getFullDetails method

console.log("Contact Only: "+ personalProfile.getContact());
console.log("Full details: "+ personalProfile.getFullDetails());