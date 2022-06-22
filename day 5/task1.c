#include <stdbool.h>

/**
 * Question:
 * A pangram is a sentence that contains every single letter of the alphabet at least once. For example, the sentence "The quick brown fox jumps over the lazy dog" is a pangram, because it uses the letters A-Z at least once (case is irrelevant).

Given a string, detect whether or not it is a pangram. Return True if it is, False if not. Ignore numbers and punctuation.

*/


bool is_pangram(const char *str_in) {

    //  <----  hajime!
  int i, a, b = 0;
  char save [26];
  save[0] = '0';

  for (i = 0; str_in[i] != '\0'; i++)
    {
    if(str_in[i] >= 'a' && str_in[i] <= 'z')
    {
        b += str_in[i] - 'a';
    } else if (str_in[i] >= 'A' && str_in[i] < 'Z')
    {
        b += str_in[i] - 'A';
    }
    
        
  }
  if(b == 26)
    {
    return true;
  } else {
    return false;
  }
}
