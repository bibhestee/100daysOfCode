#include <stdio.h>
#include <stdlib.h>

// Incomplete...

char* spinWords(const char* sentence)
{
    size_t i;

    for(i = 0; i < strlen(sentence); i++)
    {
        if(isspace(sentence[i]))
        {
            return "ok";
        }
    }
    return "nothing";
}

int main(void)
{
    char* ch;

    ch = spinWords("boy");

    printf("%s\n", ch);
    return 0;
}