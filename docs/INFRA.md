# Database Schema

## Events
- id (UUID) (PRIMARY)
- name (TEXT) (NOT NULL)
- created_at (TIMESTAMP) (NOT NULL)
- created_by (UUID) (NOT NULL)
- status (TEXT) (NOT NULL) (open|archived)

## Timers
- user_uuid (UUID) (NOT NULL)  UNIQUE(user_uuid,event_id) # uniquely identifies a user's timer for an event
- event_id (UUID) (NOT NULL)   FOREIGN KEY(event_id) REFERENCES events(id) # references to events(id)
- start_time (TIMESTAMP)
- end_time (TIMESTAMP)


## Audit Logs
- id (UUID) (PRIMARY KEY)
- event_id (UUID) (NOT NULL)
- user_uuid (UUID) (NOT NULL)
- action (TEXT) (NOT NULL)
- created_at (TIMESTAMP) (NOT NULL)

# API

## Events
- Create __[Post]__ (Name, Who created for Audit)
- Open __[Post]__ (Id, Who's opening for Audit)
- Archive __[Post]__ (Id, Who's archiving for Audit)
- Get __[Get]__ (All Events, if not eboard or RTP just give open ones)


## Timer
- Start __[Post]__ (UserUUID, EventId)
- End __[Post]__ (UserUUID, EventId)


## Leaderboard.
- Top __[Get]__ (queryParam eventId)
- Group __[Get]__ (queryParam eventId, timerVal (for pagination))