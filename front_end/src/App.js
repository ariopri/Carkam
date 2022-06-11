import "./App.css";
import NavigationBar from "./components/NavigationBar";
import "./style/SlideShow.css";
import Beranda from "./components/Beranda";
import Footer from "./components/Footer";
import TentangKami from "./components/TentangKami";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
function App() {
  return (
    <>
      <NavigationBar />
      <Router>
        <Routes>
          <Route path="/" element={<Beranda />} />
          <Route path="/Tentangkami" element={<TentangKami />} />
        </Routes>
      </Router>
      <Footer />
    </>
  );
}

export default App;
