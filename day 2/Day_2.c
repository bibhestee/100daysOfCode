#include <stdio.h>

/**
 * Comment and code explanations.
 */

int main()
{
    int i;
    int count = 0;
   char str[] = "Hello World!";
   int j = 0;

    while (str[j] != '\0')
    {
        count++;
        j++;
    }

    for (i = 0; i < count; i++)
    {
        putchar(str[i]);
    }

    putc('c')
    return 0;

}
