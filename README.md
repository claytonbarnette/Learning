## Learning Go!

![Learning Go][docker-golang]

This repo contains files from my first steps into Go! Programming language. Like learning any other language, I start with the basics and expand from there. 

## The Base 6 Languages
Go! is what I call a **Base 6**, simply put, these are language stacks that are supported by the different Cloud platform providers such as AWS, Azure, and GCP. The Base 6 languages are _Java, Javascript (_Node_), Python, PHP7, Ruby and Go!_ or *JJPPRG*. There are some language stacks that include C# or other programming languages and environments, but for most platforms the Base 6 languages are supported. 

## The "Base 6" term:
The Base 6 language term is not an industry term, but is a term used only by me to identify the language(s) as one(s) that can be used in a Cloud environments. 

## Docker

I used the **golang** official Docker image from **[Docker Hub](https://hub.docker.com/_/golang)** to get started. 

- I first installed **[Docker](https://www.docker.com/)** 
- Opened up a terminal 
- Then pulled the **golang** image using the following command: `docker pull golang`
- Once the Docker image was downloaded, in a terminal, I ran the following command: `docker run -it golang` to start the environment. 

**Note:** I found that no editors were installed on this image, so to edit files I had to
install nano. To install the editor of your choice, run the following command: `apt update` followed by the 
commands required to install your editor: `apt install nano` or `apt install vi`. 


Also, to run the Go! programs from the command line from any directory, I set my `PATH` environment variable
to include the **`$GOPATH`**. Information on setting that up can be found [here](https://golang.org/doc/code.html) if needed.

## Self Learning

![Linkedin Learning][linkedinlearning]

I completed the **"Up and Running with Go"** course with *David Gassner* on __Lynda.com__, which is now [Linkedin Learning](https://www.linkedin.com/learning/) 
 to help me get upto speed, plus I spent some quality time with the documentation at [https://golang.org](https://golang.org) 
website. 

I hope this helps you and other get started or take a second look at the [Go!](https://golang.org) programming language.
  

![Learning Go][go-logo]


-- Clayton Barnette

[linkedinlearning]: https://media.licdn.com/dms/image/C4E0BAQFfDMello2Gtg/company-logo_200_200/0?e=2159024400&v=beta&t=ectgju-UZL2EC0Fmufx6CJ9PToe2bKXwZZu6qg6qJfs
[docker-golang]: https://d1q6f0aelx0por.cloudfront.net/product-logos/81630ec2-d253-4eb2-b36c-eb54072cb8d6-golang.png
[go-logo]: https://golang.org/doc/gopher/doc.png


