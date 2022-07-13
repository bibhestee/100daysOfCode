#include <stdio.h>

/**
 * main - main functions
 * Description: remove dup lines
 *
 */

int main(void)
{
	FILE *fp = fopen("text.txt", "r");
	FILE*fd = fopen("dup.txt", "w");
	char buff[200];
	char a;

	if( !fp || !fd)
	{
		return (-1);
	}

	while(!feof(fp))
	{
		fgets(buff, 200, fp);
		if(a != buff[3])
		{
			fputs(buff, fd);
			a = buff[3];
		}
	}

	fclose(fp);
	fclose(fd);
}
