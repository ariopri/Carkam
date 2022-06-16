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
function App() {
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
        </Routes>
      </Router>
      <Footer />
    </>
  );
}

export default App;
