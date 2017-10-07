# Writing tweets to a JSONL file

## Credentials
In order to utilize the twitter API we need to create an app and get the proper
credentials.
Once we get these, we will store them in an `env` file.
```
export CONSUMER_KEY=xxxxxxxxxxxx                                   
export CONSUMER_SECRET=xxxxxxxxxxxxxx       
export ACCESS_TOKEN=xxxxxxxxxxxxxx          
export ACCESS_TOKEN_SECRET=xxxxxxxxxxxxxx
```

Notice that in this sample `env` file we added the word `export` before each
environment variable, this is so we can test out our program for local dev by
doing
```
source <environmnet.env>
```

## Getting tweets
The `twitterConfig` struct willbe used to get the env variables from our
session. This struct takes advantage of the [`env` package](https://github.com/caarlos0/env)

