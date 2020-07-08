# api contracts

## opportunities_ms [/opportunities]

### get all opportunities available [GET] [/opportunities]

Get all opportunities available

+ Request 

+ Response 200 (application/json)

	+ Body
	
		[
			{
			
				"ID": 1,
				"CreatedAt": "2020-05-12T21:12:49.305925Z",
				"UpdatedAt": "2020-05-12T21:12:49.305925Z",
				"DeletedAt": null,
				"name": "Darvoc",
				"amount": 100000.00,
				"industry": "Technology",
				"description": "Darvoc is a tech company. Our focus is changing the way farming with drones.",
				"userId": 1,
				"picture": "/images/profiles/opportunities/1.jpg",
				"returns": 0.08,
				"duration": 1,
				"location": "Tarkwa"
			}
		]
		
		
### create a new opportunity [POST] [/opportunities]

+ Request (application/json)

	{
		"name" : "Darvoc",
		"asking" : 1003000.00,
		"industry" : "Technology",
		"description" : "Darvoc is a tech company. Our focus is changing the way farming with drones.",
		"userId" : 1
	}

+ Response 201 (application/json)
	
	+ Header
	
		Location: /opportunities/1

	+ Body
		
		{
			"ID": 1,
			"CreatedAt": "2020-05-12T21:12:49.305925Z",
			"UpdatedAt": "2020-05-12T21:12:49.305925Z",
			"DeletedAt": null,
			"name": "Darvoc",
			"amount": 100000.00,
			"industry": "Technology",
			"description": "Darvoc is a tech company. Our focus is changing the way farming with drones.",
			"userId": 1,
			"picture": "/images/profiles/opportunities/1.jpg",
			"returns": 0.08,
			"duration": 1,
			"location": "Tarkwa"
		}
		
### get a list of users that created an opportunity or more opportunities [GET] [/opportunities/creators]

Query for all unique user Id get the related user. This is to say: get all the users in the opportunities repository without repeatition.

This information actually comes from the users_ms.So the opportunities_ms will query the users_ms and return the data.

+ Request 

+ Response 200 (application/json)
			
	+ Body	
	
		[
			{
				"ID": 1,
				"CreatedAt": "2020-05-12T21:12:49.305925Z",
				"UpdatedAt": "2020-05-12T21:12:49.305925Z",
				"DeletedAt": null,
				"email": "nelsonsaakekofi@gmail.com",
				"password": "password",
				"picture": "/images/profiles/nelsonsaakekofi@gmail.com.jpg",
				"firstname": "Nelson",
				"surname": "Saake",
				"dateOfBirth": "",
				"gender": "male",
				"phoneNumber": "0548876758",
				"nationality": "Ghanaian",
				"occupation": "Student",
				"address": "PC Homes, Agric Hill",
				"country": "Ghana",
				"region": "Western",
				"city": "Tarkwa",
				"accountName": "Nelson Kofi Saake",
				"accountNumber": "12345678765432",
				"bankName": "Access",
				"nkSurname": "Saake",
				"nkFirstname": "Rowland",
				"nkRelationship": "Bother",
				"nkEmail": "",
				"nkPhoneNumber": "",
				"nkAddress": ""
			}
		]
		
### get a list of opportunities created by a particullar user [GET] [/opportunities/creators/{id}/history]

Queries a list of all opportunities created by a particullar user with {id}.

+ Request

+ Response 200 (application/json)

	+ Body
		
		[
			{
				"ID": 1,
				"CreatedAt": "2020-05-12T21:12:49.305925Z",
				"UpdatedAt": "2020-05-12T21:12:49.305925Z",
				"DeletedAt": null,
				"name": "Darvoc",
				"amount": 100000.00,
				"industry": "Technology",
				"description": "Darvoc is a tech company. Our focus is changing the way farming with drones.",
				"userId": 1,
				"picture": "/images/profiles/opportunities/1.jpg",
				"returns": 0.08,
				"duration": 1,
				"location": "Tarkwa"
			}
		]


### get a specific opportunity by id [GET] [/opportunities/{id}]

Queries a specific opportunity using the id provided.

+ Request

+ Response 200 (application/json)

	+ Body
	
		{
			"ID": 1,
			"CreatedAt": "2020-05-12T21:12:49.305925Z",
			"UpdatedAt": "2020-05-12T21:12:49.305925Z",
			"DeletedAt": null,
			"name": "Darvoc",
			"amount": 100000.00,
			"industry": "Technology",
			"description": "Darvoc is a tech company. Our focus is changing the way farming with drones.",
			"userId": 1,
			"picture": "/images/profiles/opportunities/1.jpg",
			"returns": 0.08,
			"duration": 1,
			"location": "Tarkwa"
		}

