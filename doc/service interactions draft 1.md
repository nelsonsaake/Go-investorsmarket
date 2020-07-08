
1. users_db -> users_ms : read - done
2. users_ms -> users_db : write - done 

3. opportunities_db -> opportunities_ms : read - done
4. opportunities_ms -> opportunities_db : write - done

5. investments_db -> investments_ms : read - done
6. investments_ms -> investments_db : write - done

7. chats_db -> chats_ms : read - done
8. chats_ms -> chats_db : write - done

9. changepasswords_db -> changepasswords_ms : read - done
10. changepasswords_ms -> changepasswords_db : write - done

11. login_ui -> login_ms : send(email, password) * undone
12. login_ms -> login_ui : reply(created or not) * undone
12b. login_ms -> users_ms : send(email, password) - done
12c. users_ms -> login_ms : reply(all users infor or nothing) - done

13. signup_ui -> signup_ms : createUser(email, password, confirmedPassword) * undone
14. signup_ms -> signup_ui : reply(created or not) * undone
14b. signup_ms -> users_ms : exists(email, password, ...) - done
14c. users_ms -> signup_ms : reply(created status, location) - done

15. user_investments_ui -> investments_ms : read(all user investments info) - done
15b. user_single_investments_ui -> investments_ms : read(specific investment info) - done

16. market_ui -> opportunities_ms : read(all investments) - done

16b. buyopportunity_ui -> investments_ms : createInverstment(user_id, opp) - done
16c. investments_ms -> buyopportunity_ui : reply(created or not) - done

17. addopportunity_ui -> opportunities_ms : createOpportunity(opportunities info) - done
17b. opportunities_ms -> addopportunity_ui : reply(status, location, created opportunities from db) - done 

18. chats_ui -> chats_ms : sendNewMessage(message) - done
18a. chats_ms -> chats_ui : get(chats made by a user, that is every person that user initiated conversations with) - done
18b. chats_ms -> chats_ui : get(all messages between two users) - done

19. changepasswords_ui -> changepasswords_ms : create(userid) - done
20. changepasswords_ms -> userMail_ext : send(full chagepassword details) - done

21. allusers_ui -> users_ms : read(all users) - done
22. allopportunities -> opportunities_ms : read(all allopportunities) *same as 16

23. userprofile_ui -> users_ms : read( a user info) - done

24. allinvestments_ui -> investments_ms : read (all inverstments) - done
25. allchangepasswordss_ui -> changepasswords_ms : read (all change passwords) - done

26. allinverstors_ui -> investments_ms : read(all inverstors names) - done
27. allopportunitiesproviders_ui -> opportunities_ms : read(all opportunities providers names) - done

28. providerHistory_ui -> opportunities_ms : read(all opportunities provided by a particular user) - done
29. investmentorHistory -> investments_ms : read(all inverstments made by a particular user) *same as 15