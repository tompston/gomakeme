# Notes

- Link to SQLC Documentation -> https://docs.sqlc.dev/en/stable/
- If running on windows and need postgres support, install wsl and dowload sqlc binary to generate code
- The generated tables come with 3 predefined columns that would be used by most of the tables
- Each generated moudule sql file holds 5 temporary CRUD queries that need to be  updated before running sqlc. 
- Note that the "-- name" of the query is the same as the name of the controller for the module. 
- All of the "column_name" values are placeholders.  
- The name of the generated query is the same as the name of the controller function. 
