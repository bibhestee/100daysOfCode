# Day 8

## Languages
```
- C
```

## Tasks 
### 1 - Mario Bros - CS50
Toward the beginning of World 1-1 in Nintendo’s Super Mario Brothers, Mario must hop over adjacent pyramids of blocks, per the below.
![mario](https://cs50.harvard.edu/x/2022/psets/1/mario/more/pyramids.png)

Let’s recreate those pyramids in C, albeit in text, using hashes (#) for bricks, a la the below. Each hash is a bit taller than it is wide, so the pyramids themselves will also be taller than they are wide.

The program we’ll write will be called mario. And let’s allow the user to decide just how tall the pyramids should be by first prompting them for a positive integer between, say, 1 and 8, inclusive.
Here’s how the program might work if the user inputs 8 when prompted:
```
$ ./mario
Height: 8
       #  #
      ##  ##
     ###  ###
    ####  ####
   #####  #####
  ######  ######
 #######  #######
########  ########
```
Here’s how the program might work if the user inputs 4 when prompted:
```
$ ./mario
Height: 4
   #  #
  ##  ##
 ###  ###
####  ####
```
Here’s how the program might work if the user inputs 2 when prompted:
```
$ ./mario
Height: 2
 #  #
##  ##
```

### 2 - Credit Card Checker - CS50
A credit (or debit) card, of course, is a plastic card with which you can pay for goods and services. Printed on that card is a number that’s also stored in a database somewhere, so that when your card is used to buy something, the creditor knows whom to bill. There are a lot of people with credit cards in this world, so those numbers are pretty long: American Express uses 15-digit numbers, MasterCard uses 16-digit numbers, and Visa uses 13- and 16-digit numbers. 
All `American Express` numbers start with `34 or 37`; most `MasterCard numbers` start with `51, 52, 53, 54, or 55` (they also have some other potential starting numbers which we won’t concern ourselves with for this problem); and all `Visa` numbers start with `4`. But credit card numbers also have a “checksum” built into them, a mathematical relationship between at least one number and others. That checksum enables computers (or humans who like math) to detect typos (e.g., transpositions), if not fraudulent numbers, without having to query a database, which can be slow. Of course, a dishonest mathematician could certainly craft a fake number that nonetheless respects the mathematical constraint, so a database lookup is still necessary for more rigorous checks.
The task is solved using Luhn's Algorithm

```
$ ./credit

Number: 4003600000000014
 
VISA
```
