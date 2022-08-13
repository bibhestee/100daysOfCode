#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/**
 * main - main body of the file
 * @argc: argument counts
 * @argv: argument vector
 * Description: This file reads input from the command line 
 * 	./word_rm filename string
 *  where filename is the file to be modified
 *	string is the word to be removed(must be inside quotes)
 *
 * Return: 0 if success, and negative if otherwise
 */


char *strremove(char *str, const char *sub);

int main(int argc, char **argv)
{
        FILE *file, *new_file;
        char *buff, *output;
        size_t n = 0;

        if (argc != 3)
        {
                printf("Usage: ./word_rm <file.txt> <text>\n");
                exit(EXIT_FAILURE);
        }

        file = fopen(argv[1], "r");
        new_file = fopen("removed.txt", "w");

        if (!file || !new_file)
        {
                printf("File can't be open\n");
                exit(EXIT_FAILURE);
        }
        const char *word = argv[2];

        while(getline(&buff, &n , file) != -1)
        {
                output = strremove(buff, word);
                if (output != NULL)
                        fprintf(new_file, "%s", output);
        }

        fclose(file);
        fclose(new_file);

        exit(EXIT_SUCCESS);
}

/**
 * strremove - remove a specified word from the sentence input
 * @str: word to be removed
 * @sub: The sentence
 * Description: Each occurrence of the word is removed from the
 *  sentence and modified sentence is returned
 *
 * Return: modified sentence
 */

char *strremove(char *str, const char *sub) {
    char *p, *q, *r;
    if (*sub && (q = r = strstr(str, sub)) != NULL) {
        size_t len = strlen(sub);
        while ((r = strstr(p = r + len, sub)) != NULL) {
            while (p < r)
                *q++ = *p++;
        }
        while ((*q++ = *p++) != '\0')
            continue;
    }
    return str;
}
