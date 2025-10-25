# GO NOTES

Short, practical notes and commands to remember while working on this project.

## Notes
- DateTime is handled by time.Time in time package/
- Using a '_' before an import means the package will be used by the program implicitly and do not delete it.
- Struct tags are metadata providers and there must not be any space between key and value.
- nil can be assigned to a struct, but can be assigned to a pointer.
- `strcnv` is util to convert string to other types
- Use `bcrypt` library for encryption. Passwords are hashed here using bcrypt.

## Quick Commands
- run `go mod tidy` to keep the mod file clean and remove unused dependencies

## Gin Notes
- `gin.GET/gin.POST` binds path to handlers
- `gin.H` is type map[string]any and maps it to JSON directly
- call `context.ShouldBindJSON` to read the POST request body into model.

## DB notes
- use `DB.SetMaxOpenConns(10)` and `DB.SetMaxIdleConns(5)` to limit concurrent connections (default is unlimited and can overwhelm the DB).
- Query with context helps you define timeout and cancel running queries.