function solution(string)
{
    let chArray = string.split("");
    let newStr = '';
    const l = chArray.length;

    for(let i = 0; i < l; i++)
    {
        if(chArray[i] >= 'A' && chArray[i] <= 'Z')
        {
            // Replace the Uppercase with space + uppercase
            chArray[i] = ` ${chArray[i]}`
            newStr = chArray.join("")
        } else {
            newStr = chArray.join("")
        }
    }
    return newStr
}

console.log(solution("big"))
console.log(solution('camelCasing'))

// Alternative
// function solution(string) {
//     return(string.replace(/([A-Z])/g, ' $1'));
  
//   }