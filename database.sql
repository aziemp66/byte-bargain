CREATE TABLE `user` (
    `id_user` VARCHAR(36) PRIMARY KEY,
    `email` VARCHAR(225) UNIQUE NOT NULL,
    `password` VARCHAR(225) NOT NULL
);
CREATE TABLE `pembeli` (
    `id_pembeli` VARCHAR(36) PRIMARY KEY,
    `id_user` VARCHAR(36) UNIQUE NOT NULL,
    `nama` VARCHAR(225) NOT NULL,
    `alamat` VARCHAR(225) NOT NULL,
    `tgl_lahir` DATE NOT NULL,
    `no_hp` VARCHAR(20) UNIQUE NOT NULL,
    `jenis_kelamin` ENUM('Laki-laki', 'Perempuan') NOT NULL
);
CREATE TABLE `penjual` (
    `id_penjual` VARCHAR(36) PRIMARY KEY,
    `id_user` VARCHAR(36) UNIQUE NOT NULL,
    `nama` VARCHAR(225) NOT NULL,
    `saldo` FLOAT NOT NULL,
    `alamat` VARCHAR(225) NOT NULL,
    `tgl_lahir` DATE NOT NULL,
    `no_hp` VARCHAR(20) UNIQUE NOT NULL,
    `jenis_kelamin` ENUM('Laki-laki', 'Perempuan') NOT NULL,
    `kartu_identitas` VARCHAR(225) NOT NULL,
    `rekening_bank` VARCHAR(225) NOT NULL,
    `nomor_rekening` INT NOT NULL
);
CREATE TABLE `produk` (
    `id_produk` VARCHAR(36) PRIMARY KEY,
    `id_penjual` VARCHAR(36) NOT NULL,
    `nama_produk` VARCHAR(225) NOT NULL,
    `kategori` VARCHAR(225) NOT NULL,
    `deskripsi` TEXT NOT NULL,
    `harga` FLOAT NOT NULL,
    `stok` INT NOT NULL,
    `berat` FLOAT NOT NULL
);
CREATE TABLE `pesanan` (
    `id_pesanan` VARCHAR(36) PRIMARY KEY,
    `id_pembeli` VARCHAR(36) NOT NULL,
    `id_penjual` VARCHAR(36) NOT NULL,
    `waktu_transaksi` DATETIME NOT NULL,
    `status` VARCHAR(225) NOT NULL
);
CREATE TABLE `pesanan_produk` (
    `id_pesanan_produk` VARCHAR(36) PRIMARY KEY,
    `id_produk` VARCHAR(36) NOT NULL,
    `id_pesanan` VARCHAR(36) NOT NULL,
    `qty` INT NOT NULL
);
CREATE TABLE `pembayaran` (
    `id_pembayaran` VARCHAR(36) PRIMARY KEY,
    `id_pesanan` VARCHAR(36) UNIQUE NOT NULL,
    `tgl_bayar` DATE NOT NULL,
    `total_bayar` FLOAT NOT NULL,
    `metode_pembayaran` VARCHAR(20) NOT NULL
);

ALTER TABLE `pembeli`
ADD FOREIGN KEY (`id_user`) REFERENCES `user` (`id_user`);
ALTER TABLE `penjual`
ADD FOREIGN KEY (`id_user`) REFERENCES `user` (`id_user`);
ALTER TABLE `produk`
ADD FOREIGN KEY (`id_penjual`) REFERENCES `penjual` (`id_penjual`);
ALTER TABLE `pesanan`
ADD FOREIGN KEY (`id_pembeli`) REFERENCES `pembeli` (`id_pembeli`);
ALTER TABLE `pesanan`
ADD FOREIGN KEY (`id_penjual`) REFERENCES `penjual` (`id_penjual`);
ALTER TABLE `pesanan_produk`
ADD FOREIGN KEY (`id_produk`) REFERENCES `produk` (`id_produk`);
ALTER TABLE `pesanan_produk`
ADD FOREIGN KEY (`id_pesanan`) REFERENCES `pesanan` (`id_pesanan`);
ALTER TABLE `pembayaran`
ADD FOREIGN KEY (`id_pesanan`) REFERENCES `pesanan` (`id_pesanan`);