# CompetitiveProgrammingJudge

Learning how to use golang as a backend server + production level backend basics (caching, ratelimiting, batch processing, reverse proxy, databases, queues, and containerization)

## todo

I'm using LLMs btw, trying to understand how everything works. I myself don't know (or have forgotten) about a lot of the imp things. So, i'll go and revise them

- [x] Revise [GoLang](GoLang.md)
- [ ] Revise Redis
- [ ] Revise Docker
- [ ] Learn how to use Nginx (Reverse Proxies)
- [ ] Learn how to use RabbitMQ
- [ ] Graceful shutdown & signal handling
- [ ] Rate limiting (token bucket, leaky bucket)
- [ ] Caching strategies
- [ ] External API failures
- [ ] GRPC

## Project Goal

Building a **competitive programming judge system**:
- Accept code submissions
- Run them securely in isolated environments
- Evaluate against test cases
- Return verdicts (AC / WA / TLE / RE)


Things to ensure
- nginx API gateway (reverse proxy, and acts as a ratelimiter too)
- Go backend saves the metadata of the submission in local postgres, and pushes the job to a queue (perhaps rabbitmq)
- The judge instance keeps pulling jobs from the queue, sets up proper constraints, and security stuff. (sandbox)
- Cache for tcs maybe?


## File Structure
- cmd : Stores the entry point for the API
- internal : Stores the judging code, and database
- api : some other api stuff