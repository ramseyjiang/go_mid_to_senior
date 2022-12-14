Split the processes into several goroutines

1. while not creating a new goroutine every time, but simply reusing the existing ones.
2. To do 1, create a channel with jobs and the resulting channel.
3. For each worker, it will create a goroutine that will wait for a new job, apply the given function to it, and fire
   the response into the resulting channel.