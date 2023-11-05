# CSRF-Golang-Security

Use this for advanced requirements when you feel that JWT alone is not enough.

This authentication uses tokens(csrf tokens, jwt tokens) as well as RSA pair keys (private and public). 
- The private and public keys are generated using the RSA algorithm. Once you have them, the public key is embedded to the JWT token, and inorder to authenticate, the now then token is compaired with the private key, and if they match, only then will you approve.

- On top of that, CSRF tokens are added. Benefits of these are that they are completely unpredictable and they need to be sent in every http request to the server.

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
