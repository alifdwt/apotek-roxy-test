import { useMutation, useQueryClient } from "@tanstack/react-query";
import MasterBarang from "../../interface/MasterBarang";
import useToast from "../../hooks/useToast";
import axiosFetch from "../../config/axiosFetch";
import { ChangeEvent, MouseEvent, useState } from "react";
import { useDisclosure } from "@chakra-ui/react";

export default function useUpdateBarang(idBarang: string, data: MasterBarang) {
  const [nmBarang, setNmBarang] = useState(data.nm_barang);
  const [harga, setHarga] = useState(data.harga);
  const [qty, setQty] = useState(data.qty);

  const queryClient = useQueryClient();
  const { isOpen, onOpen, onClose } = useDisclosure();
  const toast = useToast();

  const mutation = useMutation({
    mutationFn: () => axiosFetch.put(`/master-barang/${idBarang}`, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["master-barang"] });
      onClose();
      toast("Success", "Barang berhasil di ubah", "success");
    },
    onError: (error) => {
      toast("Error", error.message, "error");
    },
  });

  const handleNmBarangChange = (event: ChangeEvent<HTMLInputElement>) => {
    setNmBarang(event.target.value);
  };

  const handleHargaChange = (event: ChangeEvent<HTMLInputElement>) => {
    setHarga(event.target.value);
  };

  const handleQtyChange = (event: ChangeEvent<HTMLInputElement>) => {
    setQty(parseInt(event.target.value));
  };

  const handleUpdateBarang = (e: MouseEvent<HTMLButtonElement>) => {
    e.stopPropagation();
    mutation.mutate();
  };

  return {
    nmBarang,
    harga,
    qty,
    handleNmBarangChange,
    handleHargaChange,
    handleQtyChange,
    handleUpdateBarang,
    isOpen,
    onOpen,
    onClose,
  };
}
