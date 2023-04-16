These tests perform the following operations:

TestInsertUser inserts a user within a transaction and then rolls back the transaction. This test is useful for verifying that the user insertion code
works correctly without leaving any records in the database.

TestQueryUser first inserts a user and commits the transaction. Then it starts a new transaction, queries the user, and rolls back the transaction. This
test verifies that the user can be queried correctly after being inserted into the database. The rollback at the end does not affect the records in the
database because the user was inserted in a separate transaction.
