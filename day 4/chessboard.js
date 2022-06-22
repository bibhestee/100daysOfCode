function chessBoard(m, n)
{
	for(let i = 0; i < m; i++)
	{
		let str = ""
		for(let j = 0; j < n; j++)
		{
			str += "# "
		}
		console.log(str)
		console.log(str.split("").reverse().join(""))
	}
}


// Create 8x8 chessboard
chessBoard(8, 8)
