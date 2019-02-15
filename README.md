# This project is for personal learning and documentation with resources. 

## Reading Materials

### Basic Learning
*  https://www.golang-book.com/books/intro

### Concurrency in action 
*  http://divan.github.io/posts/go_concurrency_visualize/
*  https://codeburst.io/why-goroutines-are-not-lightweight-threads-7c460c1f155f
*  https://golangbot.com/channels/
*  https://golangbot.com/goroutines/

### Performance Caveats 
*  http://container-solutions.com/surprise-golang-thread-scheduling/
* Defer, Panic and Recover: https://blog.golang.org/defer-panic-and-recover
### Project Structure

![alt text](https://github.com/nanofaroque/golang_basics/blob/master/project_structure.png)

* https://golang.org/doc/code.html

## Golang microservice with Kafka and MongoDB 
* https://www.melvinvivas.com/developing-microservices-using-kafka-and-mongodb/

### Quick Note
#### Iterating over map:
```for k, v := range m { 
       fmt.Printf("key[%s] value[%s]\n", k, v)
   }
   or
   
   for k := range m {
       fmt.Printf("key[%s] value[%s]\n", k, m[k])
   }```
   
   