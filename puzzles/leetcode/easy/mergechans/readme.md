Topic Given n channels of type chan int. And merge all the data from these channels into one and return it.

How to do it? Write a function that will asynchronously read from the source channels that will be passed to it as
arguments and write to the resulting channel that will return from the function.

Step1:  Create a channel where we will merge all the data. Explanation: It will be unbuffered because we don’t know how
much data will come from the channels.

Step2: Asynchronously read from the source channels and close the resulting channel for the merge when all reading is
over.

Step3: Wait for the end of reading, we simply wrap this loop through the channels in a wait group.