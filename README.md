# Static https website host on AppEngine
Sandbox for Go handlers working with [Let's Encrypt](https://letsencrypt.org/) and [AppEngine](https://cloud.google.com/appengine)

# Goals
- Get a starter project to create a static website working on HTTPS (via Let's Encrypt)
- Use a free platform thanks to [Google AppEngine free quota](https://cloud.google.com/appengine/docs/quotas)
- Share knowledge about AppEngine, a great PaaS
- Get feedback from the community
- Help people playing with [Go(lang)](https://golang.org), the perfect programming language for cloud platform

# How to create and deploy
- Create an app on [AppEngine](https://console.cloud.google.com/)
 - specify a name and if you want to use `us-central`, `europe-west` or `us-east1`
- Add `App Engine Admin API`
- Generate a deployment key using an account service and download the json file
- Deploy using a CI tool like [codeship](https://codeship.com/) or [circleci](https://circleci.com)
 - Add env variable `GAE_SERVICE_ACCOUNT` with `xxx@appspot.gserviceaccount.com`
 - Add env variable `GAE_KEY_FILE_CONTENT` with the content of your json file
- Use `gcloud-install.sh` to install gcloud tool
- Specify your `$APPNAME` and `$VERSION` in the script `gcloud-deploy.sh` or in your CI

# Activate HTTPS with Let's encrypt on AppEngine
- Launch the following command using `docker`
```
docker run -it -p 443:443 -p 80:80 \
  -v "$(pwd)/ssl-keys:/etc/letsencrypt" \
  quay.io/letsencrypt/letsencrypt:latest \
  -a manual certonly
```
- Go to `https.go` and change the challenge (L14)
- Enter an email for urgent notices, accept the terms and enter the domain to secure
- Commit and upload the application with the new challenge
- Complete the docker command line
- Go to `/ssl-keys/live/<mydomain>/` and upload the files on [AppEngine settings for certificate](https://console.cloud.google.com/appengine/settings/certificates):
 - fullchain.pem
 - rsa.pem using the commande line `openssl rsa -in privkey.pem -out rsa.pem`
- To get the file with Docker, execute `docker start $(docker ps -ql)`
- Then go inside with `docker exec -it $(docker ps -ql) bash` and `cd /etc/letsencrypt/live/<mydomain>`

# AppEngine static files configuration
- Static files are cached for 30 days except for `index.html` (5 minutes)
- We specify the file extensions for static file in order to avoid conflict with `golang` files

#Circle CI configuration file
See the `circle.yml` configuration file for an example. Don't forget to set $APPNAME and $VERSION in the env variables.
