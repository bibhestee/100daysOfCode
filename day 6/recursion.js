/**
 * Check whether a number is even or odd using recursion 
 *
 * isEven - checks if the input is even or odd
 * @n: Number 
 * Return: True if its even and False if its odd
 */

function isEven(n)
{    if(n < -1)
    {
        // Return false if input is negative
        return false;
    } else if (n - 2 === 0) {
        return true;
    } else if (n - 2 === 1) {
        return false;
    } else {
        // Return the result after making the function 
        return isEven(n - 2);
    }
}
// For 50
console.log(isEven(50))
// For 75
console.log(isEven(75))
// For -1
console.log(isEven(-1))
