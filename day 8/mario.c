#include <cs50.h>
#include <stdio.h>

void add_block(int n);

int main(void)
{
    int height;

    // Prompt user for input
    do{
        height = get_int("Height: ");
    } while(height < 1 || height > 8);

    // Start the building with the height specified
    for(int row = 0, Tab = height - 1; row < height; row++)
    {
        // Right alignment
        for(int i = Tab; i > 0; i--)
        {
            printf(" ");
        }

        // Left Block
        add_block(row);

        //Gap
        printf("  ");

        // Right Block
        add_block(row);

        printf("\n");
        Tab--;
    }
}


void add_block(n)
{
    for(int i = 0; i <= n; i++)
    {
        printf("#");
    }
}