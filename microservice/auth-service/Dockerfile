# Menggunakan image Docker yang berisi Go
FROM golang:1.19.5

# Menentukan direktori kerja pada container
WORKDIR /PA2/microservice/auth-service

# Menyalin seluruh kode sumber ke dalam container
COPY . .

# Mengunduh semua dependensi dari go.mod dan go.sum
RUN go mod download

# Membangun binary dari kode sumber
RUN go build -o auth-service .

# Menjalankan layanan auth-service
CMD ["./auth-service"]
