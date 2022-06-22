#include "stdio.h"
#include "string.h"
#include "stdlib.h"

/*
 * Built-in libraries are included already
 */
int main(void) {
  char* a;
  size_t i;
  int count = 0;
  printf("My first output on Atom editor\n");

  printf("Enter a string:");
  scanf("%s\n", &a);

  for (i = 0; a[i] != '\0'; i++) {
    /* code */
    count++;
  }
  printf("%d\n", count);
  return 0;
}
