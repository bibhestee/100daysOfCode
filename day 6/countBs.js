/**
 * countBs -  Takes a string and returns the total 
 * number of occurrence.
 * 
 * @param {*} str 
 * @returns Total count
 */

function countBs(str)
{
    let count = 0;

    for(let i = 0; i < str.length; i++)
    {
        if(str[i] === 'B')
        {
            count++;
        }
    }

    return count;
}
console.log(countBs("Beans is Better But not the Best."))

