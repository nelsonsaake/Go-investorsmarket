# api contracts 

## users_ms [/users]

### Get user info [POST] [/users/a/]

Get a user info based on, given the email and password in request

+ Request (application/json)

	{
		"email": "nelsonsaake@gmail.com",
		"password": "something"
	}

+ Response 200 (application/json)

    + Body

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


### Greate a new user [POST] [/users]

Create a new user

+ Request (application/json)

	{
		"email": "nelsonsaake@gmail",
		"password": "something"
	}
	
+ Response 201 (application/json)


    + Headers

        Location: /users/1
		   
	+ Body
		
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
		  
### List all users [GET] [/users]
		
Get the list of all registered users
		
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

### Get user profile [GET] [/users/{id}]

Get the all information about a user

+ Request

+ Response 200 (application/json)

	+ Body
	
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
		
### Update profile [PUT] [/users/{id}]
		
Update a user profile. Profile is all the information on a user.

+ Request (application/json)

	{
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

+ Response 200 (application/json)
		
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
		
		
		
		
		
		
		
		
		