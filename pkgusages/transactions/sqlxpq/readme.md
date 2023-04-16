Run the test, using the below command.

% go test -v

The package "github.com/lib/pq" should be used together with the package "github.com/jmoiron/sqlx" to process transaction.

When you use transactions, any changes made to the database within that transaction are not committed (made permanent)
until you explicitly call the Commit() method. If you call Rollback() instead of Commit(), any changes made within the transaction will be discarded,
and the database will revert to its state before the transaction began.

The rollback in question will only discard changes made within its own transaction. If you insert a user in a separate transaction and commit those
changes, the rollback in the other transaction will not affect the already committed user records.
