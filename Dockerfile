FROM       quay.io/kelseyhightower/scratch
MAINTAINER Kelsey Hightower <kelsey.hightower@gmail.com>
ADD        journal-2-logentries journal-2-logentries
ENTRYPOINT ["/journal-2-logentries"]
