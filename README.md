
## Note

#### How to add more test case

- use integration_test folder to add more scenario

#### How to check the result

- run unit test
- build docker as below command

#### How to run docker to see unit test result
- docker build -f Dockerfile.multistage -t docker-food-cal-test --progress plain --no-cache --target food-calculator-test .