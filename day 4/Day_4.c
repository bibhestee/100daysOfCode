// Learning Pointers, Arrays and Memory Allocations

#include <stdio.h>

int StringPrint(char *c)
{
    printf("%s", c);
    printf("\n");
    return 0;
}
int main()
{
    /*
    1. Given a character string, I like C!, write a program to pass the string to a function
    that displays the string on the screen.
    */

    char str[] = "I like C!";
    int *(ptr_str)(char str[]);

    *ptr_str = StringPrint;

    (*ptr_str)(str);

    

    
    /*
    2. Rewrite the program in exercise 1. This time, change the string of I like C! to I
    love C! by moving a pointer that is initialized with the start address of the string
    and updating the string with new characters. Then, pass the updated string to the
    function to display the content of the string on the screen.
    

    char ch[] = "I like C!";
    char *ptr_ch;

    printf("Before assigning new value\n");
    printf("ch is: %s\n", ch);

    ptr_ch = ch;

    printf("After assigning new value\n");
    printf("ch is: %s", ch);
    
    */
    
    /*
    3. Given a two-dimensional character array, str, that is initialized as
    char str[2][15] = { “You know what,”, “C is powerful.” };
    write a program to pass the start address of str to a function that prints out the
    content of the character array.
    4. Rewrite the program in Listing 16.7. This time, the array of pointers is initialized
    with the following strings:
    “Sunday”, “Monday”, “Tuesday”, “Wednesday”, “Thursday”, “Friday”, and
    “Saturday”.
    */
    return (0);
}