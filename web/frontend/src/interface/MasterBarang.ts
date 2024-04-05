interface MasterBarang {
  id_barang: string;
  nm_barang: string;
  qty: number;
  harga: string;
}

interface EditMasterBarang {
  nm_barang: string;
  qty: number;
  harga: string;
}

export default MasterBarang;

export type { EditMasterBarang };
