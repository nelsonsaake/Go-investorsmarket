# api contract

## changepasswords_ms [/changepasswords]

### request a change of password record to be created [POST] [/changepasswords]

Make sure we deactivate the old one before creating a new one for a particular user.

+ Request (application/json)
	
	{
		"userId": 1
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
			"active": true,
			"userId": 1,
			"code": "x3c4t5"
		}

### request a record of: change of password [POST] [/changepasswords/g/1]

The 1 in the route is a dummy. We don't actually use it. The informations we need is the userId and the code. This is unique because for any use there can be only one active password change request record. So we pull that record and make sure that the record code match. The actually changing of passwords will not be done here(not with this microservice), it would be done from users. users ms is the only legal modifier of users db.

+ Request (application/json)
	
	{
		"userId": 1,
		"code": "2sfT9t5U096Cim9Nv5ZogCZWQ9MfwaCP1w6mCVe6al5ohRB24dn6fpoNlQMVDnUrxtV50oOsRhe8tubmfZFANu9u"
	}

+ Response 200 (application/json)

	+ Body
	
		{
			"ID": 1,
			"CreatedAt": "2020-05-11T18:43:05.418789Z",
			"UpdatedAt": "2020-05-11T18:43:05.418789Z",
			"DeletedAt": null,
			"active": true,
			"userId": 1,
			"code": "2sfT9t5U096Cim9Nv5ZogCZWQ9MfwaCP1w6mCVe6al5ohRB24dn6fpoNlQMVDnUrxtV50oOsRhe8tubmfZFANu9u"
		}

### request all change of password records [GET] [/changepasswords]

This is intended for some troubleshoting on the admin side of things

+ Request

+ Response 200 (application/json)

	+ Body
		
		[
			{
				"ID": 1,
				"CreatedAt": "2020-05-11T18:43:05.418789Z",
				"UpdatedAt": "2020-05-11T18:43:05.418789Z",
				"DeletedAt": null,
				"active": true,
				"userId": 1,
				"code": "2sfT9t5U096Cim9Nv5ZogCZWQ9MfwaCP1w6mCVe6al5ohRB24dn6fpoNlQMVDnUrxtV50oOsRhe8tubmfZFANu9u"
			}
		]
	

