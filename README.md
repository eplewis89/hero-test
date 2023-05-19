# How to run
This program is entirely self contained and can be run without any external dependencies. In order to run this program, first make sure the go runtime is installed. This is an easy to follow process and can be found on the official golang website.

To run, cd into the base folder of this project, and type `go run main.go` - output should be displayed immediately

When done, ctrl + c can be used to cleanly exit the program.

# Prompt
You are the team leader of a very elite, very illustrious, band of super heroes which I assure you are not at all related to any rodent owned comic based team. Your team is very different, it is uhh, D. Revengers? Yeah that's it. D, standing for digital here. Write up some code to help lead your team to victory in the upcoming battle! You'll need to code up all aspects of this scenario including adding some test data in so that it can be run, and will need to factor in a number of different requirements at varying intervals throughout the fight to be successful. How many foes will your team dispatch?

## Requirements for the fight
- **DONE** At a random interval of every 5-10 seconds your arachnidSense() will trigger alerting you of potential new dangers.
- **DONE** One teammember, EagleWoman, can use their keen eyesight to scout the battlefield but it takes them 2 seconds each time to fully scout. Ask them to scout each time your arachnid senses trigger. They can't fight though because they only have a bow and arrow and seriously, what good is that going to do against an army of super villians?
- **DONE** When Eaglewoman scout()s the battlefield then after 2 seconds you should get a list of 1-3 new enemies.
- **DONE** Each new enemy you see from scout() will have the following properties, the value for each chosen at random:
    - enemyType: either "Alien", "Evil Robot", "Lizard Person"
    - powerLevel: 1, 2, or 3
    - timeToArrive: 1, 2, 3, 4, or 5 seconds until the enemy crosses the battlefield and is available to fight
- **DONE** You have two teammates who do all of the fighting. Moar, the god of static electricity from rubbing your socks on carpet, and The Semi-Forgettable Bulk, who transforms when slightly inconvenienced!
- **DONE** Both of these champions are able to dispatch certain types of enemies better than others but their abilities will increase and decrease slightly over the course of the battle. Their total fight time will need to be calculated at the time when they take on each given enemy.
- **DONE** The formulas for approximately how long each hero will take to dispatch a foe are as follows:
    - Moar + Alien: (powerLevel * 3) + random(1 or 2)
    - Moar + Evil Robot: (powerLevel * 1) + random(1, 2, or 3)
    - Moar + Lizard Person: (powerLevel * 2) + random(1 or 2)
    - Bulk + Alien: (powerLevel * 1) + random(0 or 2)
    - Bulk + Evil Robot: (powerLevel * 3) + random(0 or 1)
    - Bulk + Lizard Person: (powerLevel * 2) + random(1 or 2)
- **DONE** Things can get pretty bad if you aren't taking care of these enemies within 10 seconds of them reaching your front lines, make sure you're reprioritizing based on both how many you can dispatch of quickly, as well as making sure that no enemies get past the 10 second mark without being given to Bulk or Moar. Check every second for any enemy that's arrived at your front line and hasn't been given to Bulk or Moar within 10 seconds. If such an enemy exists, end the battle and announce your score!

# Technical requirements
- You can use whatever programming language and whatever libraries you want to complete this exercise.
- Provide the code and any additional data if needed to run the application along with some instructions for what's needed to get setup and running on someone else's machine.
- **DONE** You'll need to implement everything including providing data to your system.
    - For the enemies retuned from scout() method, the list should be generated at random given the constraints above
- **DONE** Moar and Bulk can each only take on one enemy at a time. Both of them should ask your main leader method/class for
a new enemy when they are done fighting their current enemy. At that point they should just be given the next
enemy that's been in the queue for the longest and has also fully crossed the battle field and is available to fight
(see timeToArrive above).
- **DONE** Moar and Bulk both love boasting about their victories, go ahead and print out or otherwise let the person running the application see when they've beaten a foe!
- **DONE** The program should run until one of the enemies has arrived at your frontlines and hasn't been given to Bulk or Moar within 10 seconds. The timer for this critical 10 seconds doesn't start until the timeToArrive for that enemy has elapsed. The final score of your team at the end of the program is the sum of each dispatched enemy's power level. Print out this score before exiting.

# Additional
- This exercise is probably given to people of different skill levels, even if you don't complete it fully please submit what you have written before the deadline.
- Feel free to comment as much code as you want if you think someone might not understand why some code looks the way it does.
- If there are parts of the requirements that are unclear, do your best and add comments with whatever assumptions you are making in place of an official answer.