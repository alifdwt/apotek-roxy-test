import { useMutation, useQueryClient } from "@tanstack/react-query";
import { ChangeEvent, useState } from "react";
import axiosFetch from "../../config/axiosFetch";
import useToast from "../../hooks/useToast";

export default function useSearchBarang() {
  const [searchNmBarang, setSearchNmBarang] = useState("");

  const queryClient = useQueryClient();
  const toast = useToast();
  const form = new FormData();

  const mutation = useMutation({
    mutationFn: () => axiosFetch.get(`/master-barang/${searchNmBarang}`),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["master-barang"] });
    },
    onError: (error) => {
      toast("Error", error.message, "error");
    },
  });

  const handleSearchNmBarangChange = (event: ChangeEvent<HTMLInputElement>) => {
    setSearchNmBarang(event.target.value);
  };

  const handleSubmit = () => {
    form.append("nm_barang", searchNmBarang);
    mutation.mutate();
  };

  return {
    searchNmBarang,
    handleSearchNmBarangChange,
    handleSubmit,
  };
}
