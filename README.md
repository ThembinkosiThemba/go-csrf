# CSRF-Golang-Security

Use this for advanced requirements when you feel that JWT is not enough.

Please create "keys" folder at the root level and add a pair of private and public keys using RSA algo. the names of the keys should be - app.rsa (contains private key) and app.rsa.pub(contains public key)

The program won't run without putting the RSA keys as mentioned above :)

Remember to replace mongo connection string with your string :)

### Getting Started 
```bash
git clone <repo-url>
go mod tidy
go run main.go
```

Then go to `localhost:9000/register`

Any questions, hit me up !!! (thembinkosimkhonta01@gmail.com)