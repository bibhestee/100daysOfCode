<h1># Day 8</h1>

<h2>Languages</h2>
<div>
- C
</div>


<h1> Tasks </h1>
<h3>1 - Mario Bros - CS50</h3>
Toward the beginning of World 1-1 in Nintendo’s Super Mario Brothers, Mario must hop over adjacent pyramids of blocks, per the below.
<img src="https://cs50.harvard.edu/x/2022/psets/1/mario/more/pyramids.png" alt="mario.jpg" style="width: 50px; height: 50px">
<p>Let’s recreate those pyramids in C, albeit in text, using hashes (#) for bricks, a la the below. Each hash is a bit taller than it is wide, so the pyramids themselves will also be taller than they are wide.</p>
<div>
   #  #
  ##  ##
 ###  ###
####  ####
</div>
The program we’ll write will be called mario. And let’s allow the user to decide just how tall the pyramids should be by first prompting them for a positive integer between, say, 1 and 8, inclusive.

Here’s how the program might work if the user inputs 8 when prompted:

<div>
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
</div>
