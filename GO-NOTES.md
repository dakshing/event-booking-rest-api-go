# GO NOTES

Short, practical notes and commands to remember while working on this project.

## Notes
- DateTime is handled by time.Time in time package/
- Using a '_' before an import means the package will be used by the program implicitly and do not delete it.

## Quick Commands
- run `go mod tidy` to keep the mod file clean and remove unused dependencies

## Gin Notes
- `gin.GET/gin.POST` binds path to handlers
- `gin.H` is type map[string]any and maps it to JSON directly
- call `context.ShouldBindJSON` to read the POST request body into model.

## DB notes
- use `DB.SetMaxOpenConns(10)` and `DB.SetMaxIdleConns(5)` to limit concurrent connections (default is unlimited and can overwhelm the DB).