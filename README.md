# Flake-auth
Flake-auth this authentication service

# Config 

```
addr = "Address"
log_level = "debug"

[store]

db_url = "URL DATABASE"

[auth]

token = "Token Key"

```
# Start params

```config-path = "Path conf.toml"```

# EndPoints 

```
/reg

{
    name: string,
    email: string,
    password: string,
    fullname: string,
}

/login

{
    name: string
    password: string,
}

/token 

{
    token: string,
}

```