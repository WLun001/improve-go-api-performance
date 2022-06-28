# warning
# this docker file is for educative onlu

FROM golang:1.18
RUN git clone https://github.com/WLun001/improve-go-api-performance.git \
  && cd improve-go-api-performance \
  && go mod download
