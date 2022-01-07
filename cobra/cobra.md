# Cobra

Cobra help create an application and add any commands you want.

cobra install

```bash
go install github.com/spf13/cobra/cobra
cobra
```

Initializing a new Go module

```bash
mkdir myapp
cd myapp
go mod init github.com/spf13/myapp
cat myapp/go.mod
```

Initializing an Cobra CLI application

```bash
cd myapp
cobra init --viper
go run main.go
```

`--viper` flag automatically set up viper

Viper is a companion to Cobra to provide easy handling of environment variables and config files and connect them to the appropriate flags

Add commands

```bash
cobra add serve
cobra add config
cobra add create -p 'configCmd'
```

```bash
go run main.go
go run main.go serve
go run main.go config
go run main.go config create
```

You now have a basic Cobra-based application up and running. Next step is to edit the files in cmd and customize them.

rootCmd in `app/cmd/root.go`

additionally define flags and handle configuration in init() function