As i have found out:
it seems that there is a main go routine that fires up when we start the go program

we can create child go routine using the go keyword

it just basically fires up a parallel computing process to execute the block of code it was ment to

if we do not use channels then we cannot communicate between the main goroutine and the child go routines

this is why we need to create channels and pass them onto the process that are being handled by the child goroutines

the main go routine needs to wait and listen for a message from the child goroutine in order to complete its execution

whenever a message comes through from the child goroutines to the main go routine..

the main go routine assumes that all the child go routine process has been completed and 
then the main go rotine completes and exits the program completely. if we do not make the main go routine wait for another channel or stuff like that.