// Learning getc(), getchar(), putc() and putchar().
// Pointers
// Arrays and Multidimensional array
#include <stdio.h>

/**
 *  Using the getc(), getchar(), putc() and putchar() functions.
 * getc() - gets a character from standard input and converts it to integer then it returns the ASCII value of 
 * the character.
 * getchar() - works like getc() but its the equivalent of getc(stdin)
 * putc(int c, FILE *stream) - gets an ASCII value and converts it to a character then returns the character to the file stream
 * putchar(int c) - accept only one argument and returns the character set for the ASCII value
 */

int main()
{
    // INPUT AND OUTPUT FILE HANDLING
    //1. Write a program to put the characters B, y, and e together on the screen.

    //putchar(66); //Capital A is 65 and small a is 97 and 10 for \n
    //putchar(121);
    //putchar(101);

    //2. Display the two numbers 123 and 123.456 and align them at the left edge of the field.
    /*
    int x;
    double y;

    x = 123;
    y = 123.456;

    printf("\n%d\n", x);
    printf("%lf", y);
    */

    //3. Given three integers, 15, 150, and 1500, write a program that prints the integers on the screen in the hex format.

    /*
    printf("\n%x\n", 15);
    printf("%x\n", 150);
    printf("%x", 1500);
    */

    //4. Write a program that uses getchar() and putchar() to read in a character entered by the user and write the character to the screen.

    /*
    printf("\nEnter a character: ");
    int x = getchar();
    putchar(x);
    */

   // POINTERS
   /*
    1. Given three integer variables, x = 512, y = 1024, and z = 2048, write a program
    to print out their left values as well as their right values.
  
    int x = 512;
    int y = 1024;
    int z = 2048;

    printf("\nThe left value of x is %p and the right value of x is %d\n", &x, x);
    printf("The left value of y is %p and the right value of y is %d\n", &y, y);
    printf("The left value of z is %p and the right value of z is %d\n", &z, z);

  
    2. Write a program to update the value of the double variable flt_num from 123.45
    to 543.21 by using a double pointer.
 

    double flt_num = 123.45;
    double *ptr_flt_num = &flt_num;

    printf("\nThe value of flt_num is: %lf\n", flt_num);
    *ptr_flt_num = 543.21;

    printf("The value of flt_num is: %lf", *ptr_flt_num);
  
    3. Given a character variable ch and ch = ‘A’, write a program to update the value of
    ch to decimal 66 by using a pointer.

    char ch;
    ch = 'A';
    char *ptr_ch = &ch;

    printf("\nThe value of ch is: %c\n", ch);
    *ptr_ch = 66;

    printf("The value of ch is: %c", *ptr_ch);

    4. Given that x=5 and y=6, write a program to calculate the multiplication of the two
    integers and print out the result, which is saved in x, all in the way of indirection
    (that is, using pointers).

    int x = 5;
    int y = 6;
    int *ptr_x = &x;
    int *ptr_y = &y;

    printf("\nThe product of x and y using pointers: %d", *ptr_x * *ptr_y);
    */

    // ARRAYS AND MULTIDIMENSIONAL ARRAYS
    /*
    1. Given this character array:
    char array_ch[5] = {‘A’, ‘B’, ‘C’, ‘D’, ‘E’};
    write a program to display each element of the array on the screen.
   

    char array_ch[5] = {'A', 'B', 'C', 'D', 'E'};

  
    for (int i = 0; i < sizeof(array_ch); i++)
    {
        printf("array_ch[%d]: %c\n", i, array_ch[i]);
    }
  
  
    2. Rewrite the program in Exercise 1, but this time use a for loop to initialize the
    character array with ‘a’, ‘b’, ‘c’, ‘d’, and ‘e’, and then print out the value of
    each element in the array.
  
    char array_ch[5] = {'a', 'b', 'c', 'd', 'e'};

  
    for (int i = 0; i < sizeof(array_ch); i++)
    {
        printf("array_ch[%d]: %c\n", i, array_ch[i]);
    }
 
    3. Given this two-dimensional unsized array:
    char list_ch[][2] = {
    ‘1’, ‘a’,
    ‘2’, ‘b’,
    ‘3’, ‘c’,
    ‘4’, ‘d’,
    ‘5’, ‘e’,
    ‘6’, ‘f’};
    write a program to measure the total bytes taken by the array, and then print out all
    elements of the array.
   

    char list_ch[][2] = {'1', 'a',
    '2', 'b',
    '3', 'c',
    '4', 'd',
    '5', 'e',
    '6', 'f'};

    printf("The total bytes taken by the array is: %d", sizeof(list_ch));
    for (int i = 0; i < 10; i++)
    {
        printf("\n");
        for (int j = 0; j < 2; j++)
        {
            printf("%c", list_ch[i][j]);
        }
        
    }
   


    4. Rewrite the program in Listing 12.5. This time put a string of characters, I like
    C!, on the screen.
    5. Given the following array:
    double list_data[6] = {
    1.12345,
    2.12345,
    3.12345,
    4.12345,
    5.12345};
    use the two equivalent ways taught in this lesson to measure the total memory
    space taken by the array, and then display the results on the screen.
   

    double list_data[6] = {
    1.12345,
    2.12345,
    3.12345,
    4.12345,
    5.12345};
   
    printf("\nUsing Method I: Total memory taken by array is %d\t", sizeof(list_data));
    printf("Using Method II: Total memory taken by array is %d\n", sizeof(double)* (sizeof(list_data)/sizeof(list_data[0])));
    */
    return (0);
}
