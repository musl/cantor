FROM scratch
ADD cantor /cantor
ADD build /build
EXPOSE 9000
CMD ["/cantor"]
