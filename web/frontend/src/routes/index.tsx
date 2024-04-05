import MasterBarangPage from "@/pages/MasterBarang/MasterBarang";
import TransaksiPage from "@/pages/Transaksi/Transaksi";
import { BrowserRouter, Route, Routes as RouterRoutes } from "react-router-dom";

export default function Routes() {
  return (
    <BrowserRouter>
      <RouterRoutes>
        {/* <Route path="/test"> */}
        <Route path="/test/master" element={<MasterBarangPage />} />
        <Route path="/test/transaksi" element={<TransaksiPage />} />
        {/* </Route> */}
      </RouterRoutes>
    </BrowserRouter>
  );
}
