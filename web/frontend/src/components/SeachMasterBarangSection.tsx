import { Box, Button, Input } from "@chakra-ui/react";
import useSearchBarang from "./hooks/useSearchBarang";

export default function SearchMasterBarangSection() {
  const { searchNmBarang, handleSearchNmBarangChange, handleSubmit } =
    useSearchBarang();
  return (
    <Box>
      <Box as={"form"}>
        <Input
          placeholder="Search"
          value={searchNmBarang}
          onChange={handleSearchNmBarangChange}
        />
        <Button onClick={handleSubmit}>Search</Button>
      </Box>
    </Box>
  );
}
