# api contract 

## chats_ms [/chats]

Chat is a message from a user(sender) to another user(receiver).

### create a new chat [POST] [/chats]

Create a new chat. That is represent message sent from one user to another. This might be upgraded to use websockets kinda thing.

Sender cannot be equal to receiver of the same chat record.

+ Request (application/json)
	
	{
		"senderId": 1,
		"receiverId": 2,
		"message": "Hello",
		"sentAt": "2020-05-11T18:42:05.418789Z"
	}

+ Response 201 (application/json)

	+ Header
		
		Location : /chats/1

	+ Body
	
		{
			"ID": 1,
			"CreatedAt": "2020-05-11T18:43:05.418789Z",
			"UpdatedAt": "2020-05-11T18:43:05.418789Z",
			"DeletedAt": null,
			"senderId": 1,
			"receiverId": 2,
			"message": "Hello",
			"sentAt": "2020-05-11T18:42:05.418789Z",
			"delieveredAt": "2020-05-11T18:43:10.418789Z",
			"readAt": "2020-05-11T18:43:55.418789Z"
		}

### get all people a user as communicated with [/chats/counterpartsof/{id}]

For a user uniquely identified by {id}, find all users he/she has chat(s) with.

+ Request

+ Response 200 (application/json)

	+ Body	
	
		[
		  {
			"ID": 1,
			"CreatedAt": "2020-05-11T18:43:05.418789Z",
			"UpdatedAt": "2020-05-11T18:43:05.418789Z",
			"DeletedAt": null,
			"Name": "Nelson",
			"Email": "nelsonsaake@gmail",
			"Password": "something",
			"Picture": "/images/profiles/nelsonsaakekofi@gmail.com.jpg",
			"About": "I like banku and tuna. I'm a UMaT student."
		  },
		  {
			"ID": 3,
			"CreatedAt": "2020-05-12T21:12:49.305925Z",
			"UpdatedAt": "2020-05-12T21:12:49.305925Z",
			"DeletedAt": null,
			"Name": "Nelson",
			"Email": "nelsonsaake@gmail.com",
			"Password": "something",
			"Picture": "/images/profiles/nelsonsaakekofi@gmail.com.jpg",
			"About": "I like banku and tuna. I'm a UMaT student."
		  }
		]



### get all sent betweeb two people [GET] [/chats/chats/{u1}/{u2}]

Get all the messages sent between two users.

+ Request

+ Response 200 (application/json)
	
	+ Body
		
		[
			{
				"ID": 1,
				"CreatedAt": "2020-05-11T18:43:05.418789Z",
				"UpdatedAt": "2020-05-11T18:43:05.418789Z",
				"DeletedAt": null,
				"senderId": 1,
				"receiverId": 2,
				"message": "Hello",
				"sentAt": "2020-05-11T18:42:05.418789Z",
				"delieveredAt": "2020-05-11T18:43:10.418789Z",
				"readAt": "2020-05-11T18:43:55.418789Z"
			}
		]
		
		