import { useMutation, useQueryClient } from "@tanstack/react-query";
import useToast from "../../hooks/useToast";
import axiosFetch from "../../config/axiosFetch";
import { MouseEvent } from "react";

export default function useDeleteBarang(id: string) {
  const queryClient = useQueryClient();
  const toast = useToast();

  const mutation = useMutation({
    mutationFn: () => axiosFetch.delete(`/master-barang/${id}`),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["master-barang"] });
      toast("Success", "Barang berhasil di hapus", "success");
    },
    onError: (error) => {
      toast("Error", error.message, "error");
    },
  });

  const handleDeleteBarang = (e: MouseEvent<HTMLButtonElement>) => {
    e.stopPropagation();
    mutation.mutate();
  };

  return { handleDeleteBarang };
}
