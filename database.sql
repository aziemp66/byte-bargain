CREATE TABLE user (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
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
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    kategori VARCHAR(255) NOT NULL,
    deskripsi TEXT NOT NULL,
    harga INT NOT NULL,
    stok INT NOT NULL,
    berat FLOAT NOT NULL
);