# Documentation

## Task Section
### Task 3
LISTEN diganti menjadi EXPOSE

### Task 4
Untuk memudahkan maintenance and upgrade

### Task 5
Index bekerja dengan cara menyusun data dengan terurut dengan cara membuat struktur data yang berisi kumpulan keys beserta referensinya ke actual data di table sehingga dapat mempercepat pencarian data

## Guide Section

### Rest API
Click [Postman Documentation](https://documenter.getpostman.com/view/4459576/VUxXJNt9#intro) to go to Rest API guide

### Run Unit Test
##### Step :
- Move to directory 1, 2 or 6/src
- type ``` go test -v ``` and then press enter

### Prepare & Run Docker
##### Step :
- Move to directory 1
- type ``` docker build -t fikriprayoga1/telkom-server:1.0 . ``` and then press enter
- If you don't have mongo image, type ``` docker pull mongo ``` and then press enter
- type ``` docker compose up -d ``` and then press enter 