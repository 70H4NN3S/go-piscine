## Small test on main.go with an arr size of 90.000.000

Only the amount of respective threads make sense on the go routine, the macbook M4 has 8 cores here, thus, beyond 8 goroutines gives no leverage

Sync SUM: 
11205034142
--------> 23.482708ms
Async SUM: 
11205034142
2 --------> 11.389958ms
Async SUM: 
11205034142
3 --------> 7.797542ms
Async SUM: 
11205034142
4 --------> 5.970875ms
Async SUM: 
11205034142
5 --------> 7.354083ms
Async SUM: 
11205034142
6 --------> 6.109834ms
Async SUM: 
11205034142
7 --------> 5.083667ms
Async SUM: 
11205034142
8 --------> 4.477833ms
Async SUM: 
11205034142
9 --------> 4.657917ms
Async SUM: 
11205034142
10 --------> 4.250709ms
Async SUM: 
11205034142
11 --------> 4.649042ms
Async SUM: 
11205034142
12 --------> 4.028833ms
Async SUM: 
11205034142
13 --------> 4.322041ms
Async SUM: 
11205034142
14 --------> 4.000792ms
Async SUM: 
11205034142
15 --------> 4.652791ms
Async SUM: 
11205034142
16 --------> 4.386959ms
Async SUM: 
11205034142
17 --------> 4.152375ms
Async SUM: 
11205034142
18 --------> 3.918917ms
Async SUM: 
11205034142
19 --------> 4.329375ms
Async SUM: 
11205034142
20 --------> 4.129666ms
Async SUM: 
11205034142
21 --------> 3.950792ms
Async SUM: 
11205034142
22 --------> 3.817166ms
Async SUM: 
11205034142
23 --------> 4.1285ms
Async SUM: 
11205034142
24 --------> 3.690417ms
