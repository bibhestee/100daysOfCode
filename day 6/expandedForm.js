/**
 * Description: You will be given a number and you will need to return it as a string in Expanded Form. 
 * For example:
 * expandedForm(12); // Should return '10 + 2'
 * expandedForm(42); // Should return '40 + 2'
 * expandedForm(70304); // Should return '70000 + 300 + 4'
 * NOTE: All numbers will be whole numbers greater than 0.
 * @param {*} num 
 * @returns expandedForm
 */
function expandedForm(num) {
    // Convert the num to string
   let ch = num.toString()
   
   let output = ''
   let newStr = ''
   let j = ch.length - 1
   for(let i = 0; i < ch.length; i++)
   {
      if(ch[i] !== '0')
       {
         output += ch[i] + '0'.repeat(j)
         // Check if not at last string and str length is more than 1 
        if(i < ch.length - 1 && ch.length > 0)
         {
           output += " + "
         }
       }
       j--;
    }
   let k = output.length - 1
   // Remove last three characters " + "
   if(output[k] === " ")
     {
       newStr += output.slice(0, -3)
       return newStr
     }
    return output
}
// Testing
console.log(expandedForm(98219200))