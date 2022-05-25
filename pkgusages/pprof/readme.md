go get github.com/pkg/profile


pprof is a tool for visualization and analysis of profiling data. 
It reads a collection of profiling samples in profile.proto format and generates reports visualizing and help analyze the data. 
It can generate both text and graphical reports.

The objective of pprof is to generate a report for a profile. The report is generated from a location hierarchy, which is reconstructed from the profile samples. Each location contains two values:

**flat**: the value of the location itself.
**cum**: the value of the location plus all its descendants.
Samples that include a location multiple times (e.g. for recursive functions) are counted only once per location.

More details please access https://github.com/google/pprof/blob/master/doc/README.md#interpreting-the-callgraph


**% pprof -alloc_objects http://:1234/debug/pprof/allocs**
Fetching profile over HTTP from http://:1234/debug/pprof/allocs
Saved profile in /Users/daweijiang/pprof/pprof.alloc_objects.alloc_space.inuse_objects.inuse_space.002.pb.gz
Type: alloc_objects
Time: May 10, 2022 at 11:14am (NZST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) nodefraction=0
(pprof) **top**
Showing nodes accounting for 39532, 100% of 39532 total
Showing top 10 nodes out of 43
flat  flat%   sum%        cum   cum%
16384 41.44% 41.44%      16384 41.44%  net/textproto.(*Reader).ReadLine (inline)
13108 33.16% 74.60%      13108 33.16%  net.(*conn).Read
2521  6.38% 80.98%       2521  6.38%  runtime.malg
2380  6.02% 87.00%       2380  6.02%  bytes.makeSlice
2341  5.92% 92.92%      19238 48.66%  net/http.(*conn).readRequest
2277  5.76% 98.68%       2277  5.76%  runtime.allocm
513  1.30%   100%        513  1.30%  bufio.NewWriterSize (inline)
4  0.01%   100%          4  0.01%  compress/flate.newDeflateFast
2 0.0051%   100%          6 0.015%  compress/flate.(*compressor).init
2 0.0051%   100%          8  0.02%  compress/flate.NewWriter
(pprof)

In the above output, you can see two columns, flat and cum. 
Flat means only by this function, while cum (cumulative) means by this function and functions called down the stack.

For better understanding, let’s assume the following function, the flat time of function A is 4s and the cum is 11s.

func A() {
    B()             // takes 1s
    DO STH DIRECTLY // takes 4s
    C()             // takes 6s
}

**How to view SVG directly using pprof?**

The first way is not flexible.
Step1: brew install Graphviz
Step2: pprof -web http://:1234/debug/pprof/goroutine

% **pprof -web http://:1234/debug/pprof/goroutine**
Fetching profile over HTTP from http://:1234/debug/pprof/goroutine

The above command should automatically open your web browser at the graph page
Please access this link "https://git.io/JfYMW" to know how to view the graph.

The second way is highly recommended.
Step1: % curl http://localhost:1234/debug/pprof/allocs > allocs.out
Step2: go tool pprof -http=:12345 allocs.out  #12345 means any port you can choose  


CPU Profiler
The Go CPU profiler uses a SIGPROF signal to record code execution statistics. Once the signal got registered, it will deliver every specified time interval. This timer unlike typical timers will increment when the CPU is executing the process. The signal will interrupt code execution and make it possible to see which code was interrupted.

When the pprof.StartCPUProfile function is called, the SIGPROF signal handler will be registered to call every 10 milliseconds interval by default (100 Hz). On Unix, it uses a setitimer(2) syscall to set the signal timer.

On every invocation signal, the handler will trace back the execution by unwinding it from the current PC value. This will generate a stack trace and increment its hit count. The whole process will result in a list of stack traces grouped by the number of times each stack trace is seen.

After the profiler is stopped, the stack traces will be converted to pprof-supported stack frames with source file names, line numbers, and function names.

Memory Profiler
The memory profiler samples heap allocations. It will show function calls allocations. Recording all allocation and unwinding the stack trace would be expensive, therefore a sampling technique is used.

The sampling process relies on a pseudo-random number generator based on exponential distribution to sample only a fraction of allocations. The generated numbers define the distance between samples in terms of allocated memory size. This means that only allocations that cross the next random sampling point will be sampled.

The sampling rate defines the mean of the exponential distribution. The default value is 512 KB which is specified by the runtime.MemProfileRate. That means to tune the size of the fraction to be sampled, you have to change runtime.MemProfileRate value. Setting the value to 1, will include every allocated block in the profile. Obviously, this will slow down your application.

Stack traces with corresponding sampled allocations information is available at any time since the memory allocation profiler is always active. Setting the runtime.MemProfileRate to 0, will turn off the memory sampling entirely.

Once the stack traces got ready, those will be converted to pprof-supported stack frames.


**How to use hey?**

Step1: brew install hey

Step2: go run .

Step3: using hey in command line
% hey http://localhost:1234/debug/pprof/