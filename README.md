# Static https website host on AppEngine
Sandbox for Go handlers working with [Let's Encrypt](https://letsencrypt.org/) and [AppEngine](https://cloud.google.com/appengine)

# Goals
- Get a starter project to create a static website working on HTTPS (via Let's Encrypt)
- Use a free platform thanks to [Google AppEngine free quota](https://cloud.google.com/appengine/docs/quotas)
- Share knowledge about AppEngine, a great PaaS 
- Get feedbacks from the community
- Help people playing with [Go(lang)](https://golang.org), the perfect programming language for cloud platform

# How to create and deploy
- Create an app on [AppEngine](https://console.cloud.google.com/)
 - specify a name and if you want to use `us-central`, `europe-west` or `us-east1`
- Add `App Engine Admin API`
- Generate a deployment key using an account service and download the json file
- Deploy using a CI tool like [codeship](https://codeship.com/) or [circleci]()https://circleci.com)
 - Add env variable `GAE_SERVICE_ACCOUNT` with `xxx@appspot.gserviceaccount.com`
 - Add env variable `GAE_KEY_FILE_CONTENT` with the content of your json file
- Use `gcloud-install.sh` to install gcloud tool
- Specify your `<appname>` and `<version>` in the `gcloud-deploy.sh` 

# Activate HTTPS


# AppEngine static files configuration 
- Static files are cached for 30 days except for `index.html` (5 minutes)
- We specify the file extensions for static file in order to avoid conflict with `golang` files
