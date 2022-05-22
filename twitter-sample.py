import tweepy
import sys
def main():
    twitter_auth_keys = {
        "consumer_key"        : "changeme",
        "consumer_secret"     : "changeme",
        "access_token"        : "changeme",
        "access_token_secret" : "changeme"
    }
    auth = tweepy.OAuthHandler(
            twitter_auth_keys['consumer_key'],
            twitter_auth_keys['consumer_secret']
            )
    auth.set_access_token(
            twitter_auth_keys['access_token'],
            twitter_auth_keys['access_token_secret']
            )
    api = tweepy.API(auth)
    # Upload image
    media = api.media_upload("twitter.png")
    url = sys.argv[2]
    title = sys.argv[1]
    tweet = "New Post: " + title + "\n \n Click link to read article ⬇️ \n" + url
    post_result = api.update_status(status=tweet, media_ids=[media.media_id])
if __name__ == "__main__":
    main()