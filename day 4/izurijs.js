// Conditional statements
function calc(x, operator, y){
  if(operator === "*")
  {
    return x * y
  } else if (operator === "+") {
    return x + y
  } else if (operator === "-") {
    return x - y
  } else if (operator === "/") {
    return x/y
  }
}

console.log(calc(12, "+", 9))
