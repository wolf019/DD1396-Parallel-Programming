## Ovn3

### Task 1 - Matching Behaviour 

 - If we remove `go-command` from `Seek` in `main` we will just run one `Seek`- function after the other, resulting in predictable matchmaking as we are iterating through the `people` slice in order. 
 - If we change `new` to `var` we will get a panik as `Wait` will wait forever to be "unlocked". This is because when we have our `WaitGroup` as a `var` we don´t get `wg` as a pointer. This is why we need to change `Seek` func so it receives the actual type and not a type pointer. What we will call `Seek` with is a copy of `wg`, which will make our `Wait` wait forever as stated.
 - We will get a deadlock as the last person will wait forever for a match / another person who seeks someone. 
 - Nothing. But if we add or remove a name in the `people` slice, we will get a deadlock as we will get stuck in the `select` statement in `main`. 

### Task 2 - Fractal Images

- I have 8 CPUs
- Given `Julia.go` version terminates in 16.4 seconds, my first parallel version [julia_v1.go](https://gits-15.sys.kth.se/dd1396ht20/taxberg-ovn3/blob/master/julia_v1.go) terminates in 3 seconds and my second [julia_v2.go](https://gits-15.sys.kth.se/dd1396ht20/taxberg-ovn3/blob/master/julia_v2.go) terminates in 2.5 seconds. Result: Parallel second version is 13.9 seconds faster than the given non-parallel `Julia.go`.


### Task 3 - Weather station

[client.go](https://gits-15.sys.kth.se/dd1396ht20/taxberg-ovn3/blob/master/client.go)
