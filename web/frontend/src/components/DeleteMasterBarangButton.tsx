import { Button } from "@chakra-ui/react";
import useDeleteBarang from "./hooks/useDeleteBarang";

export default function DeleteMasterBarangButton({ id }: { id: string }) {
  const { handleDeleteBarang } = useDeleteBarang(id);

  return (
    <>
      <Button onClick={handleDeleteBarang}>Hapus</Button>
    </>
  );
}
