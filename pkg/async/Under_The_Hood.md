# Under the Hood: Go Concurrency Model


> How does the Go scheduler handle the scheduling of goroutines across multiple OS threads?

Go's concurrency model is based on goroutines and channels. Here's how it works under the hood:

1. **Goroutines**: A goroutine is a lightweight thread managed by the Go runtime. They're lighter than traditional operating system threads and the Go runtime can manage hundreds of thousands to millions of goroutines in a single process. When a program starts, it has one goroutine (the main goroutine), and it creates additional goroutines with the go keyword.

2. **Go Scheduler**: The Go runtime contains a scheduler that distributes goroutines over multiple OS threads (which can run on multiple cores). The scheduler's job is to schedule the execution of goroutines. It's a M:N scheduler because it maps M goroutines onto N OS threads. The Go scheduler uses a technique known as work stealing, where idle threads take work from busy ones, leading to better load balancing.

3. **GOMAXPROCS**: By default, Go uses the number of cores available on the machine as the default value for GOMAXPROCS. GOMAXPROCS determines how many OS threads are used by the Go scheduler. You can also manually set GOMAXPROCS to control the number of cores your Go program can utilize.

4. **Threads vs Processes**: In general, a thread is a sequence of instructions within a program that can be executed independently of other code. A process is a running instance of a program and can consist of multiple threads. Threads within the same process share the same memory space, which means they can read from and write to the same variables and data structures. This is why synchronization is necessary when multiple threads access the same data. In contrast, each Go goroutine has its own stack and is scheduled independently by the Go runtime, which simplifies concurrency. The Go runtime handles the scheduling of goroutines across OS threads, so you don't have to worry about managing threads explicitly.

In your code, each go `func() {...}()` statement starts a new goroutine. These goroutines can run concurrently, and the Go runtime will handle scheduling them onto OS threads. The channels (`input`, `doubleOutput`, `squareOutput`) are used to synchronize and pass data between the goroutines.

> How does the Go scheduler handle the scheduling of goroutines across multiple OS threads?


> I've always considered Processes to Threads as a One-to-Many relationship. Is that correct? If so, how does the GOMACPROCS variable fit into this model? Shouldn't it be a One-to-One relationship? As a follow-up (if it's relevant), how many OS threads can a single process have? or how is that number (if not fixed) determined?