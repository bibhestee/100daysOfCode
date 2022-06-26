#include <cs50.h>
#include <stdio.h>

bool valid_checker(long number);
void check_card(long number);
bool visa(long number);
bool mc(long number);
bool amex(long number);


int main(void)
{
    long number;
    do
    {
        // Prompt user for input
        number = get_long("Number: ");
    }
    while (number < 0);

    if (valid_checker(number))
    {
        check_card(number);
    }
    else
    {
        printf("INVALID\n");
    }
}

/**
 * valid_checker - checks if a card detail is valid
 *
 * Return: returns true if valid or INVALID
 *
 */
bool valid_checker(long number)
{
    int digit, sum = 0;

    for (int i = 0; number / 10 > 0 ; i++)
    {
        // Alternate number
        digit = (number % 100) / 10;
        // Add mod
        sum += (number % 100) % 10;
        // Multiply alternate by 2
        if ((digit * 2) >= 10)
        {
            sum += (digit * 2) - 9;
        }
        else
        {
            sum += digit * 2;
        }
        number /= 100;
        // Add Odd digits last number i.e 13 and 15 digits
        if (number < 10)
        {
            sum += number;
        }
    }
    // Use checksum to verify validity
    if (sum % 10 == 0)
    {
        return true;
    }
    else
    {
        return false;
    }
}

/**
 * check_card - checks the card type
 *
 * Return: Nothing
 */
void check_card(long number)
{
    if (amex(number))
    {
        printf("AMEX\n");
    }
    else if (mc(number))
    {
        printf("MASTERCARD\n");
    }
    else if (visa(number))
    {
        printf("VISA\n");
    }
    else
    {
        printf("INVALID\n");
    }
}

/**
 * visa - checks if the card number is Visa
 *
 * Return: returns true if the card is Visa
 */
bool visa(long number)
{
    long visa_l = 4000000000000;
    long visa_h = 5000000000000000;
    if (number >= visa_l && number < visa_h)
    {
        for (int i = 0; i <= 15; i++)
        {
            number = number / 10;
            if (number < 10 && number == 4)
            {
                return true;
            }
            else if (number < 10 && number != 4)
            {
                return false;
            }
        }
    }
    return false;
}

/**
 * mc - checks if the card number is Mastercard
 *
 * Return: returns true if the card is Mastercard
 */
bool mc(long number)
{
    long mc_l = 5100000000000000;
    long mc_h = 5600000000000000;
    if (number >= mc_l && number < mc_h)
    {
        return true;
    }
    return false;
}

/**
 * amex - checks if the card number is American Express
 *
 * Return: returns true if the card is American Express
 */
bool amex(long number)
{
    long amex_l = 340000000000000, amex_ll = 350000000000000;
    long amex_h = 370000000000000, amex_hh = 380000000000000;
    if ((number >= amex_l && number < amex_ll) || (number >= amex_h && number < amex_hh))
    {
        return true;
    }
    return false;
}