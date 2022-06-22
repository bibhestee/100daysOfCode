// Write a function that takes an integer as input, 
// and returns the number of bits that are equal to one in the binary representation of that number. 
// You can guarantee that input is non-negative.

// Example: The binary representation of 1234 is 10011010010, so the function should return 5 in this case

// var countBits = function(n) {
//     // Program Me
//     let bit = "";
//     while(n !== 0)
//     {
//         bit += n % 2;
//         n = n/2;
//     }
   
//   };

function countBits(n) {
    // Program Me
    let bit = "";
    for(let i = 0; n !== 0; )
    {
        bit += n % 2;
        n = n/2;
    }
    return bit;
};

console.log(countBits(10));