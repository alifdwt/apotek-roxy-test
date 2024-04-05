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
} from "@chakra-ui/react";
import useCreateBarang from "./hooks/useCreateBarang";

export default function AddMasterBarangButton() {
  const {
    idBarang,
    nmBarang,
    qty,
    harga,
    handleIdBarangChange,
    handleNmBarangChange,
    handleQtyChange,
    handleHargaChange,
    handleSubmit,
    isOpen,
    onClose,
    onOpen,
    mutation,
  } = useCreateBarang();

  return (
    <>
      <Button onClick={onOpen}>Add</Button>

      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>Add Master Barang</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            <Box as={"form"}>
              <Box>
                <Text>ID Barang</Text>
                <Input
                  placeholder="Barang"
                  value={idBarang}
                  onChange={handleIdBarangChange}
                />
              </Box>
              <Box>
                <Text>Nama Barang</Text>
                <Input
                  placeholder="Nama Barang"
                  value={nmBarang}
                  onChange={handleNmBarangChange}
                />
              </Box>
              <Box>
                <Text>Quantity</Text>
                <Input
                  placeholder="Quantity"
                  value={qty}
                  onChange={handleQtyChange}
                />
              </Box>
              <Box>
                <Text>Harga</Text>
                <Input
                  placeholder="Harga"
                  value={harga}
                  onChange={handleHargaChange}
                />
              </Box>
              <Button
                onClick={handleSubmit}
                isLoading={mutation.isPending}
                w={"full"}
                mt={4}
              >
                Submit
              </Button>
            </Box>
          </ModalBody>
          <ModalFooter>
            <Button onClick={onClose}>Close</Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  );
}
