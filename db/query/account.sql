
/* name: CreateAccounts :exec */
INSERT INTO accounts (
    owner,
    balance,
    currency,
    created_at
) VALUES (
    ?,
    ?,
    ?,
    ?
);