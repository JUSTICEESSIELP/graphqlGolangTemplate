# Clone the repository
git clone https://github.com/JUSTICEESSIELP/graphqlGolangTemplate.git

# Navigate to the cloned directory
cd graphqlGolangTemplate

# Run the server
go run server.go




******************************************* OR ****************************************************

# Setup a new directory for gqlgen-todos
mkdir gqlgen-todos
cd gqlgen-todos

# Initialize Go module
go mod init github.com/[username]/gqlgen-todos

# Create tools.go file
printf '// +build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go

# Download gqlgen and introspection packages
go get -d github.com/99designs/gqlgen@v0.17.20
go get github.com/99designs/gqlgen/graphql/introspection

# Initialize gqlgen configuration
go run github.com/99designs/gqlgen init

# Run the server
go run server.go
