# Todo App

currently working on my spare time. some parts of the code is really janky :P. i hope i will fix these parts later.
open to any reviews (very grateful also)

Current Milestone:

- [ ] Fix the authontication (currently every request made to todo service goes to auth service to auth user. But i
  should use public keys to auth validate the token)
- [ ] do better error handling (custom error classes...)
- [ ] add friend service
- [ ] add todo service
- [ ] add async messagingq service between services (rabbitmq)
- [ ] dockerize every service
- [ ] make a frontend for the app
- [ ] deploy using kubernetes

Ideas:

- [ ] try grpc
- [ ] maybe get rid of gorm ?
- [ ] try deploying without k8s (maybe using terraform)
- [ ] implement a notification service using websockets
- [ ] healtcheck and monitoring services 

