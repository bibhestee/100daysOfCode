#include <stdlib.h>
#include <stdio.h>

char Food_type(char *food);

int main()
{
    int food_code;
    char food;
    char menu[] = "Rice, Beans, Plantain, Egg, Fish";
    

    printf("Enter the following numbers for each food items.\n");
    printf("Rice - 1\n Beans - 2\n Egg - 3\n");

    printf("Input the food items number: ");
    scanf("%d", food_code);

    food = menu[food_code];
    
    printf("The selected food is %s",Food_type(food));
    return 0;
}

char Food_type(char *food)
{
    if (*(food) == 'Rice')
    {
        return 1;
    }
    else if (*(food) == 'Beans' || *(food) == 'Egg')
    {
        return 2;
    }    
    
}