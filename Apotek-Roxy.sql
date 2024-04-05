CREATE TABLE `master_barang` (
  `id_barang` int PRIMARY KEY,
  `nm_barang` varchar(30),
  `Qty` int(10),
  `Harga` text
);

CREATE TABLE `transaksi_hider` (
  `id_trans` varchar(15) PRIMARY KEY,
  `tgl_trans` datetime,
  `total` int
);

CREATE TABLE `transaksi_detail` (
  `id_trans` varchar(15) PRIMARY KEY,
  `id_barang` varchar(15),
  `qty` int(10)
);

ALTER TABLE `transaksi_detail` ADD FOREIGN KEY (`id_trans`) REFERENCES `transaksi_hider` (`id_trans`);

ALTER TABLE `transaksi_detail` ADD FOREIGN KEY (`id_barang`) REFERENCES `master_barang` (`id_barang`);
