# Övning 1 pallinda20

This work should be completed before the exercise on **Friday 30th October**.

Samtliga uppgifter på kursen ska lämnas in på
[organisationen dd1396ht20 på KTH GitHub](https://gits-15.sys.kth.se/dd1396ht20).

*Gå till [https://gits-15.sys.kth.se/](https://gits-15.sys.kth.se/) och logga in för att aktivera ditt konto.*

Repos för de tre övningarna heter `username-ovnx`.
Användaren `nisse` hittar alltså sitt repo för övning 1 på
[https://gits-15.sys.kth.se/dd1396ht20/nisse-ovn1](https://gits-15.sys.kth.se/dd1396ht20/nisse-ovn1).

Gör (minst) en fil per uppgift. Utgå från mallarna i
[/pallinda20/ovn0/](https://github.com/yourbasic/pallinda20/tree/master/ovn0).
Lösningar skall vara inlämnade **innan övningen börjar**.

- Vid övningen ska du vara beredd att muntligt presentera och diskutera dina lösningar och din programkod.
- Uppgifter märkta med HANDIN ska ockå lämnas in skriftligt innan övningens start.


### Homework
Study the following course literature:

- Read [Why Go? – Key advantages you may have overlooked](https://yourbasic.org/golang/advantages-over-java-python/)
- Read the following from the [Step-by-step guide to concurrency](http://yourbasic.org/golang/concurrent-programming/)
  - [Goroutines](http://yourbasic.org/golang/goroutines-explained/)
  - [Channels](https://yourbasic.org/golang/channels-explained/)

### Task 1 - Go Environment

The first task is to determine that you have a functioning Go environment on the computer that you are working from.

- On a KTH computer - Go should be installed and ready to use
- On your own computer - Goto the [Go getting started page](https://golang.org/doc/install) and follow the installation instructions for your preferred operating system.

### Task 2 - A Tour of Go (HANDIN)

In this task we shall follow the online exercises hosted on [A Tour of Go](http://tour.golang.org/welcome/1).  Start at the beginning and read through the tutorial. You are expected to submit solutions for the following exercises:

- [Loops and Functions](http://tour.golang.org/flowcontrol/8)
- [Slices](http://tour.golang.org/moretypes/18)
- [Maps](http://tour.golang.org/moretypes/23)
- [Fibbonacci Closure](http://tour.golang.org/moretypes/26)

Remember to format your code.  Go has a unapologetic tool built-in that will reformat your code according to a set of style rules made by the designers of the language. To run the format utility, use the following command for all submissions:

    $ go fmt

### Task 3 - Alarm Clock (HANDIN)

In this task you will explore time functions using Go.  Write a function `Remind(text string, delay time.Duration)` that will print the following output:

    Klockan är XX.XX: + <text>

The output will repeatedly print the output after the given delay, and `XX.XX` should be replaced with the current time, and `<text>` should be replaced by the contents of `text`.

Now, write a complete program that runs indefinitely and prints the following reminders:

* every 3rd hour: `Klockan är XX.XX: Dags att äta`
* every 8th hour: `Klockan är XX.XX: Dags att arbeta`
* every 24th hour: `Klockan är XX.XX: Dags att sova`

To prevent the main program from exiting early, the following statement can be used:

```Go
select { }
```

In order to access time related functions, you should investigate the [time package](https://golang.org/pkg/time/), and discover how to get the current time in Go and also how you can format it neatly for human users to understand.  Remember to test and format your code.

### Task 4 - Two Part Sum (HANDIN)

In this task you will complete the following partial program.  It adds all of the numbers in an array by splitting the array in half, then having two Go routines take care of each half.  Partial results are then sent over a channel.  Remember to test and format your code.

```Go
package main

// Add adds the numbers in a and sends the result on res.
func Add(a []int, res chan <- int) {
    // TODO
}

func main() {
    a := []int{1, 2, 3, 4, 5, 6, 7}
    n := len(a)
    ch := make(chan int)
    go Add(a[:n/2], ch)
    go Add(a[n/2:], ch)

    // TODO: Get the subtotals from the channel and print their sum.
}
```
