# About
I needed a fast way to add gpg keys to repos on my machine for different email addresses, this small tool solves that for me

## Steps
1. Run ```gpg-gen``` copy, paste and run cmd on about section to generate a key, pick default opts
2. Run ```gpg-gen keys```

3. To pick a key for step:
Example:
```
sec   4096R/3AA5C34371567BD2 2016-03-10 [expires: 2017-03-10]
uid                          Hubot <hubot@example.com>
ssb   4096R/4BB6D45482678BE3 2016-03-10
```
key id would be: ```3AA5C34371567BD2```
Email would be: ```hubot@example.com```

Run ```gpg-gen add --key <key-id> --email <email>```

4. Run ```gpg-gen github``` add the generated token to your github profile GPG Keys, make sure the same email address used with the key is also linked to your github account

# Install
1. Copy repo to local
2. Run ```go build && go install``` now ```gpg-gen``` is part of your ```./bin/go``` path
