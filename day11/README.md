# Day 11 of 100 days of coding

# Language

- C programming

# Project


## Word_remover

### About:
 
  This project is used for text file modification. 
  
  
  A sentence or word can be removed from the file easily using this file. 
  
  
  This will be useful for handling repetitive task of removing words line by line or searching through the file to remove each occurence of the word.

  
### How To:
- First, clone the repository on your local machine using the command below:

    ``` 
        git clone github.com/bibhestee/100daysOfCode/day11/ 
    ```

- Compile the program.

    ``` 
        gcc word_remover.c -o word_rm 
    ```

- Congratulations 🎉🎉🎉.. You have successfully install the program.

### Run the program: 

- Move your file to this current directory or create new file then copy and paste the contents.

    ``` 
        bibest@DESKTOP-TSF2R7V:/100daysOfCoding/day11/ cat > new_file.txt
        paste the content here then use ctrl + d to exit writing mode 
    ```
   
- Your new file is saved, next is to run the program.

- Using this format: ./word_rm new_file.txt "the".
  
  new_file.txt is my file and "the" is the word to be removed

- If successful, a new file is created automatically with the name `removed.txt`.
  The file contains the modified file.

- View the file to show the changes made.

    ` cat removed.txt `

- Example:

      
        bibest@DESKTOP-TSF2R7V:/100daysOfCoding/day11/ cat > new_file.txt

        paste the content here then use ctrl + d to exit writing mode

        bibest@DESKTOP-TSF2R7V:/100daysOfCoding/day11/ ./word_rm new_fiie.txt "the"

        bibest@DESKTOP-TSF2R7V:/100daysOfCoding/day11/ ls

        bibest@DESKTOP-TSF2R7V:/100daysOfCoding/day11/ new_file.txt removed.txt word_rm word_remover.c

        bibest@DESKTOP-TSF2R7V:/100daysOfCoding/day11/ cat removed.txt

        bibest@DESKTOP-TSF2R7V:/100daysOfCoding/day11/ paste content here then use ctrl + d to exit writing mode 
      


Hurray! 🎈🎈🎉🎉 Congratulations, you have successfully removed each occurrence of the word from the file.
