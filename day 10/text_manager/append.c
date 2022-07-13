#include <stdio.h>

/**
 * main - main functions
 * Description: add new string to each lines
 *
 */

int main(void)
{
	FILE *fp = fopen("dup.txt", "r");
	FILE*fd = fopen("final2.txt", "w");
	int c;

	if( !fp || !fd)
	{
		return (-1);
	}

	while(!feof(fp))
	{
		c = fgetc(fp);
		if (c == ')')
		{
			fputs(") \n", fd);
		}
		fputc(c, fd);
	}

	fclose(fp);
	fclose(fd);
}
