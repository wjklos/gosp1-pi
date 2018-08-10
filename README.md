# Title

## Intro
The goal is to build upon a basic construct utilizing some of the built in features of Go, and show how a simple, self-contained Golang application can be used to show some pretty advanced concepts without having to utilize any external infrastructure.
Over the course of this post, as well as the planned follow-up posts, we will examine how we can use Google’s Go language to not only do basic web services, but also to explore how we can leverage Go’s concurrent primitives to add some fun, yet powerful abilities that make an application more suited for distributed cloud environments as well as any environment where additional “situational awareness” will lead to a more preditable, flexible, and scaleable application architecture.
## The Web Service Part
Typically, a web service application is deisgned to answer REST calls coming from some sort of front-end (e.g. web, voice, other web service) calculate an answer based on internal knowledge or the result of a database query - then send it back along the same channel from which the request originated.  Preferably along with a 200 as an HTTP response code.  But sometimes, the conditions in which the application lives or the realities of the tasks it needs to perform require some additional flexibilities.  
 Some conditions that may require more than a standard request/response model in our web serices application:
–	Assets on which our web services rely are no longer available and mitigation to those conditions need to be executed;
–	more?
For these tasks, we will want to add a piece to our web services router.  For this piece, we need something we can use to push these tasks to the background. 
The Background Process Part
Google’s Go language is obviously not the first, nor is it the only language that allows for concurrent programming.  However, it is arguably one of the easiest and more robust tools with which to employ it.

Some reasons why we may want to add web service ability to a process that is traditionally, for lack of a better term, a “batch” routine:
–	We want to pull status or operating metrics from the process;
–	We want to change the behavior, rules, or parameters of a running process;
–	more?

–	How We are Going to Achieve It
A heartbeat timer is good for showing the same thing a heartbeat for any living thing is good - it shows you’re alive.  For our application, a heartbeat can be used to keep a stream connection alive, send data to a downstream connection on a timed basis or simply just be our pulse to show that everything is still ok.  If the pulse stops, it may be a good indication that our little application needs to be garbage collected by its higher power (e.g. its container). 

–	Setting Up the Heartbeat Ticker
–	Watching for the Heartbeat
–	Outputting the Proof
## Conclusion

