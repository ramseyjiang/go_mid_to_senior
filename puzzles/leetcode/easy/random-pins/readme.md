Write a library for generating random PIN codes. You probably know what a PIN code is; it’s a short sequence of numbers,
often used as a passcode for bank cards.

The library should export a function that returns a batch of 1,000 PIN codes in random order Each PIN code in the batch
should be unique Each PIN should be:
4 digits long Two consecutive digits should not be the same (e.g. 1156 is invalid)
Three consecutive digits should not be incremental (e.g. 1236 is invalid)
The library should have automated tests.
