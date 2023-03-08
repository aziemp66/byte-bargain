CREATE TABLE user (
    id_user INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    jenis_user ENUM('Penjual','Pembeli'),
    email VARCHAR(255) NOT NULL,
    nama VARCHAR(255) NOT NULL,
    alamat VARCHAR(255) NOT NULL,
    tgl_lahir DATE NOT NULL,
    no_hp VARCHAR(20) NOT NULL,
    jenis_kelamin ENUM('P', 'L') NOT NULL,
    kartu_identitas VARCHAR(255) NULL,
    rekening_bank VARCHAR(255) NULL
);
CREATE TABLE product (
    id_produk INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    nama_produk VARCHAR(255) NOT NULL,
    kategori VARCHAR(255) NOT NULL,
    deskripsi TEXT NOT NULL,
    harga INT NOT NULL,
    stok INT NOT NULL,
    berat FLOAT NOT NULL
);
CREATE TABLE pesanan (
    id_pesanan INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    id_user INT NOT NULL,
    id_produk INT NOT NULL,
    tgl_pesan DATE NOT NULL,
    CONSTRAINT fk_user_pesanan
    FOREIGN KEY (id_user) REFERENCES user(id_user),
    CONSTRAINT fk_produk_pesanan
    FOREIGN KEY (id_produk) REFERENCES product(id_produk)
);
CREATE TABLE bayar (
    id_bayar INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    id_pesanan INT NOT NULL,
    tgl_bayar DATE NOT NULL,
    total_bayar INT NOT NULL,
    CONSTRAINT fk_bayar_pesanan
    FOREIGN KEY (id_pesanan) REFERENCES bayar(id_pesanan)
);