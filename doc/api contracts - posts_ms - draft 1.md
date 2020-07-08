# api contracts

## posts_ms [/posts]

Posts are used to show the progress of opportunities. 

### create a post [POST] [/posts]

+ Request (application/json)

	{
		"picture": "/images/posts/opportunity1.2.jpg",
		"description": "Request stage!",
		"opportunityId": 1
	}

+ Response 201 (application/json)

	+ Header
		
		Location: /posts/1
		
	+ Body
	
		{
			"ID": 1,
			"CreatedAt": "2020-05-11T18:43:05.418789Z",
			"UpdatedAt": "2020-05-11T18:43:05.418789Z",
			"DeletedAt": null,
			"picture": "/images/posts/opportunity1.2.jpg",
			"description": "Request stage!",
			"opportunityId": 1
		}

### get all posts related to a particular opportunity [GET] [/posts/opportunity/{id}] 

+ Request 

+ Response 200 (application/json)

	+ Body
	
		[
			{
				"ID": 1,
				"CreatedAt": "2020-05-11T18:43:05.418789Z",
				"UpdatedAt": "2020-05-11T18:43:05.418789Z",
				"DeletedAt": null,
				"picture": "/images/posts/opportunity1.2.jpg",
				"description": "Request stage!",
				"opportunityId": 1
			}
		]
		
		