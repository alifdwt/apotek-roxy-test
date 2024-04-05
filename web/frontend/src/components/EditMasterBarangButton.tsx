import {
  Box,
  Button,
  Input,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  Text,
  useDisclosure,
} from "@chakra-ui/react";
import useToast from "../hooks/useToast";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { ChangeEvent, MouseEvent, useEffect, useState } from "react";
import useFetchWithId from "../hooks/useFetchWithId";
import axiosFetch from "../config/axiosFetch";
import { EditMasterBarang } from "../interface/MasterBarang";

export default function EditBarangButton({ id }: { id: string }) {
  const toast = useToast();
  const queryClient = useQueryClient();
  const [nmBarang, setNmBarang] = useState("");
  const [qty, setQty] = useState(0);
  const [harga, setHarga] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const { data: fetchedMasterBarangData } = useFetchWithId({
    queryKey: "master-barang",
    id: id,
    fetchRoute: "/master-barang",
  });

  const masterBarangData: EditMasterBarang = fetchedMasterBarangData;

  const { isOpen, onOpen, onClose } = useDisclosure();

  const mutation = useMutation({
    mutationFn: (editedMasterBarang: EditMasterBarang) => {
      return axiosFetch.put(`/master-barang/${id}`, editedMasterBarang);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["master-barang"] });
      toast("Success", "Barang berhasil di ubah", "success");
      setIsLoading(false);
      onClose();
    },
    onError: (error) => {
      // @ts-expect-error error
      toast("Error", error.response.data.message, "error");
      setIsLoading(false);
    },
  });

  useEffect(() => {
    const { nm_barang, qty, harga } = masterBarangData;
    setNmBarang(nm_barang);
    setQty(qty);
    setHarga(harga);
  }, [masterBarangData]);

  const handleOnChangeNmBarang = (event: ChangeEvent<HTMLInputElement>) => {
    setNmBarang(event.target.value);
  };

  const handleOnChangeQty = (event: ChangeEvent<HTMLInputElement>) => {
    setQty(parseInt(event.target.value));
  };

  const handleOnChangeHarga = (event: ChangeEvent<HTMLInputElement>) => {
    setHarga(event.target.value);
  };

  const handleUpdateBarang = (e: MouseEvent<HTMLButtonElement>) => {
    e.stopPropagation();
    if (nmBarang === "" || qty === 0 || harga === "") {
      toast("Error", "Mohon lengkapi semua data", "error");
    }

    mutation.mutate({
      nm_barang: nmBarang,
      qty: qty,
      harga: harga,
    });
  };

  return (
    <>
      <Button onClick={onOpen}>Edit</Button>

      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>Add Master Barang</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            <Box as={"form"}>
              <Box>
                <Text>Nama Barang</Text>
                <Input
                  placeholder="Nama Barang"
                  value={nmBarang}
                  onChange={handleOnChangeNmBarang}
                />
              </Box>
              <Box>
                <Text>Quantity</Text>
                <Input
                  placeholder="Quantity"
                  value={qty}
                  onChange={handleOnChangeQty}
                />
              </Box>
              <Box>
                <Text>Harga</Text>
                <Input
                  placeholder="Harga"
                  value={harga}
                  onChange={handleOnChangeHarga}
                />
              </Box>
              <Button w={"full"} mt={4} onClick={handleUpdateBarang}>
                Submit
              </Button>
            </Box>
          </ModalBody>
          <ModalFooter>
            <Button isLoading={isLoading}>Close</Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  );
}
