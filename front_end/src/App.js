import "./App.css";
import NavigationBar from "./components/NavigationBar";
import "./style/SlideShow.css";
import Beranda from "./components/Beranda";
import Footer from "./components/Footer";
import TentangKami from "./components/TentangKami";
import Login from "./components/Login-Page/Login";
import Register from "./components/Register-Page/Register";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import FAQ from "./components/FAQ";
import Kampus from "./components/Kampus";
import DetailKampus from "./components/DetailKampus";
import Masukkanreview from "./components/Masukkanreview";
import InputUnivBaru from "./components/InputUnivBaru";
import EditUniv from "./components/EditUniv";
function App() {
  const peringatan = () => {
    alert("Mohon mengisi dengan baik sebelum meriview");
  };
  const peringatan2 = () => {
    alert("Mohon memperhatikan data untuk diubah");
  };
  return (
    <>
      <NavigationBar />
      <Router>
        <Routes>
          <Route path="/" element={<Beranda />} />
          <Route path="/Tentangkami" element={<TentangKami />} />
          <Route path="/Login" element={<Login />} />
          <Route path="/Register" element={<Register />} />
          <Route path="/FAQ" element={<FAQ />} />
          <Route path="/Kampus" element={<Kampus />} />
          <Route path="/Kampus/detail" element={<DetailKampus />} />
          <Route
            path="/Kampus/detail/review"
            element={<Masukkanreview peringatan={peringatan} />}
          />
          <Route path="/Kampus/create" element={<InputUnivBaru />} />
          <Route
            path="/Kampus/detail/update"
            element={<EditUniv peringatan2={peringatan2} />}
          />
        </Routes>
      </Router>
      <Footer />
    </>
  );
}

export default App;
