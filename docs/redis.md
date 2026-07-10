# Redis Integration

Redis stores the application visitor count.

Current key:

visits

Commands

View value

GET visits

Increment

INCR visits

Reset

DEL visits

Current implementation

Application increments the Redis key on every request to "/".