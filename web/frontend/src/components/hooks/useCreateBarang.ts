import { useMutation, useQueryClient } from "@tanstack/react-query";
import { ChangeEvent, useState } from "react";
import useToast from "../../hooks/useToast";
import MasterBarang from "../../interface/MasterBarang";
import axiosFetch from "../../config/axiosFetch";
import { useDisclosure } from "@chakra-ui/react";

export default function useCreateBarang() {
  const queryClient = useQueryClient();
  const [idBarang, setIdBarang] = useState("");
  const [nmBarang, setNmBarang] = useState("");
  const [qty, setQty] = useState(0);
  const [harga, setHarga] = useState("");

  const toast = useToast();
  const form = new FormData();
  const { isOpen, onOpen, onClose } = useDisclosure();

  const mutation = useMutation({
    mutationFn: (newMasterBarang: MasterBarang) =>
      axiosFetch.post("/master-barang/", newMasterBarang),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["master-barang"] });
      onClose();
      toast("Success", "Barang berhasil ditambahkan", "success");
    },
    onError: (error) => {
      console.log(error);
      toast("Error", error.message, "error");
    },
  });

  const handleIdBarangChange = (event: ChangeEvent<HTMLInputElement>) => {
    setIdBarang(event.target.value);
  };

  const handleNmBarangChange = (event: ChangeEvent<HTMLInputElement>) => {
    setNmBarang(event.target.value);
  };

  const handleQtyChange = (event: ChangeEvent<HTMLInputElement>) => {
    setQty(parseInt(event.target.value));
  };

  const handleHargaChange = (event: ChangeEvent<HTMLInputElement>) => {
    setHarga(event.target.value);
  };

  const handleSubmit = () => {
    form.append("id_barang", idBarang);
    form.append("nm_barang", nmBarang);
    form.append("qty", qty.toString());
    form.append("harga", harga);
    mutation.mutate({
      id_barang: idBarang,
      nm_barang: nmBarang,
      qty: qty,
      harga: harga,
    });

    setIdBarang("");
    setNmBarang("");
    setQty(0);
    setHarga("");
  };

  return {
    idBarang,
    nmBarang,
    qty,
    harga,
    isOpen,
    onOpen,
    onClose,
    setIdBarang,
    setNmBarang,
    setQty,
    setHarga,
    handleIdBarangChange,
    handleNmBarangChange,
    handleQtyChange,
    handleHargaChange,
    handleSubmit,
    mutation,
  };
}
