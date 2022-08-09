# Telkom TNT Test

##### Q : 3. Apakah ada kesalahan dari script di bawah ini? Jika ada tolong jelaskan dimana letak kesalahannya dan bagaimana anda memperbaikinya. Jika tidak ada, tolong jelaskan untuk apa script di bawah ini.

##### FROM golang
##### ADD . /go/src/github.com/telkomdev/indihome/backend
##### WORKDIR /go/src/github.com/telkomdev/indihome
##### RUN go get github.com/tools/godep
##### RUN godep restore
##### RUN go install github.com/telkomdev/indihome
##### ENTRYPOINT /go/bin/indihome
##### LISTEN 80
#

## A : LISTEN diganti menjadi EXPOSE

### Q : 4. Menurut anda apakah tujuan penggunaan microservices?
## A : Untuk memudahkan maintenance and upgrade


### Q : Bagaimana cara index bekerja pada sebuah database?

## A : Index bekerja dengan cara menyusun data dengan terurut dengan cara membuat struktur data yang berisi kumpulan keys beserta referensinya ke actual data di table sehingga untuk mempercepat pencarian data


