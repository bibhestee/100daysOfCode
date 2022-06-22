function loopTriangle(n){
for (let i = 1; i < n +1; i++)
{
	let str = ""
	for(let j = 0; j < i; j++)
	{
		 str += "#"
	}
	console.log(str)
}
}

loopTriangle(5)
loopTriangle(7)
