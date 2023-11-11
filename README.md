### Pluggable OAuth library for social login (Google, Facebook, Github, Linkedin)


[![GoDoc](https://godoc.org/github.com/wuriyanto48/go-social?status.svg)](https://godoc.org/github.com/wuriyanto48/go-social)
[![Build Status](https://travis-ci.org/wuriyanto48/go-social.svg?branch=master)](https://travis-ci.org/wuriyanto48/go-social)
[![CircleCI](https://circleci.com/gh/wuriyanto48/go-social.svg?style=svg)](https://circleci.com/gh/wuriyanto48/go-social)

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com) [![forthebadge](https://forthebadge.com/images/badges/fuck-it-ship-it.svg)](https://forthebadge.com) [![forthebadge](https://forthebadge.com/images/badges/contains-technical-debt.svg)](https://forthebadge.com)

### Install
```shell
$ go get github.com/wuriyanto48/go-social
```

### Usage

OAuth2 using Facebook login

* Getting Authorization Code First
`https://www.facebook.com/dialog/oauth?client_id={your_client_id}&redirect_uri=http://localhost:8080/callback&response_type=code`

* Place the Authorization Code to the second parameter of `GetAccessToken(ctx, "authorization_code")` function

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/wuriyanto48/go-social"
	"github.com/wuriyanto48/go-social/pkg/facebook"
)

func main() {
	f, err := social.New(social.Facebook, "client_id", "client_secret", "", "http://localhost:8080/callback", "")

	if err != nil {
		fmt.Println(err)
	}

        // using context for cancellation
        ctx := context.Background()
        // set context timeout
	ctx, _ = context.WithTimeout(ctx, 2000*time.Millisecond)

	err = f.GetAccessToken(ctx, "authorization_code")

	if err != nil {
		fmt.Println(err)
	}

	result, err := f.GetUser(ctx)

	if err != nil {
		fmt.Println(err)
	}

	user, _ := result.(*facebook.User)

	fmt.Println(user.Picture)
}

```

### Todo
- Add Twitter implementation