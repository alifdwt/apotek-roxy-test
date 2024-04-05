import AddMasterBarangButton from "@/components/AddMasterBarangButton";
import DeleteMasterBarangButton from "@/components/DeleteMasterBarangButton";
import EditBarangButton from "@/components/EditMasterBarangButton";
import SaveMasterBarangButton from "@/components/SaveMasterBarangButton";
import SearchMasterBarangSection from "@/components/SeachMasterBarangSection";
import axiosFetch from "@/config/axiosFetch";
import MasterBarang from "@/interface/MasterBarang";
import {
  Box,
  Flex,
  Heading,
  Table,
  TableContainer,
  Tbody,
  Th,
  Thead,
  Tr,
} from "@chakra-ui/react";
import { useQuery } from "@tanstack/react-query";

export default function MasterBarangPage() {
  const { data, isLoading } = useQuery({
    queryKey: ["master-barang"],
    queryFn: async () => {
      const { data } = await axiosFetch.get("/master-barang/");
      return data;
    },
  });

  const masterBarangs: MasterBarang[] = data;

  if (isLoading) {
    return <Box>Loading...</Box>;
  }

  return (
    <Box>
      <Flex gap={4} flexDir={"column"} p={4}>
        <Heading textAlign={"center"}>Master Barang</Heading>
        <SearchMasterBarangSection />
        <Flex gap={2}>
          <AddMasterBarangButton />
          <SaveMasterBarangButton data={masterBarangs} />
        </Flex>
      </Flex>
      <TableContainer>
        <Table variant={"simple"}>
          <Thead>
            <Tr>
              <Th>ID</Th>
              <Th>Barang</Th>
              <Th>Qty</Th>
              <Th>Harga</Th>
              <Th>Option</Th>
            </Tr>
          </Thead>
          <Tbody>
            {masterBarangs.map((masterBarang) => (
              <Tr key={masterBarang.id_barang}>
                <Th>{masterBarang.id_barang}</Th>
                <Th>{masterBarang.nm_barang}</Th>
                <Th>{masterBarang.qty}</Th>
                <Th>{parseInt(masterBarang.harga).toLocaleString()}</Th>
                <Th as={Flex} gap={2}>
                  <EditBarangButton id={masterBarang.id_barang} />
                  <DeleteMasterBarangButton id={masterBarang.id_barang} />
                </Th>
              </Tr>
            ))}
          </Tbody>
        </Table>
      </TableContainer>
    </Box>
  );
}
