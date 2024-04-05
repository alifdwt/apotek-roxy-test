import {
  Box,
  Button,
  Flex,
  Modal,
  ModalBody,
  ModalContent,
  ModalOverlay,
  Table,
  TableContainer,
  Tbody,
  Text,
  Th,
  Thead,
  Tr,
  useDisclosure,
} from "@chakra-ui/react";
import MasterBarang from "../interface/MasterBarang";

export default function SaveMasterBarangButton({
  data,
}: {
  data: MasterBarang[];
}) {
  const { isOpen, onOpen, onClose } = useDisclosure();
  // time now 04-04-2024 16:48
  const timeNow = new Date().toLocaleString();

  return (
    <>
      <Button onClick={onOpen}>Save</Button>

      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalBody>
            <Box as={Flex} flexDir={"column"} gap={4}>
              <TableContainer>
                <Table>
                  <Thead bg={"gray.100"}>
                    <Tr>
                      <Th>Nama barang</Th>
                      <Th>Qty</Th>
                      <Th>Total Harga</Th>
                    </Tr>
                  </Thead>
                  <Tbody>
                    {data.map((masterBarang) => (
                      <Tr key={masterBarang.id_barang}>
                        <Th>{masterBarang.nm_barang}</Th>
                        <Th>{masterBarang.qty}</Th>
                        <Th>
                          {(
                            parseInt(masterBarang.harga) * masterBarang.qty
                          ).toLocaleString()}
                        </Th>
                      </Tr>
                    ))}
                    <Tr bg={"gray.100"}>
                      <Th>Total</Th>
                      <Th>
                        {data.reduce((total, masterBarang) => {
                          return total + masterBarang.qty;
                        }, 0)}
                      </Th>
                      <Th>
                        {data
                          .reduce((total, masterBarang) => {
                            return (
                              total +
                              parseInt(masterBarang.harga) * masterBarang.qty
                            );
                          }, 0)
                          .toLocaleString()}
                      </Th>
                    </Tr>
                  </Tbody>
                </Table>
              </TableContainer>
              <Text>Print: {timeNow} By ALIF</Text>
              <Box textAlign={"center"}>
                <Text>Apotek Roxy</Text>
                <Text>Jl. Raya Jatikramat No. 14 Bekasi</Text>
                <Text>No. Telp (021) 9497-0558</Text>
              </Box>
            </Box>
          </ModalBody>
        </ModalContent>
      </Modal>
    </>
  );
}
