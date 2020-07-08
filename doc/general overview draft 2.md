application overview
	the application is to advertise opportunities.
	so that potential investors can see these opportunities.
	so that potential investors can buy into some of these opportunities, as in pay, actually give money to into the opportunities.

users
	all the users will be recognised as the system as one
		anyone can browse all opportunities
		anyone can buy into any opportunities
		anyone can provide an opportunities - only admin can provide
	
ui screens
	the system will use 6 pages to interact with the users
		investments page
		browse/market page
		add opportunities page
		chat page
		profile page
		login in page
		signup page
		change password
	
databases
	tables
		users database
		opportunities database
		investments 
			investment table is a transaction table. It holds information on which user buys into what opportunity. 
		chats database
			stores the chats, one-on-conversation made by users on the platform.
			sender of a message cannot be the receiver of that message.
		change password
		post 
			to provide investors with updates on the investments they made.
			The post doesn't actually refer to the investment database, it refers to the opportunities table. Every investment involves a single user and an opportunity. If it was built on investments, the updates will have to be applied individually to all the investments involved. This is because the investment is a linking table, a transaction made by a user to buy into an opportunity. So the post is tied to the opportunities table. It store posts related to opportunities and only reveals those opportunities to investors the concern.
		auths
			holds a record on authorisation codes and the roles they grant users. Such as Admin, User 
			
		
	trascations
		user can buy into an investment
		user can create an opportunity
		user can send a message to another user
		user can request a change of password
		
	relationships 
		every investment is created and maintained by one user, users, opportunities
		every investment can be bought into by more than one user: users, opportunities, investments
		every chat is sent from one user to another which is an admin, users, users, chats
		every change password belongs to a particular user


	tables breakdown		
		users
			id
			createdAt
			updatedAt
			deletedAt
			name 
			email
			password
			picture
			firstname
			surname
			dateOfBirth
			gender
			phoneNumber
			nationality
			occupation
			address
			country
			region
			city
			accountName
			accountNumber 
			bankName
			nkSurname 
			nkFirstname
			nkRelationship
			nkEmail
			nkPhoneNumber
			nkAddress
			
		opportunities
			id
			createdAt
			updatedAt
			deletedAt
			name
			industry
			description
			user_id 
				// who created it
			picture
			amount
			returns
			duration
			location			
			
		investments
			id
			createdAt
			updatedAt
			deletedAt
			user_id
			opportunity_id
			amountBought
			
		chats 
			id
			createdAt
			updatedAt
			deletedAt
			sender_id
			receiver_id
			message
			sentAt
			delieveredAt
			readAt
			
		changepasswords
			id
			createdAt
			updatedAt
			deletedAt
			active
			user_id
			code			
				// Make sure we deactivate the old one before creating a new one for a particular user.
				
		auths
			id
			createdAt
			updatedAt
			deletedAt
			user_id uint64 `unique`
			token string `unique`
			active bool 
				// this table is mainly for authentication, we won't maintain tokens in memory
			role
				// admin or user
				
		posts
			id
			createdAt
			updatedAt
			deletedAt
			pictures
			description
			opportunity_id
				// posts. This is how you update the investors on their investments. A post will expected to come with a picture, and some descriptions. 
				// This is tied to the opportunities instead of investments. Because, investments is a transaction on the opportunities.
			
			
microservices
	database
		tables
			users
			opportunities
			investments
			chats
			changepasswords
			posts
			auths
				// for authentication and authorisation
			
	users
		ui
			investments page
				renders investments of a particular user
			opportunities page
				renders all opportunities created by a particular user
			browse/market page
				renders all investments
			add opportunities page
				allows a user to create an opportunity
			chat page
				allows two users to communicate
			profile page
				dispalys information about a user
			login in page
				allows users to be logged in
			signup page
				allows users to register themselves in the system
			change password
				allows users to change their password
			
	admin
		tables
			users
				renders all users in a list 
				can expand information on a single user
			opportunities
				renders all opportunities
				can expand infromation on a single opportunity
			investments
				renders all investments and 
				can expand opportunities on a single investment
			chats
				renders all chats initiated and
				can expand charts between any two users
			changepasswords
				renders all password change request
				can expand a particular change instance
			
		views
			investors
				provides a list of all investors
					an investor is a user that is in the investment table
				and can provide investment history on any one investor
			opportunity providers
				provides a list of all opportunity providers 
					an opportunity provider is any user in the opportunities table
				and can provide a history on any one provider
	
API contracts	
	



