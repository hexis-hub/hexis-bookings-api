FROM iron/go:dev

WORKDIR /hexis-bookings
# Set an env var that matches your github repo name, replace treeder/dockergo here with your repo name
ENV SRC_DIR=/go/src/github.com/hexis-hub/hexis-bookings-api/
# Add the source code:
ADD . $SRC_DIR

EXPOSE 3000
# Build it:
RUN cd $SRC_DIR; go build -o hexis-bookings; cp hexis-bookings /hexis-bookings/

ENTRYPOINT ["./hexis-bookings"]