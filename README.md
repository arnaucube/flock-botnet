# flock-botnet [![Go Report Card](https://goreportcard.com/badge/github.com/arnaucode/flock-botnet)](https://goreportcard.com/report/github.com/arnaucode/flock-botnet)

A twitter botnet with autonomous bots replying tweets with text generated based on probabilities in Markov chains

### generating text with Markov chains
Markov chain: https://en.wikipedia.org/wiki/Markov_chain

The algorithm calculates the probabilities of Markov chains, analyzing a considerable amount of text, for the examples, I've done it with the book "The Critique of Pure Reason", by Immanuel Kant (http://www.gutenberg.org/cache/epub/4280/pg4280.txt).

### Replying tweets with Markov chains
When the botnet is up working, the bots start streaming all the twitter new tweets containing the configured keywords. Each bot takes a tweet, analyzes the containing words, and generates a reply using the Markov chains previously calculated, and posts the tweet as reply.

In the following examples, the bots ("andreimarkov", "dodecahedron", "projectNSA") are replying some people.

![flock-botnet](https://raw.githubusercontent.com/arnaucode/flock-botnet/master/screenshots/01.png "01")

-

![flock-botnet](https://raw.githubusercontent.com/arnaucode/flock-botnet/master/screenshots/02.jpeg "02")

-

![flock-botnet](https://raw.githubusercontent.com/arnaucode/flock-botnet/master/screenshots/03.jpeg "03")

-

![flock-botnet](https://raw.githubusercontent.com/arnaucode/flock-botnet/master/screenshots/04.jpeg "04")


configuration file example (flockConfig.json):
```
[{
        "title": "account1",
        "consumer_key": "xxxxxxxxxxxxx",
        "consumer_secret": "xxxxxxxxxxxxx",
        "access_token_key": "xxxxxxxxxxxxx",
        "access_token_secret": "xxxxxxxxxxxxx"
    },
    {
        "title": "account2",
        "consumer_key": "xxxxxxxxxxxxx",
        "consumer_secret": "xxxxxxxxxxxxx",
        "access_token_key": "xxxxxxxxxxxxx",
        "access_token_secret": "xxxxxxxxxxxxx"
    },
    {
        "title": "account3",
        "consumer_key": "xxxxxxxxxxxxx",
        "consumer_secret": "xxxxxxxxxxxxx",
        "access_token_key": "xxxxxxxxxxxxx",
        "access_token_secret": "xxxxxxxxxxxxx"
    }
]

```
