
# Accept the Go version for the image to be set as a build argument.
# Default to Go 1.12
ARG GO_VERSION=1.12

# First stage: build the executable.
FROM golang:${GO_VERSION}-alpine AS builder

# Create the user and group files that will be used in the running container to
# run the process as an unprivileged user.
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

# Set the working directory outside $GOPATH to enable the support for modules.
RUN mkdir /src
WORKDIR /src

# SSH stuff to pull from private repo
#ADD id_rsa /root/.ssh/id_rsa
#RUN chmod 700 /root/.ssh/id_rsa
#RUN echo "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config
#RUN apk add --no-cache git
#RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/

# Git is required for fetching the dependencies.
# gcc and libc-dev also required for go build
RUN apk update
#RUN apk add --no-cache git gcc libc-dev libgit2 libssh2 make cmake pkgconf openssl-dev 
RUN apk add --no-cache git gcc libc-dev libgit2-0.27-dev pkgconf 

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
# We want to populate the module cache based on the go.{mod,sum} files. 
COPY ./go.mod ./go.sum ./
RUN go mod download

# Import the code from the context.
COPY ./ ./ 

# Build the executable to `/app`. Mark the build as statically linked.
# CGO needs to be enabled because of the C bindings in libgit2
RUN CGO_ENABLED=1 GOOS=linux go build \
    -installsuffix 'static' \
    -o /app .

# Final stage: the running container.
FROM scratch AS final

# Import the user and group files from the first stage.
COPY --from=builder /user/group /user/passwd /etc/

# Import the compiled executable from the first stage.
COPY --from=builder /app /app

EXPOSE 9000

# Perform any further action as an unprivileged user.
USER nobody:nobody

# Run the compiled binary.
ENTRYPOINT ["/app"]