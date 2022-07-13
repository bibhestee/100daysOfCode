#include <stdio.h>

/**
 * main - main functions
 * Description: remove unwanted lines that start with a character
 *
 */

int main(void)
{
	FILE *fp = fopen("new.txt", "r");
	FILE*fd = fopen("text.txt", "w");
	char buff[100];
	char c;

	if( !fp || !fd)
	{
		return (-1);
	}

	while(!feof(fp))
	{
		fgets(buff, 100, fp);
		c = buff[0];
		if(c == '0')
		{
			fputs(buff, fd);
		}
	}

	fclose(fp);
	fclose(fd);
}
