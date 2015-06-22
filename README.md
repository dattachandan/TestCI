# TestCI
Google Go test CI


[![Build Status](https://travis-ci.org/dattachandan/TestCI.svg?branch=master)](https://travis-ci.org/dattachandan/TestCI)


Setup tutorial on Windows
=========


## Step 1 - Check GOROOT 


Refer [GO docs](http://golang.org/doc/code.html) and [screencast](https://www.youtube.com/watch?v=XCsL89YtqCs)

![Check GOROOT ](/snaps/1.jpg?raw=true "Check GOROOT ")
![Check GOPATH ](/snaps/2.jpg?raw=true)

## Step 2 - Test Coverage

We can use gocov to analyse how much of our codebase is covered by tests. Install gocov. Note that gocov has both a library and a CLI, thus the gocov/gocov path
``` go get -v github.com/axw/gocov/gocov ```

![gocov error](/snaps/gocov_error.jpg?raw=true)

If you get the error below, add the GOPATH bin to the PATH environment variable

![gocov env](/snaps/gopath_env.jpg?raw=true)

If the system variables are set correctly, then you should be able to see the output as below
![gocov output](/snaps/gocov_output.jpg?raw=true)

## Step 3 - Setup a CI pipeline with github and Travis 

I want a feature where when I push a commit to Github, it will automatically kick off a Travis build

* Travis reads its configuration from a file in the repository root. Create a .travis.yml file
* Activate Travis - Open https://travis-ci.org in your browser. Click on “Sign in with Github”, in the upper right corner of the screen. If you have never logged in to travis before you will be prompted by Github to grant OAuth permissions to Travis. Once you are logged in, click on your username in the upper right corner, and choose “Accounts” from the dropdown. Travis will display a list of all your Github repositories. Click the OFF/ON toggle beside it. The toggle will slide to “ON”.
* Now we push our commits to Github, which will automatically kick off a Travis build

![git push](/snaps/git_push.jpg?raw=true)

Check Travis for automatically building once code is pushed

### Travis failed
![travis failed](/snaps/travis_failed.jpg?raw=true)

### Travis success
![travis success](/snaps/travis_success.jpg?raw=true)

	