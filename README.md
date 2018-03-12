# README #


### What is this repository for? ###

--------        ----------  ----                ----          -------           ------                  ------------------
        ---------------       ------    -------          -------       -------------        ------                  ------------------
    -----           -----     ------    --------        --------     -----------------      ------                  ------------------
    -----            ----     ------    ----------    ----------    --------    -------     ------                  ------
    -----                     ------    ----------- ------------    ------      -------     ------                  ------
     -----                    ------    ------------------------    --------   -------      ------                  ------
      -----                   ------    ------            ------    ----------------        ------                  ------
        -----                 ------    ------            ------    -------------           ------                  ------------------
          -----               ------    ------            ------    ------                  ------                  ------------------
            -----             ------    ------            ------    ------                  ------                  ------------------
              -----           ------    ------            ------    ------                  ------                  ------
                -----         ------    ------            ------    ------                  ------                  ------
                   -----      ------    ------            ------    ------                  ------                  ------
                    -----     ------    ------            ------    ------                  ------                  ------
                    -----     ------    ------            ------    ------                  ------                  ------
    ------         ------     ------    ------            ------    ------                  --------------------    ------------------
     -----------------        ------    ------            ------    ------                  --------------------    ------------------
         ---------          ----------  ------            ------    ------                  --------------------    ------------------

* Quick summary: Simple is a programming language i am working on for educational purpose 
* Version : 1.0.0

### How to run your code ###
You can write your code directly in the simple console and have evaluated on the fly 

Or you can run your .sim file by typing
```text
simply run fileName.sim
```

### Examples ###

* Reduce function 
```javascript
var reduce = fn(arr, initial, f){
    var iter = fn(arr, result){
        if (len(arr) == 0) {
            result
        } else {
            iter(rest(arr), f(result, first(arr)));
        }
    };
    iter(arr, initial);
};

var sum = fn(arr){
    reduce(arr, 0, fn(initial, el) { initial + el });
};

println(sum([1, 2, 3, 4, 5]));
```

* Simple Interest
```javascript
printLn("Please enter the principal"); 

var p = parseInt(readLn());

printLn("Please enter the interest rate");
var r = parseInt(readLn());

printLn("Please enter the time");
var t = parseInt(readLn());

var simepleInterest = fn(p, t, r) {
    return (p*t*r)/100
}

var i = simepleInterest(p, t, r);

print("The interest is ", i);

printLn("Thanks");

```

* Factorial
```javascript
var factorial = fn(n){
    if(n == 1){
        n
    }else{
        n * factorial(n-1)
    }
}

printLn("The factorial of", 20, "is", factorial(20));
```

### Contribution guidelines ###

* Writing tests
* Code review

### Who do I talk to? ###

* Repo owner or admin
* Other community or team contact