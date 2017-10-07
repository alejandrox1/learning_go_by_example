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
We will be making this a sort of command line tool which will take the number
of tweets to gather as a command line argument.
In `main` you can see that we start by getting the value for the number of
tweets we want to save into the var `tweetCount` and if no value is passed then
we default to 20.

After this we call the func `writeTweets` which will do the following:

1. Get the credentials to authenticate against the twitter API from the
environment. Here the `twitterConfig` struct willbe used to get the env variables from our 
session. This struct takes advantage of the [`env` package](https://github.com/caarlos0/env)

2. We utilize the package [`oauth1](https://github.com/dghubble/go-twitter/tree/master/examples)` 
to authenticate or app.

3. Instantiate a new client.

4. Get the latest `tweetCount` tweets from our timeline. 
   `client.Timelines.HomeTimeline` will return a slice of our tweets, a http
   response, and an error.

5. We will create a file with the name `path`, which in our case is hardcoded
   to `"tweets.jsonl"`.

6. Then for our last trick we will output the number, screen name, and test of
   each tweet to stdout and convert the struct to json then to a string to be
   writen to file.

