# api contracts

## investments_ms [/investments]

### get all investments made by a particular user [GET] [/investments/investors/{id}]

Get all investments made a user. Investment is where a user buys into an opportunity. A user can buy into as many opportunities he or she likes. This retrieves all opportunities with a particular user_id. Check investments table for more info.

+ Request 

+ Response 200 (application/json)

	+ Body
	
		[
			{
				"ID": 1,
				"CreatedAt": "2020-05-11T18:43:05.418789Z",
				"UpdatedAt": "2020-05-11T18:43:05.418789Z",
				"DeletedAt": null,
				"userId": 1,
				"opportunityId": 1,
				"amountBought": 10000
			}
		]
		
### get a specific investments info [GET] [/investments/{id}]

+ Request

+ Response 200 (application/json)

	+ Body

		{
			"ID": 1,
			"CreatedAt": "2020-05-11T18:43:05.418789Z",
			"UpdatedAt": "2020-05-11T18:43:05.418789Z",
			"DeletedAt": null,
			"userId": 1,
			"opportunityId": 1,
			"amountBought": 10000
		}
		
### create a new investments [POST] [/investments]

+ Request (application/json)
	
	{
		"userId": 1,
		"opportunityId": 1,
		"amountBought": 10000
	}

+ Response 201 (application/json)

	+ Header
		
		Location : /investments/1

	+ Body
	
		{
			"ID": 1,
			"CreatedAt": "2020-05-11T18:43:05.418789Z",
			"UpdatedAt": "2020-05-11T18:43:05.418789Z",
			"DeletedAt": null,
			"userId": 1,
			"opportunityId": 1,
			"amountBought": 10000
		}
		
### get all investments [GET] [/investments]

+ Request

+ Response 200 (application/json)
	
	+ Body
	
		[
			{
				"ID": 1,
				"CreatedAt": "2020-05-11T18:43:05.418789Z",
				"UpdatedAt": "2020-05-11T18:43:05.418789Z",
				"DeletedAt": null,
				"userId": 1,
				"opportunityId": 1,
				"amountBought": 10000
			}
		]
		
		
	