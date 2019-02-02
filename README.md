## Learning Go!

![Learning Go][docker-golang]

This repo contains files from my first steps into Go! Programming language. Like many of the other
languages that I've learnt over the years, I start with the basics and expand from there. Go! is 
becoming one of the **Base 6** languages which are used by most of the cloud providers. The Base 6
languages are *JJPPRG*: Java, Javascript, Python, PHP, Ruby and Go or *JJPPRC* which includes C#. Both
are valid Base 6 language sets and *"Base 6"* is not an industry term, but is just a term used my me to
reference the language stacks used in a cloud environment. 

### Docker

I used the **golang** official Docker image to get started, this image is located on 
**[Docker Hub](https://hub.docker.com/_/golang)**. To pull this image, I first installed Docker,
then use the following command `docker pull golang` to pull down the golang image. 

Once the Docker image was downloaded, I started a terminal, then ran the following command: `docker run -it golang` 
to get the Go! environment up and running. I found that no editors are installed on this image, so to edit files I
installed nano, but you can install the editor of your choice by first doing an `apt update` followed by the 
commands required to install your editor: `apt install nano` or `apt install vi`. 


Also, to run the Go! programs from the command line from any directory, I set my `PATH` environment variable
to include the **`$GOPATH`**. Information on setting that up can be found [here](https://golang.org/doc/code.html) if needed.

## Self Learning

![Linkedin Learning][linkedinlearning]

I completed the **"Up and Running with Go"** course on __Lynda.com__, which is now [Linkedin Learning](https://www.linkedin.com/learning/) 
with *David Gassner* to help me get upto speed, plus I spent some quality time with the documentation at [https://golang.org](https://golang.org) 
website. I hope this helps other get started or take a second look at the Go! programming language.
  

![Learning Go][go-logo]


-- Clayton Barnette

[linkedinlearning]: https://media.licdn.com/dms/image/C4E0BAQFfDMello2Gtg/company-logo_200_200/0?e=2159024400&v=beta&t=ectgju-UZL2EC0Fmufx6CJ9PToe2bKXwZZu6qg6qJfs
[docker-golang]: https://d1q6f0aelx0por.cloudfront.net/product-logos/81630ec2-d253-4eb2-b36c-eb54072cb8d6-golang.png
[go-logo]: https://golang.org/doc/gopher/doc.png


