#include <stdlib.h>

/**
 * Question:
 * In this kata you are required to, given a string, replace every letter with its position in the alphabet.

 * If anything in the text isn't a letter, ignore it and don't return it.

 * "a" = 1, "b" = 2, etc.

 * Example
 * alphabet_position("The sunset sets at twelve o' clock.")
 * Should return "20 8 5 19 21 14 19 5 20 19 5 20 19 1 20 20 23 5 12 22 5 15 3 12 15 3 11" ( as a string )
*/

// returned string has to be dynamically allocated and will be freed by the caller
char *alphabet_position(const char *text) {
  char* str;
  str = malloc(sizeof(text));
  int i;
 
  // Loop through the string
  for(i = 0; text[i] != '\0'; i++)
    {
    // Check if the character is in Uppercase or Lowercase
    if(text[i] >= 'a' && text[i] <= 'z')
      {
      str[i] = text[i] - 'a';
    } else if(text[i] >= 'A' && text[i] <= 'Z')
      {
      str[i] = text[i] -'A';
    }
  }
  
  return str;
}