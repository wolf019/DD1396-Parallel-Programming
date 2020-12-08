## Ovn2

### Task1

- Bug 1: It is because we cant know if "Hello world!" have been sent before we want to print what's on ch. Println will try to print something that isn't there."Hello world!" is on the channel at some point, but Println is not waiting for that. Solution here is to make the channel buffered. By making the channel hold one element, Println could only get one value, which it will wait for as it's the only thing to find:

Correction: Channel is blocked until both sender and receiver are ready. When we try to send "Hello world!" we can't because the channel's capacity is zero or absent as no receiver is established and "Hello world!" can't be sent. When Println(<-ch) executes we have a receiver but "Hello world!" has not been sent.
[task1_bug1.go](https://gits-15.sys.kth.se/dd1396ht20/taxberg-ovn2/blob/master/task1_bug1.go)

- Bug 2: Its because we don't know if the goroutine Print will handle the last i in the channel before main goroutine terminates. To make sure that the main goroutine doesn't terminate until Print is done, we include "sync" to wait until it is done:
[task1_bug2.go](https://gits-15.sys.kth.se/dd1396ht20/taxberg-ovn2/blob/master/task1_bug2.go)

### Task2

- Question  1: If we close the channel, before wait, we will close it before the producers and consumers have exchanges strings with each other.

- Question  2: If we move close, we will have four producers that all could close the channel separately. Once one is done sending strings, it will close the channel, and no more producers will be able to send strings. We will get panic.

- Question  3: What happens if we remove close is nothing. When the main goroutine terminates, the producers are done, and we are done.

- Question  4: Changing the number of consumers to four will make the producers get done fasters because the strings are consumed faster. Adding more consumers than four will not make it faster as the strings produced are limited. 

- Question  5: We can't be sure that all things are printed. We never check if consumers are done. 

- Modified code: [many2many.go](https://gits-15.sys.kth.se/dd1396ht20/taxberg-ovn2/blob/master/many2many.go)

### Task 3

- Code: [oracle.go](https://gits-15.sys.kth.se/dd1396ht20/taxberg-ovn2/blob/master/oracle.go)
