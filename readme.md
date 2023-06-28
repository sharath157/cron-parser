# Cron Expression Parser

This Project Parse the Basic cron Expression & Expands the fields to understand and show times at which would run

## Cron Expression Format

```[minutes] [hours] [Days of the months] [Months]  [day of week]```

 Each field can take any of the below formats

| Format | Description                     |
|--------|---------------------------------|
| '*'    | Accepts All values              |
| 1-2    | describes startRange - endRange |
| */2    | frequency of every 2nd unit     |
| 1,2    | Fixed values                    |
| 1      | Fixed value                     |



### Example

```
Intput Arguments */15 0 1,15 * 1-5 [cmd]

Output:

minute          0 15 30 45
hour            0
day of month    1 15
month           1 2 3 4 5 6 7 8 9 10 11 12
day of week     1 2 3 4 5
command         /usr/bin/find
```

# Environmental Setup (pre-requesites)

- Install Go ``brew install go`` or Install via link & follow steps - https://go.dev/doc/install

# How to Run the program

- Move to root dir of the git repo

- Download dependencies using ``go mod tidy``

 Run the code as 
 ``go run main.go <arguments>``

# How to Run Test cases
- Move to root dir of the git repo, Run 
``go test ./...``