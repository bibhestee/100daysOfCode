#include <stdio.h>

/**
 * main - main functions
 * Description: remove unwanted lines that start with a character
 *
 */

int main(void)
{
	FILE *fp = fopen("text.txt", "r");
	FILE *fd = fopen("new.txt", "w");
	char buff[200];
	char c;

	if( !fp || !fd)
	{
		return (-1);
	}

	while(!feof(fp))
	{
		fgets(buff, 200, fp);
		c = buff[0];
		if((c >= 'a' && c <= 'z')||(c >= 'A' && c <= 'Z'))
		{
			continue;
		}
		else
		{
			fputs(buff, fd);
		}
	}

	fclose(fp);
	fclose(fd);
}
