# api contract

## auths_ms [/auths]

### create a new authorisation package [POST] [/auths]

Create a new authorisation package. Normally a header with a location will be returned. But this authorisation and no expected to request an authrisation package so that would not be provided.

+ Request (application/json)

{
	"email": "nelsonsaakekofi@gmail.com",
	"password": "something"
}

+ Response 201 (application/json)
	
	+ Body 
		
		{
			"userId": 1,
			"token": "2sfT9t5U096Cim9Nv5ZogCZWQ9MfwaCP1w6mCVe6al5ohRB24dn6fpoNlQMVDnUrxtV50oOsRhe8tubmfZFANu9u",
			"role": "admin"
		}

### get authorisation package given email and password [POST] [/auths/ep]

get the userId, token, and role given the user email and password

+ Request (application/json)

	{
		"email": "nelsonsaakekofi@gmail.com",
		"password": "something"
	}

+ Response 200 (application/json)

	+ Body 
	
		{
			"userId": 1,
			"token": "2sfT9t5U096Cim9Nv5ZogCZWQ9MfwaCP1w6mCVe6al5ohRB24dn6fpoNlQMVDnUrxtV50oOsRhe8tubmfZFANu9u",
			"role": "admin"
		}

### get authorisation package given token [POST] [/auths/t1]

get the user id, token, and role given the user token

+ Request (application/json)

	{
		"token": "2sfT9t5U096Cim9Nv5ZogCZWQ9MfwaCP1w6mCVe6al5ohRB24dn6fpoNlQMVDnUrxtV50oOsRhe8tubmfZFANu9u",
	}

+ Response 200 (application/json)

	+ Body 
	
		{
			"userId": 1,
			"token": "2sfT9t5U096Cim9Nv5ZogCZWQ9MfwaCP1w6mCVe6al5ohRB24dn6fpoNlQMVDnUrxtV50oOsRhe8tubmfZFANu9u",
			"role": "admin"
		}

### send  [POST] [/auths/ut]

validate user id and token. We get the authorisation package. If the user id does not match the token. We get unauthorised.

+ Request (application/json)

	{
		"userId": 1,
		"token": "2sfT9t5U096Cim9Nv5ZogCZWQ9MfwaCP1w6mCVe6al5ohRB24dn6fpoNlQMVDnUrxtV50oOsRhe8tubmfZFANu9u",
	}

+ Response 200 (application/json)

	+ Body 
	
		{
			"userId": 1,
			"token": "2sfT9t5U096Cim9Nv5ZogCZWQ9MfwaCP1w6mCVe6al5ohRB24dn6fpoNlQMVDnUrxtV50oOsRhe8tubmfZFANu9u",
			"role": "admin"
		}
		
		