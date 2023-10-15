# Important Note
- Not finished (I deeply apologize for this but I don't have enough time ATM to complete the mvp)

# Clean Code Architecture
![Architecture](https://imgur.com/D8cr9ze.png)

# Tech Stacks
- Server   -> Go:Fiber
- Database -> SQLite
- ORM      -> GORM
> why use ORM? to save time.

# Installation

- Clone the Repo
```bash
git clone https://github.com/resqiar/synapsis_test.git
```

- Run the Server
```bash
go run main.go
```

# API Documentations

## Register User

```bash
POST /auth/register
```

Body
```go
{
	username string, // required, min=3, max=100
	password string, // required, min=8, max=100"`
}
```

Possible Return
```
200, 500 Status
```
```json
{
    "error": string,
}
```

## Create Product

```bash
POST /product/create
```

Body
```go
{
	Title       string // required, min=3, max=100"`
	Description string // max=100
	ImageURL    string // url
}
```

Possible Return
```
200, 500 Status
```
```json
{
    "error": string,
}
```

## Create Category

```bash
POST /category/create
```

Body
```go
{
	Title       string // required, min=3, max=100"`
	Description string // max=100
}
```

Possible Return
```
200, 500 Status
```
```json
{
    "error": string,
}
```
