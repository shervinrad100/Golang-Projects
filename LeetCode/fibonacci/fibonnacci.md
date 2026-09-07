Write a function to calculate the fibonnaci of n where the function is calculated as such:

```
    fibonnacci(n) = fibonnacci(n-1) + fibonnacci(n-2);
    fibonnacci(1) = fibonnacci(2) = 1
```

This is to practice dynamic programming. 


### Learnings
without memoization you are duplicating the same calculation. When memory is available I think go's compiler does some caching on the recursions but when you try to calculate larger fibonnacci numbers you can see that it takes a long time. But when you record the values in a cache you are immediately fetching them rather than calculating. 

Another thing to keep an eye out for is that the iterative function runs a lot faster. Not sure why but one assumption is that there's less overhead because it's not creating more function calls and stacks which causes the data to be written to a heap so it can be accessd by other function calls. 