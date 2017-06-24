FROM alpine:3.5
ADD build/simpleCRUD /
CMD ["/simpleCRUD"]