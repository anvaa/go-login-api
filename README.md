go 1.21.3
<br>
A login api written in Go, GIN and GORM using Sqlite3 as data source. The middelware uses JWT and cookies for authentication.
<br>
<br><br>
Usage:<br>
- clone the reposetory<br>
- cd into go-login-api<br>
- Execute: go run .<br>
- data/data.db and .env will be created on first run<br>
- First user to sign up will be superadmin<br><br>

Before relese:<br>
- Remeber to set GIN_MODE="release" in .env<br>
- REM out test secret in loadEnv.go<br><br>

- Hope you enjoy it. Good luck<br>