# Go Vanity URL Server
A minimal implementation of a Go vanity URL server.

## Introduction:
Want to use your own domain name to share your Go modules like this?
```console
go get bjvanbemmel.nl/cosmos
go install bjvanbemmel.nl/kyoto
```
This repo will allow you to get your own vanity URL up and running within seconds.

## Installation:
The recommended way to run this proxy is by cloning the repository and starting the Docker container.

1. Clone the repository:
```console
git clone git@github.com:bjvanbemmel/go_proxy
cd go_proxy
```

2. Configure the environment file:
```console
cp .env.example .env
vi .env
```
```env
HOST_PORT=88 # Example port. Just pick any port you want.
```

3. Start the Docker container:
```console
docker compose up -d
```

### Optional: Separation between a normal user and the `go get` client:
In case you want to only redirect the `go get` client and not all regular visitors, apply the following code snippet to your NGINX proxy:

4. Within your NGINX config, add the following code snippet:
```nginx
server {
        # server_name mywebsite.com www.mywebsite.com;

        # root /var/www/website;

        location ~ /(.*) {
                if ($http_user_agent = "Go-http-client/1.1") {
                        proxy_pass http://127.0.0.1:88/$1; # Enter whichever port you defined in the .env file
                }
        }

        # location / {
            # if (!-e $request_filename) {
                # rewrite ^ / permanent;
            # }
        #}
}
```
*Commented lines don't have to be copied and exist solely to show you where to place the snippet.*

5. Restart your NGINX service and you're done:
```console
systemctl restart nginx
```
